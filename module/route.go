package module

import (
	"sbb-golang-template/pkg/logger"
	"sbb-golang-template/pkg/routes"
)

func setAllRoutes() error {
	allModules := GetAllModules()
	for _, m := range allModules {
		if err := m.SetRoutes(); err != nil {
			return err
		}
	}
	printAllRoutes()
	return nil
}

func printAllRoutes() {
	allRoutes := routes.GetRouteService().ListRoutes()
	for _, r := range allRoutes {
		logger.Get().Info("route ", r.Method, r.Path)
	}
}
