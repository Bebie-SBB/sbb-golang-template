package template

import (
	"sbb-golang-template/pkg/interfaces"
	"sbb-golang-template/pkg/logger"
	"sbb-golang-template/pkg/models"
)

type TemplateModule struct {
	models.Module[any, TemplateServiceInterface]
}

var module *TemplateModule

func (m *TemplateModule) SetRoutes() error {
	return nil
}

func (m *TemplateModule) Init() error {
	m.ModuleName = "template"

	service, err := NewService()
	if err != nil {
		module.Log.Fatal("error initial " + module.ModuleName + " service :" + err.Error())
	}

	module.Service = service
	module.Log = logger.Get()

	return nil
}

func GetInstance() interfaces.IModule {
	if module == nil {
		module = &TemplateModule{}
		if err := module.Init(); err != nil {
			module.Log.Fatal("error initial " + module.ModuleName + " module :" + err.Error())
		}
	}
	return module
}
