package http

import (
	"encoding/json"
	"io"
)

func ReadBodyToStruct(body io.ReadCloser, v any) error {
	defer body.Close()
	return json.NewDecoder(body).Decode(v)
}
