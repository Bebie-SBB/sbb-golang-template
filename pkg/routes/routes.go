package routes

import (
	"fmt"
	"net/http"
)

type HttpRoute struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

type RouteEntry struct {
	Method string
	Path   string
}

type RouteGroup struct {
	prefix string
	mux    *http.ServeMux
}

type RouteService struct {
	mux    *http.ServeMux
	routes []RouteEntry
	root   *RouteGroup
}

var routeService *RouteService

func InitRoutes() {
	mux := http.NewServeMux()
	routeService = &RouteService{
		mux:  mux,
		root: &RouteGroup{mux: mux},
	}
}

func GetRouteService() *RouteService {
	return routeService
}

func (rs *RouteService) GetRouteGroup() *RouteGroup {
	return rs.root
}

func (rs *RouteService) ListRoutes() []RouteEntry {
	return rs.routes
}

func (rs *RouteService) GetMux() http.Handler {
	return rs.mux
}

func (g *RouteGroup) Group(path string) *RouteGroup {
	return &RouteGroup{
		prefix: g.prefix + path,
		mux:    g.mux,
	}
}

// Handle registers a route using Go 1.22+ "METHOD /path" pattern syntax.
func (g *RouteGroup) Handle(method, path string, handler http.HandlerFunc) {
	fullPath := g.prefix + path
	pattern := fmt.Sprintf("%s %s", method, fullPath)
	g.mux.HandleFunc(pattern, handler)

	if routeService != nil {
		routeService.routes = append(routeService.routes, RouteEntry{
			Method: method,
			Path:   fullPath,
		})
	}
}
