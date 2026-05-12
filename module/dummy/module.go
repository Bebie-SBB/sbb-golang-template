package dummy

import (
	"sbb-golang-template/pkg/interfaces"
	"sbb-golang-template/pkg/logger"
	"sbb-golang-template/pkg/models"
)

type DummyModule struct {
	models.Module[DummyServiceHandlerInterface, DummyServiceInterface]
}

var module *DummyModule

func (m *DummyModule) SetRoutes() error {
	module.setRoutes()
	return nil
}

func (m *DummyModule) Init() error {
	m.ModuleName = "dummyService"
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
		module = &DummyModule{}

		if err := module.Init(); err != nil {
			module.Log.Fatal("error initial " + module.ModuleName + " module :" + err.Error())
		}
	}
	return module
}
