package cmd

import (
	"flag"
	"fmt"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
	"sbb-golang-template/pkg/config"
	"sbb-golang-template/pkg/logger"
	"sbb-golang-template/pkg/routes"
)

func InitCmd(initFn func() error) error {
	env := flag.String("env", "local", "environment (local, dev, production)")
	loadEnvFile := flag.Bool("loadEnvFile", true, "load .env file")
	loadConfFile := flag.Bool("loadConfFile", false, "load config JSON file")
	flag.Parse()

	if err := config.Init(*env, *loadEnvFile, *loadConfFile); err != nil {
		return fmt.Errorf("config init failed: %w", err)
	}

	routes.InitRoutes()

	if err := initFn(); err != nil {
		return fmt.Errorf("module init failed: %w", err)
	}

	// Swagger UI
	mux := routes.GetRouteService().GetMux().(*http.ServeMux)
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	cfg := config.GetConfig()
	port := cfg.Server.Port
	if port == 0 {
		port = 9090
	}

	addr := fmt.Sprintf(":%d", port)
	logger.Get().Infof("Starting server on %s (env=%s) — Swagger: http://localhost%s/swagger/", addr, *env, addr)

	return http.ListenAndServe(addr, withMiddleware(routes.GetRouteService().GetMux()))
}

func withMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg := config.GetConfig()

		if cfg.Server.Cors != nil {
			origins := cfg.Server.Cors.AllowOrigin
			if len(origins) > 0 && origins[0] == "*" {
				w.Header().Set("Access-Control-Allow-Origin", "*")
			} else {
				origin := r.Header.Get("Origin")
				for _, o := range origins {
					if o == origin {
						w.Header().Set("Access-Control-Allow-Origin", origin)
						break
					}
				}
			}
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		}

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
