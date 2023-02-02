package handlers

import (
	"gropi/cmd/internal"
	"gropi/cmd/web/errors"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errors.ErrorPage(w, http.StatusText(404), 404)

		return
	}

	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		errors.ErrorPage(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)

		return
	}

	files := []string{
		"./ui/html/home.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		errors.ErrorPage(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	err = internal.GetAllInfo()
	if err != nil {

		errors.ErrorPage(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	ts.Execute(w, &internal.List)
}
