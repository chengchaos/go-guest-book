package services

import "github.com/chengchaos/go-guest-book/entities"

type ArticleService interface {
	GetArticles(int) []*entities.Article
	GetArticleById(int) *entities.Article
	AddSave(*entities.Article)
	EditSave(*entities.Article)
}
