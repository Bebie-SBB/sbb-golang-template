package template

import "sbb-golang-template/pkg/migration"

func beginMigration() error {
	return migration.InitialMigration(Template{})
}
