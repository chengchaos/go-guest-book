package services

import (
    "github.com/chengchao/go-guest-book/entities"
    "log"
    "strconv"
)

type ArticleServiceImpl struct {

}

func NewArticleService() ArticleService {
    as := &ArticleServiceImpl{}
    return as
}


func (asi *ArticleServiceImpl) GetArticles(page int) []*entities.Article {

    res := make([]*entities.Article, 10)
    for i := 0; i < 10; i++ {
        a:= &entities.Article{
            Id : int64(i),
            Title: "Title "+ strconv.Itoa(i),
            Content: "<div>Content .... "+ strconv.Itoa(i),
        }

        res[i] = a

    }
    return res
}

func (asi *ArticleServiceImpl) GetArticleById(articleId int) *entities.Article {

    art := &entities.Article{
        Id: int64(articleId),
        Title: "又是一个星期天",
        Content: "<div>这让我如何是好？</div>",
    }

    return art
}

func (asi *ArticleServiceImpl) AddSave(article *entities.Article) {
    log.Println("do save")
}

func (asi *ArticleServiceImpl) EditSave(article *entities.Article) {
    log.Println("do save")
}
