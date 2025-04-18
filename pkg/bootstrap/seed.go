package bootstrap

import (
	"web/internal/database/seeder"
	"web/pkg/config"
	"web/pkg/database"
)

func Seed() {
	config.Set()
	database.Connect()
	seeder.Seed()
}
