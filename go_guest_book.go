package main

import (
	"log"

	"github.com/chengchao/go-guest-book/handlers"
	"github.com/chengchao/go-guest-book/server"
	"github.com/chengchao/go-guest-book/services"
)

func main() {

	log.Println("it Works!")

	as := services.NewArticleService()

	indexHandler := handlers.NewIndexHandler(as)
	articleView := handlers.NewArticlViewHandler(as)
	articleManagerController := handlers.NewArticleManageController(as)
	server.
		NewServer("127.0.0.1:8080").
		AddHandler("/index.asp", indexHandler).
		AddHandler("/article.asp", articleView).
		AddHandleFuncs("/login-ready.asp", handlers.LoginReady).
		AddHandleFuncs("/login-session.asp", handlers.LoginSession).
		AddHandleFuncs("/logout-session.asp", handlers.LogoutSession).
		AddHandleFuncs("/admin/articles.asp", articleManagerController.AdminList()).
		AddHandleFuncs("/admin/add-ready.asp", articleManagerController.AddReady()).
		AddHandleFuncs("/admin/add-save.asp", articleManagerController.AddSave()).
		AddHandleFuncs("/admin/edit-ready.asp", articleManagerController.EditReady()).
		AddHandleFuncs("/admin/edit-save.asp", articleManagerController.EditSave()).
		Start()
}
