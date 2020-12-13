package handlers

import (
    "html/template"
    "log"
    "net/http"
    "time"
)

type LoginController struct {

}


func LoginReady(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("./templates/login-ready.html")
    t.Execute(w, nil)
}

func LoginSession(w http.ResponseWriter, r *http.Request) {

    username := r.FormValue("username")
    password := r.FormValue("password")

    log.Println("username and password => ", username, password)

    theCookie := http.Cookie{
        Name : "go-guest-book",
        Value: "1",
        HttpOnly: true,
    }

    w.Header().Set("Location", "http://localhost:8080/admin/articles.asp")
    //w.Header().Set("Set-Cookie", theCookie.String())
    http.SetCookie(w, &theCookie)
    w.WriteHeader(302)

}

func LogoutSession(w http.ResponseWriter, r *http.Request) {

    theCookie := http.Cookie{
        Name : "go-guest-book",
        Value : "",
        HttpOnly: true,
        Expires: time.Unix(1, 0),
        MaxAge: -1,
    }

    http.SetCookie(w, &theCookie)
    w.Header().Set("Location", "/login-ready.asp")
    w.WriteHeader(302)
}