package services

import (
	ArticleModel "web/internal/modules/article/models"
	ArticleRepository "web/internal/modules/article/repositories"
)

type ArticleService struct {
	articleRepository ArticleRepository.ArticleRepositoryInterface
}

func New() *ArticleService {
	return &ArticleService{
		articleRepository: ArticleRepository.New(),
	}
}

func (ArticleService *ArticleService) GetFeaturedArticles() []ArticleModel.Article {
	return ArticleService.articleRepository.List(4)
}

func (ArticleService *ArticleService) GetStoriesArticles() []ArticleModel.Article {
	return ArticleService.articleRepository.List(6)
}
