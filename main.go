// @title           SBB Golang Template API
// @version         1.0
// @description     Go REST API template with modular structure.
// @host            localhost:9090
// @BasePath        /
// @schemes         http https

package main

import (
	"log"

	_ "sbb-golang-template/docs"
	"sbb-golang-template/module"
	"sbb-golang-template/pkg/cmd"
	"sbb-golang-template/pkg/database"
	"sbb-golang-template/pkg/health"
)

func main() {
	if err := cmd.InitCmd(InitModules); err != nil {
		log.Fatalln(err)
	}
}

func InitModules() error {
	if err := module.PrependModules(database.GetInstance(), health.GetInstance()); err != nil {
		return err
	}
	return module.Init(&module.ModuleConfig{})
}
