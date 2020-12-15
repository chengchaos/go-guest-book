package dao

import (
	"testing"

	"github.com/chengchao/go-guest-book/entities"
)

func TestGetArticleById(t *testing.T) {

	id := 2
	articleDao := NewArticleDao()
	result := articleDao.GetArticleById(id)

	t.Log("result => ", result)
	t.Log("title => ", result.Title)
	t.Log("content => ", result.Content)
	t.Log("createAt => ", result.CreateAt)

}

func TestSaveArticle(t *testing.T) {
	art1 := &entities.Article{
		UserID:  1,
		Title:   "测试文档标题",
		Content: "测试文档内容不为空",
	}
	articleDao := NewArticleDao()
	articleDao.SaveArticle(art1)

	t.Log("art1 id : ", art1.ID)
}

func TestUpdateArticle(t *testing.T) {

	articleDao := NewArticleDao()

	art1 := articleDao.GetArticleById(1)
	art1.Title = "测试修改文档标题"
	art1.Content = "而那些请求都是无意间说出的。"

	articleDao.UpdateArticle(art1)

	t.Log("art1 id : ", art1.ID)
}

func TestDeleteArticle(t *testing.T) {

	articleDao := NewArticleDao()
	articleDao.DeleteArticleById(3)

}

func TestListArticles(t *testing.T) {

	articleDao := NewArticleDao()

	articles := articleDao.ListArticles(1)

	for i, art := range articles {
		t.Log(i, ", ID => ", art.ID)
	}
}
