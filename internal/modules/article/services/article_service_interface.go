package services

import ArticleModel "web/internal/modules/article/models"

type ArticleServiceInterface interface {
	GetFeaturedArticles() []ArticleModel.Article
	GetStoriesArticles() []ArticleModel.Article
}
