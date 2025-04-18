package repositories

import (
	ArticleModel "web/internal/modules/article/models"
)

type ArticleRepositoryInterface interface {
	List(limit int) []ArticleModel.Article
}
