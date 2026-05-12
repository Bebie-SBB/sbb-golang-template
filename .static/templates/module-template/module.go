package template

import (
	"sbb-golang-template/pkg/interfaces"
	"sbb-golang-template/pkg/logger"
	"sbb-golang-template/pkg/models"
)

type TemplateModule struct {
	models.Module[TemplateServiceHandlerInterface, TemplateServiceInterface]
}

var module *TemplateModule

func (m *TemplateModule) SetRoutes() error {
	module.setRoutes()
	return nil
}

func (m *TemplateModule) Init() error {
	m.ModuleName = "templateService"
	module.Log = logger.Get()

	service, err := NewService()
	if err != nil {
		module.Log.Fatal("error initial " + module.ModuleName + " service :" + err.Error())
	}

	handler, err := NewHandler(service)
	if err != nil {
		module.Log.Fatal("error initial " + module.ModuleName + " Handler :" + err.Error())
	}

	module.Handler = handler
	module.Service = service

	if err := beginMigration(); err != nil {
		return err
	}

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
