package dummy

import (
	"net/http"

	"sbb-golang-template/pkg/routes"
)

var baseRoute *routes.RouteGroup

func setBaseRoute(group *routes.RouteGroup) {
	if group != nil {
		baseRoute = group
	} else {
		baseRoute = routes.GetRouteService().GetRouteGroup()
	}
}

func (m *DummyModule) GetRoutes() *[]routes.HttpRoute {
	r := []routes.HttpRoute{
		{Method: http.MethodGet, Path: "/search", Handler: m.Handler.list},
		{Method: http.MethodPost, Path: "/info", Handler: m.Handler.info},
		{Method: http.MethodPost, Path: "", Handler: m.Handler.create},
		{Method: http.MethodPut, Path: "", Handler: m.Handler.update},
		{Method: http.MethodDelete, Path: "", Handler: m.Handler.delete},
	}
	return &r
}

func (m *DummyModule) setRoutes() {
	m.Log.Info("setting " + m.ModuleName + " routes...")
	group := getBaseRoute().Group("/dummys")

	for _, r := range *m.GetRoutes() {
		group.Handle(r.Method, r.Path, r.Handler)
	}
}

func getBaseRoute() *routes.RouteGroup {
	if baseRoute == nil {
		setBaseRoute(nil)
	}
	return baseRoute
}
