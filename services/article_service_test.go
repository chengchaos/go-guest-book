package services

import (
	"testing"

	"github.com/chengchao/go-guest-book/dao"
)

func TestGetArticleById(t *testing.T) {
	id := 1
	articleDao := dao.NewArticleDao()
	userDao := dao.NewUserDao()

	articleService := NewArticleService(articleDao, userDao)

	article := articleService.GetArticleById(id)
	t.Log(article.Title)

}
