package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func HandlerResponse(w http.ResponseWriter, r *http.Request, data any, err error) {
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		msg := formatError(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Status:  http.StatusBadRequest,
			Message: msg,
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    data,
	})
}

func formatError(err error) string {
	if ve, ok := err.(validator.ValidationErrors); ok {
		return fmt.Sprintf("code=400, message=%s", ve.Error())
	}
	msg := err.Error()
	if len(msg) > 0 {
		msg = strings.ToUpper(msg[:1]) + msg[1:]
		if !strings.HasSuffix(msg, ".") {
			msg += "."
		}
	}
	return msg
}
