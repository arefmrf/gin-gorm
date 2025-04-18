package repositories

import (
	"gorm.io/gorm"
	articleModels "web/internal/modules/article/models"
	"web/pkg/database"
)

type ArticleRepository struct {
	DB *gorm.DB
}

func New() *ArticleRepository {
	return &ArticleRepository{
		DB: database.Connection(),
	}
}

func (articleRepository *ArticleRepository) List(limit int) []articleModels.Article {
	var articles []articleModels.Article
	articleRepository.DB.Limit(limit).Joins("User").Find(&articles)
	return articles
}
