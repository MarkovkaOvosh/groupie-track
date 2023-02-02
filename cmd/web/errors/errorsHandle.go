package errors

import (
	"html/template"
	"net/http"
)

type dataErr struct {
	Id   int
	Text string
}

func ErrorPage(w http.ResponseWriter, statusText string, statusID int) {
	temp, err := template.ParseFiles("./ui/html/errors.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	res := dataErr{
		Id:   statusID,
		Text: statusText,
	}
	w.WriteHeader(statusID)
	temp.Execute(w, res)
}
