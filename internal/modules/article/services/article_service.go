package services

import (
	ArticleRepository "web/internal/modules/article/repositories"
	ArticleResponse "web/internal/modules/article/responses"
)

type ArticleService struct {
	articleRepository ArticleRepository.ArticleRepositoryInterface
}

func New() *ArticleService {
	return &ArticleService{
		articleRepository: ArticleRepository.New(),
	}
}

func (ArticleService *ArticleService) GetFeaturedArticles() ArticleResponse.Articles {
	articles := ArticleService.articleRepository.List(4)
	return ArticleResponse.ToArticles(articles)
}

func (ArticleService *ArticleService) GetStoriesArticles() ArticleResponse.Articles {
	articles := ArticleService.articleRepository.List(6)
	return ArticleResponse.ToArticles(articles)
}
