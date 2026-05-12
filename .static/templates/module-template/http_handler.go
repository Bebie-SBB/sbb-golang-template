package template

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	httpRequest "sbb-golang-template/pkg/http"
	"sbb-golang-template/pkg/models"
	"sbb-golang-template/pkg/response"
)

var validate = validator.New()

type TemplateServiceHandlerInterface interface {
	list(w http.ResponseWriter, r *http.Request)
	info(w http.ResponseWriter, r *http.Request)
	create(w http.ResponseWriter, r *http.Request)
	update(w http.ResponseWriter, r *http.Request)
	delete(w http.ResponseWriter, r *http.Request)
}

type templateServiceHandler models.Handler[TemplateServiceInterface]

func NewHandler(service *TemplateService) (*templateServiceHandler, error) {
	return &templateServiceHandler{Service: service}, nil
}

func (h *templateServiceHandler) list(w http.ResponseWriter, r *http.Request) {
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

func (h *templateServiceHandler) info(w http.ResponseWriter, r *http.Request) {
	args := ReadArgs{}
	httpRequest.ReadBodyToStruct(r.Body, &args)
	if err := validate.Struct(args); err != nil {
		response.HandlerResponse(w, r, &args, err)
		return
	}
	result, err := h.Service.Info(&args)
	response.HandlerResponse(w, r, result, err)
}

func (h *templateServiceHandler) create(w http.ResponseWriter, r *http.Request) {
	args := CreateArgs{}
	httpRequest.ReadBodyToStruct(r.Body, &args)
	if err := validate.Struct(args); err != nil {
		response.HandlerResponse(w, r, &args, err)
		return
	}
	result, err := h.Service.Create(&args)
	response.HandlerResponse(w, r, result, err)
}

func (h *templateServiceHandler) update(w http.ResponseWriter, r *http.Request) {
	args := UpdateArgs{}
	httpRequest.ReadBodyToStruct(r.Body, &args)
	if err := validate.Struct(args); err != nil {
		response.HandlerResponse(w, r, &args, err)
		return
	}
	result, err := h.Service.Update(&args)
	response.HandlerResponse(w, r, result, err)
}

func (h *templateServiceHandler) delete(w http.ResponseWriter, r *http.Request) {
	obj := Template{}
	httpRequest.ReadBodyToStruct(r.Body, &obj)
	if err := validate.Struct(obj); err != nil {
		response.HandlerResponse(w, r, &obj, err)
		return
	}
	err := h.Service.Delete(&DeleteArgs{Template: obj})
	response.HandlerResponse(w, r, &obj, err)
}
