package database

import (
	"fmt"
	"log"
	"os"
	config2 "trip/config"
	"trip/pkg/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() {
	cfg := config.Get()
	dbCfg := cfg.DB
	if os.Getenv("APP_ENV") == "test" {
		dbCfg = config2.DB{
			Username: cfg.TestDB.Username,
			Password: cfg.TestDB.Password,
			Host:     cfg.TestDB.Host,
			Port:     cfg.TestDB.Port,
			Name:     cfg.TestDB.Name,
		}
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbCfg.Host, dbCfg.Username, dbCfg.Password, dbCfg.Name, dbCfg.Port,
	)
	fmt.Println()
	fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return
	}
	DB = db
}
