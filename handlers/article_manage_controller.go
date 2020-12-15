package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/chengchao/go-guest-book/entities"
	"github.com/chengchao/go-guest-book/services"
)

type ArticleManageController struct {
	as services.ArticleService
}

func NewArticleManageController(as services.ArticleService) *ArticleManageController {
	amc := &ArticleManageController{
		as: as,
	}
	return amc
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	h := r.Header["Cookie"]
	log.Println("cookie => ", h)
	theCookie, err := r.Cookie("go-guest-book")
	if err != nil {
		log.Println("Connot get the cookie")

		w.Header().Set("Location", "/login-ready.asp")
		w.WriteHeader(302)
		return
	}
	log.Println("theCookie => ", theCookie)
	userId := theCookie.Value
	log.Println("userId => ", userId)

	cs := r.Cookies()
	for i, c := range cs {
		log.Println("Cookie => ", i, c)
	}
}

func myFilter(handlerFunc http.HandlerFunc) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		getCookie(w, r)
		handlerFunc(w, r)
	}
	return fn
}

func (amc *ArticleManageController) AdminList() func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		reqPage := r.FormValue("page")
		page, err := strconv.Atoi(reqPage)
		if err != nil {
			page = 1
		}

		arts := amc.as.GetArticles(page)

		t, err := template.ParseFiles("./templates/admin/articles.html")
		if err != nil {
			log.Println("err => ", err)
		} else {
			t.Execute(w, arts)
		}
	}
	return myFilter(fn)
}

func (amc *ArticleManageController) AddReady() func(w http.ResponseWriter, r *http.Request) {

	fn := func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("./templates/admin/add-ready.html")
		if err != nil {
			log.Println("err => ", err)
		}
		t.Execute(w, nil)
	}

	return myFilter(fn)
}

func (amc *ArticleManageController) EditReady() func(w http.ResponseWriter, r *http.Request) {

	fn := func(w http.ResponseWriter, r *http.Request) {
		articleId := r.FormValue("id")
		log.Println("articleId => ", articleId)
		if articleId != "" {
			id, err := strconv.Atoi(articleId)
			if err != nil {
				log.Println("err => ", err)
			}
			art := amc.as.GetArticleById(id)
			log.Println("art => ", art)
			t, err := template.ParseFiles("./templates/admin/edit-ready.html")
			if err != nil {
				log.Println("err => ", err)
			}
			t.Execute(w, art)
		}
	}

	return myFilter(fn)
}

func (amc *ArticleManageController) AddSave() func(w http.ResponseWriter, r *http.Request) {

	fn := func(w http.ResponseWriter, r *http.Request) {
		title := r.FormValue("title")
		content := r.FormValue("content")

		art := &entities.Article{
			Title:   title,
			Content: content,
			// CreateAt: time.Now(),
		}
		amc.as.AddSave(art)
		w.Header().Set("Location", "../admin/articles.asp")
		w.WriteHeader(302)

	}

	return myFilter(fn)
}

func (amc *ArticleManageController) EditSave() func(w http.ResponseWriter, r *http.Request) {

	fn := func(w http.ResponseWriter, r *http.Request) {
		title := r.FormValue("title")
		content := r.FormValue("content")

		art := &entities.Article{
			Title:    title,
			Content:  content,
			CreateAt: time.Now(),
		}
		amc.as.EditSave(art)
		w.Header().Set("Location", "../admin/articles.asp")
		w.WriteHeader(302)

	}

	return myFilter(fn)
}
