package migration

import (
	"sbb-golang-template/pkg/database"
)

func InitialMigration(models ...any) error {
	dbService := database.GetDatabaseService()
	if dbService == nil {
		return nil
	}

	db := dbService.DefaultDatabase()
	if db == nil {
		return nil
	}

	if len(models) > 0 {
		return db.AutoMigrate(models...)
	}
	return nil
}
