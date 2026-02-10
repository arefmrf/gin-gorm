package testutils

import (
	"fmt"
	"os"
	"testing"
	"trip/pkg/database"

	"gorm.io/gorm"
)

func SetupTestDB(t *testing.T, models ...interface{}) *gorm.DB {
	t.Helper()
	os.Setenv("APP_ENV", "test")
	InitTestConfig()
	fmt.Println("==============1")
	//viper.AddConfigPath("config")
	database.Connect()
	db := database.Connection()
	fmt.Println("==============2")

	if len(models) > 0 {
		if err := db.AutoMigrate(models...); err != nil {
			fmt.Println("==============3")
			t.Fatalf("failed to migrate models: %v", err)
		}
	}
	fmt.Println("==============4")
	// Cleanup after test
	tx := db.Begin()
	t.Cleanup(func() {
		tx.Rollback()
	})

	return tx
}
