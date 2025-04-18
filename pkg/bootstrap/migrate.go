package bootstrap

import (
	"web/internal/database/migration"
	"web/pkg/config"
	"web/pkg/database"
)

func Migrate() {
	config.Set()
	database.Connect()
	migration.Migrate()
}
