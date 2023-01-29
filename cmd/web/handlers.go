package web

import (
	"fmt"
	"gropi/cmd/internal"
	"html/template"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Fatalln(http.StatusMethodNotAllowed)
	}

	// if r.URL.Path != "/" {
	// 	log.Fatalln("Url Error")
	// }

	files := []string{
		"./ui/html/home.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		fmt.Println("aaskdmlkas")
		return
	}

	list := internal.GetArtists()
	if err = ts.Execute(w, list); err != nil {
		fmt.Println("Execute ")
		return
	}
}
