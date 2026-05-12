package health

import (
	"encoding/json"
	"net/http"

	"sbb-golang-template/pkg/interfaces"
	"sbb-golang-template/pkg/routes"
)

type healthModule struct{}

func (m *healthModule) Init() error {
	return nil
}

func (m *healthModule) SetRoutes() error {
	group := routes.GetRouteService().GetRouteGroup().Group("/health")
	group.Handle(http.MethodGet, "", handleHealth)
	return nil
}

func GetInstance() interfaces.IModule {
	return &healthModule{}
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
