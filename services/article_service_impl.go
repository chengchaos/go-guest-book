package services

import (
	"github.com/chengchao/go-guest-book/dao"
	"github.com/chengchao/go-guest-book/entities"
)

type ArticleServiceImpl struct {
	articleDao dao.ArticleDao
	userDao    dao.UserDao
}

func NewArticleService(ad dao.ArticleDao,
	ud dao.UserDao) ArticleService {

	as := &ArticleServiceImpl{
		articleDao: ad,
		userDao:    ud,
	}
	return as
}

func (impl *ArticleServiceImpl) GetArticles(page int) []*entities.Article {

	return impl.articleDao.ListArticles(page)
}

func (impl *ArticleServiceImpl) GetArticleById(articleId int) *entities.Article {
	return impl.articleDao.GetArticleById(articleId)
}

func (impl *ArticleServiceImpl) AddSave(article *entities.Article) {
	impl.articleDao.SaveArticle(article)
}

func (impl *ArticleServiceImpl) EditSave(article *entities.Article) {

	impl.articleDao.UpdateArticle(article)
}
