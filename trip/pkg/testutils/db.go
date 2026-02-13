package testutils

import (
	"os"
	"testing"
	"trip/pkg/database"

	"gorm.io/gorm"
)

func SetupTestDB(t *testing.T, models ...interface{}) *gorm.DB {
	t.Helper()
	os.Setenv("APP_ENV", "test")
	InitTestConfig()
	//viper.AddConfigPath("config")
	database.Connect()
	db := database.Connection()

	if len(models) > 0 {
		if err := db.AutoMigrate(models...); err != nil {
			t.Fatalf("failed to migrate models: %v", err)
		}
	}
	// Cleanup after test
	tx := db.Begin()
	t.Cleanup(func() {
		tx.Rollback()
	})

	return tx
}
