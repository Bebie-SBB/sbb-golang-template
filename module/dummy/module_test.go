package dummy

import (
	"testing"

	"sbb-golang-template/pkg/config"
	"sbb-golang-template/pkg/database"
	"sbb-golang-template/pkg/routes"
)

func TestDummyModule(t *testing.T) {
	testConfigFile(t)
	mod := testCreateModule(t)
	testInitModule(t, mod)
	testModuleSetRoute(t, mod)
}

func testCreateModule(t *testing.T) *DummyModule {
	testConfigFile(t)
	database.GetInstance()
	routes.InitRoutes()

	modInterface := GetInstance()
	mod, ok := modInterface.(*DummyModule)
	if !ok {
		t.Fatal("cannot cast IModule to DummyModule")
	}
	if err := mod.Init(); err != nil {
		t.Error(err)
	}
	return mod
}

func testInitModule(t *testing.T, mod *DummyModule) {
	if err := mod.Init(); err != nil {
		t.Error(err)
	}
}

func testModuleSetRoute(t *testing.T, mod *DummyModule) {
	routes.InitRoutes()
	group := routes.GetRouteService().GetRouteGroup().Group("dummy-test")
	setBaseRoute(group)
	if err := mod.SetRoutes(); err != nil {
		t.Error(err)
	}
}

func testConfigFile(t *testing.T) {
	testConfig := config.ConfigFile{
		Databases: []config.Database{
			{
				Name: "sqlite",
				Config: config.DatabaseConfig{
					Dialect: "sqlite",
				},
			},
		},
		Server: config.ServerConfig{
			Name: "dummy-service",
			Cors: &config.CorsConfig{
				AllowHeaders: []string{"*"},
				AllowOrigin:  []string{"*"},
			},
		},
		Configs: config.Configs{
			DatabaseList:            "sqlite",
			IsMigration:             true,
			TargetMigrationDatabase: "sqlite",
			DisplayResponseError:    true,
		},
	}
	config.SetConfig(&testConfig)
}
