package dummy

import "sbb-golang-template/pkg/migration"

func beginMigration() error {
	return migration.InitialMigration(Dummy{})
}
