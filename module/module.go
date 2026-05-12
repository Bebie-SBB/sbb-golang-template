package module

import (
	"errors"

	"sbb-golang-template/module/dummy"
	"sbb-golang-template/pkg/interfaces"
)

var modules []interfaces.IModule
var moduleConfig *ModuleConfig

func init() {
	modules = []interfaces.IModule{}
}

type ModuleConfig struct{}

func Init(cfg *ModuleConfig) error {
	if cfg == nil {
		return errors.New("no config provided")
	}

	// add modules here
	modules = append(modules, dummy.GetInstance())

	return setAllRoutes()
}

func GetAllModules() []interfaces.IModule {
	return modules
}

func GetModuleConfig() *ModuleConfig {
	return moduleConfig
}

func AppendModules(moduleArr ...interfaces.IModule) error {
	if len(moduleArr) == 0 {
		return errors.New("no module to add")
	}
	modules = append(modules, moduleArr...)
	return nil
}

func PrependModules(moduleArr ...interfaces.IModule) error {
	if len(moduleArr) == 0 {
		return errors.New("no module to add")
	}
	modules = append(moduleArr, modules...)
	return nil
}
