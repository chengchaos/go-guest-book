package handlers

import (
    "github.com/chengchao/go-guest-book/services"
    "html/template"
    "log"
    "net/http"
)

type IndexHandler struct {
    articleService services.ArticleService
}

func NewIndexHandler(as services.ArticleService) *IndexHandler{
    ih := &IndexHandler{
        articleService: as,
    }
    return ih
}


func (ih *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    t, err := template.ParseFiles("./templates/index.html")
    if err != nil {
        log.Fatalln(err )
    }
    articles := ih.articleService.GetArticles(1)

    t.Execute(w, articles)
}
