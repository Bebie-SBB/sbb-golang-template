package template

import (
	"sbb-golang-template/pkg/logger"
	"sbb-golang-template/pkg/models"
)

type TemplateServiceInterface interface {
	Init() error
}

type TemplateService models.Service[any]

func NewService() (*TemplateService, error) {
	s := &TemplateService{}
	return s, s.Init()
}

func (s *TemplateService) Init() error {
	s.Log = logger.Get()
	return nil
}
