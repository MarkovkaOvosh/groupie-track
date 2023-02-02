package web

import (
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	httpServer http.Server
}

type Router struct {
	mux *http.ServeMux
}

func Run() {
	server := new(Server)
	router := new(Router)
	if err := server.Serv(":4000", router.routes()); err != nil {
		fmt.Println(err)
		return
	}
}

func (s *Server) Serv(port string, handler http.Handler) error {
	s.httpServer = http.Server{
		Addr:         port,
		Handler:      handler,
		ReadTimeout:  time.Second * 3,
		WriteTimeout: time.Second * 3,
	}

	fmt.Printf("click on the link http://localhost%s/", port)
	return s.httpServer.ListenAndServe()
}
