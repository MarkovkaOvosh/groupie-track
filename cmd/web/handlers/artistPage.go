package handlers

import (
	"gropi/cmd/internal"
	"gropi/cmd/web/errors"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

type allData struct {
	Art internal.Artist
	Rel internal.Relations
}

func ArtistPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		errors.ErrorPage(w, http.StatusText(http.StatusInternalServerError), http.StatusMethodNotAllowed)
		return
	}

	files := []string{
		"./ui/html/art.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		errors.ErrorPage(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	strUrl := strings.Split(r.URL.Path, "/")
	idArtist, err := strconv.Atoi(strUrl[2])
	if err != nil {
		errors.ErrorPage(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	if idArtist < 1 || idArtist > 52 {
		errors.ErrorPage(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	err = internal.GetAllInfo()
	if err != nil {
		errors.ErrorPage(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	res := oneArtAllData(idArtist)
	res.Art.Data = res.Rel.Data

	if err := ts.Execute(w, res.Art); err != nil {
		errors.ErrorPage(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
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
