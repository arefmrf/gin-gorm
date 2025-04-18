package responses

import (
	"fmt"
	"strconv"
	articleModel "web/internal/modules/article/models"
	UserResponses "web/internal/modules/user/responses"
)

type Article struct {
	ID        string
	Image     string
	Title     string
	Content   string
	CreatedAt string
	User      UserResponses.User
}

type Articles struct {
	Data []Article
}

func ToArticle(article articleModel.Article) Article {
	return Article{
		ID:        strconv.Itoa(int(article.ID)),
		Title:     article.Title,
		Content:   article.Content,
		Image:     "/assets/img/demopic/10.jpg",
		CreatedAt: fmt.Sprintf("%d/%02d/%02d", article.CreatedAt.Year(), article.CreatedAt.Month(), article.CreatedAt.Day()),
		User:      UserResponses.ToUser(article.User),
	}
}

func ToArticles(articles []articleModel.Article) Articles {
	var respons Articles
	for _, article := range articles {
		respons.Data = append(respons.Data, ToArticle(article))
	}
	return respons
}
