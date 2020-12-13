package server

import (
    "log"
    "net/http"
)

type GuestBookServer struct {
    server *http.Server
}

// AddHandleFuncs to server
func (gbs *GuestBookServer) AddHandleFuncs(
    pattern string, handler func(http.ResponseWriter, *http.Request)) *GuestBookServer {
    http.HandleFunc(pattern, handler)
    return gbs
}

// AddHandler to server
func (gbs *GuestBookServer) AddHandler(
    pattern string, handler http.Handler) *GuestBookServer {
    http.Handle(pattern, handler)
    return gbs
}

func (gbs *GuestBookServer) Start() {
    log.Println("the server is listen and serve ..")
    gbs.server.ListenAndServe()
}



// GetServer is create a http.Server
func NewServer(addr string) *GuestBookServer {
    server := &http.Server{
        Addr: addr,
    }

    gbs := &GuestBookServer{
        server: server,
    }

    return gbs

}
