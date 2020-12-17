package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/chengchaos/go-guest-book/services"
)

type ArticleViewHandler struct {
	as services.ArticleService
}

func NewArticlViewHandler(as services.ArticleService) *ArticleViewHandler {
	avh := &ArticleViewHandler{
		as: as,
	}
	return avh
}

func (avh *ArticleViewHandler) ServeHTTP(
	w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")
	longid, err := strconv.Atoi(id)
	if err != nil {
		longid = 1
	}

	art := avh.as.GetArticleById(longid)

	t, err := template.ParseFiles("./templates/article-view.html")
	t.Execute(w, art)
}
