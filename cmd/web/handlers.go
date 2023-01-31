package web

import (
	"fmt"
	"gropi/cmd/internal"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Fatalln(http.StatusMethodNotAllowed)
	}

	files := []string{
		"./ui/html/home.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = internal.GetAllInfo()
	if err != nil {
		return
	}

	if err := ts.Execute(w, &internal.List); err != nil {
		fmt.Println(err)
		return
	}
}

func artPage(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/art.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		fmt.Println(err)
		return
	}

	strUrl := strings.Split(r.URL.Path, "/")
	idArtist, err := strconv.Atoi(strUrl[2])
	if err != nil {
		return
	}

	err = internal.GetAllInfo()
	if err != nil {
		return
	}

	res := oneArtAllData(idArtist)
	res.Art.Data = res.Rel.Data

	if err := ts.Execute(w, res.Art); err != nil {
		fmt.Println(err)
		return
	}
}
