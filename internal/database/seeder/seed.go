package seeder

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	articleModels "web/internal/modules/article/models"
	userModels "web/internal/modules/user/models"
	"web/pkg/database"
)

func Seed() {
	db := database.Connection()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
		return
	}
	user := userModels.User{Name: "Random name", Email: "random@example.com", Password: string(hashedPassword)}
	db.Create(&user)
	fmt.Printf("User successfully created with email %s and id %d\n", user.Email, user.ID)
	for i := 0; i < 10; i++ {
		article := articleModels.Article{
			Title:   fmt.Sprintf("Random title %d", i),
			Content: fmt.Sprintf("Random content %d", i),
			UserID:  user.ID}
		db.Create(&article)
		fmt.Printf("article successfully created with id %d\n", article.ID)
	}
	log.Println("Seeder Done.")
}
