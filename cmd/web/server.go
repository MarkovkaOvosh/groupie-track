package web

import (
	"gropi/cmd/internal"
	"log"
	"net/http"
)

func Serv() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/artist/", artPage)

	err := http.ListenAndServe(":4000", mux)
	log.Fatalln(err)
}

type allData struct {
	Art internal.Artist
	Rel internal.Relations
}

func oneArtAllData(id int) allData {
	oneArtistData := internal.List[id-1]
	oneArtistDataRelation := internal.ListRel.Index[id-1]
	res := allData{
		Art: oneArtistData,
		Rel: oneArtistDataRelation,
	}
	return res
}
