package server

import (
	"log"
	"net/http"
)

type GuestBookServer struct {
	server *http.Server
	mux    *http.ServeMux
	addr   string
}

// AddHandleFuncs to server
func (gbs *GuestBookServer) AddHandleFuncs(
	pattern string, handlerFunc func(http.ResponseWriter, *http.Request)) *GuestBookServer {
	// http.HandleFunc(pattern, handler)
	gbs.mux.HandleFunc(pattern, handlerFunc)
	return gbs
}

// AddHandler to server
func (gbs *GuestBookServer) AddHandler(
	pattern string, handler http.Handler) *GuestBookServer {
	//http.Handle(pattern, handler)
	gbs.mux.Handle(pattern, handler)
	return gbs
}

// Start method do start
func (gbs *GuestBookServer) Start() {

	server := &http.Server{
		Addr:    gbs.addr,
		Handler: gbs.mux,
	}
	gbs.server = server

	log.Println("the server is listen and serve ..")
	gbs.server.ListenAndServe()
}

// NewServer is create a http.Server
func NewServer(addr string) *GuestBookServer {
	mux := http.NewServeMux()
	// http.FileServer 函数创建了一个能够为指定目录中的静态文件服务的处理器。
	files := http.FileServer(http.Dir("./public"))
	log.Println("./public => ", http.Dir("./public"))

	// 使用 StripPrefix 函数去除请求中的指定前缀
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	gbs := &GuestBookServer{
		mux:  mux,
		addr: addr,
	}

	return gbs

}

func m2() {
	// mux := http.NewServeMux()
	// files := http.FileServer(http.Dir("/public"))
	// mux.Handle("/static/", http.StripPrefix("/static/", files))
	// mux.HandleFunc("/", index)

	// server := &http.Server{
	// 	Addr:    "127.0.0.1:8080",
	// 	Handler: mux,
	// }
}
