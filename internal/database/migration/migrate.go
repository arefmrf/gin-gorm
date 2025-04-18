package migration

import (
	"fmt"
	"log"
	articleModels "web/internal/modules/article/models"
	userModels "web/internal/modules/user/models"
	"web/pkg/database"
)

func Migrate() {
	db := database.Connection()
	err := db.AutoMigrate(&userModels.User{}, &articleModels.Article{})

	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("migrate done.")
}
