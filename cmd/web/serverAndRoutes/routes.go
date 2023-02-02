package web

import (
	handlers "gropi/cmd/web/handlers"
	"net/http"
)

func (r *Router) routes() *http.ServeMux {
	r.mux = http.NewServeMux()
	r.mux.HandleFunc("/", handlers.Home)
	r.mux.HandleFunc("/artist/", handlers.ArtistPage)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	r.mux.Handle("/static", http.NotFoundHandler())
	r.mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	return r.mux
}
