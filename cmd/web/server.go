package web

import (
	"log"
	"net/http"
)

func Serv() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)

	err := http.ListenAndServe(":4000", mux)
	log.Fatalln(err)
}
