package yinyang

import (
	"html/template"
	"net/http"
)

func Page(w http.ResponseWriter, r *http.Request) {
	artists := Parse()
	if r.URL.Path != "/" {
		Error(w, http.StatusNotFound, " Not correct file location!")
		return
	}
	if r.Method != "GET" {
		Error(w, http.StatusMethodNotAllowed, " Incorrect method!")
		return
	}
	tmpl, err := template.ParseFiles("templates/main.html")
	if err != nil {
		Error(w, http.StatusInternalServerError, " Incorrect page!")
		return
	}
	tmpl.Execute(w, artists)
}
