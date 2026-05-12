package dummy

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	httpRequest "sbb-golang-template/pkg/http"
	"sbb-golang-template/pkg/models"
	"sbb-golang-template/pkg/response"
)

var validate = validator.New()

type DummyServiceHandlerInterface interface {
	list(w http.ResponseWriter, r *http.Request)
	info(w http.ResponseWriter, r *http.Request)
	create(w http.ResponseWriter, r *http.Request)
	update(w http.ResponseWriter, r *http.Request)
	delete(w http.ResponseWriter, r *http.Request)
}

type dummyServiceHandler models.Handler[DummyServiceInterface]

func NewHandler(service *DummyService) (*dummyServiceHandler, error) {
	return &dummyServiceHandler{Service: service}, nil
}

// @Summary     List dummys
// @Description Search dummys by keyword with pagination
// @Tags        dummy
// @Produce     json
// @Param       keyword  query    string  true  "Search keyword"
// @Param       page     query    int     false "Page number"
// @Param       limit    query    int     false "Items per page"
// @Success     200      {object} response.Response{data=[]Dummy}
// @Failure     400      {object} response.Response
// @Router      /dummys/search [get]
func (h *dummyServiceHandler) list(w http.ResponseWriter, r *http.Request) {
	search := SearchQuery{
		Keyword: r.URL.Query().Get("keyword"),
		Page:    r.URL.Query().Get("page"),
		Limit:   r.URL.Query().Get("limit"),
	}

	if err := validate.Struct(search); err != nil {
		response.HandlerResponse(w, r, &search, err)
		return
	}

	result, err := h.Service.List(&search)
	response.HandlerResponse(w, r, &result, err)
}

// @Summary     Get dummy by code
// @Description Get a single dummy record by its code
// @Tags        dummy
// @Accept      json
// @Produce     json
// @Param       body  body     ReadArgs  true  "Dummy code"
// @Success     200   {object} response.Response{data=Dummy}
// @Failure     400   {object} response.Response
// @Router      /dummys/info [post]
func (h *dummyServiceHandler) info(w http.ResponseWriter, r *http.Request) {
	args := ReadArgs{}
	httpRequest.ReadBodyToStruct(r.Body, &args)
	if args.Code == "" {
		response.HandlerResponse(w, r, &args, fmt.Errorf("code is required"))
		return
	}

	result, err := h.Service.Info(&args)
	response.HandlerResponse(w, r, result, err)
}

// @Summary     Create dummy
// @Description Create a new dummy record
// @Tags        dummy
// @Accept      json
// @Produce     json
// @Param       body  body     CreateArgs  true  "Dummy data"
// @Success     200   {object} response.Response{data=Dummy}
// @Failure     400   {object} response.Response
// @Router      /dummys [post]
func (h *dummyServiceHandler) create(w http.ResponseWriter, r *http.Request) {
	args := CreateArgs{}
	httpRequest.ReadBodyToStruct(r.Body, &args)
	if err := validate.Struct(args); err != nil {
		response.HandlerResponse(w, r, &args, err)
		return
	}

	result, err := h.Service.Create(&args)
	response.HandlerResponse(w, r, result, err)
}

// @Summary     Update dummy
// @Description Update an existing dummy record
// @Tags        dummy
// @Accept      json
// @Produce     json
// @Param       body  body     UpdateArgs  true  "Dummy data"
// @Success     200   {object} response.Response{data=Dummy}
// @Failure     400   {object} response.Response
// @Router      /dummys [put]
func (h *dummyServiceHandler) update(w http.ResponseWriter, r *http.Request) {
	args := UpdateArgs{}
	httpRequest.ReadBodyToStruct(r.Body, &args)
	if err := validate.Struct(args); err != nil {
		response.HandlerResponse(w, r, &args, err)
		return
	}

	result, err := h.Service.Update(&args)
	response.HandlerResponse(w, r, result, err)
}

// @Summary     Delete dummy
// @Description Delete a dummy record by code
// @Tags        dummy
// @Accept      json
// @Produce     json
// @Param       body  body     Dummy  true  "Dummy data"
// @Success     200   {object} response.Response
// @Failure     400   {object} response.Response
// @Router      /dummys [delete]
func (h *dummyServiceHandler) delete(w http.ResponseWriter, r *http.Request) {
	dummyObj := Dummy{}
	httpRequest.ReadBodyToStruct(r.Body, &dummyObj)
	if err := validate.Struct(dummyObj); err != nil {
		response.HandlerResponse(w, r, &dummyObj, err)
		return
	}

	err := h.Service.Delete(&DeleteArgs{Dummy: dummyObj})
	response.HandlerResponse(w, r, &dummyObj, err)
}
