package yinyang

import (
	"html/template"
	"net/http"
)

func Error(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	response := struct {
		ErrorCode int
		ErrorText string
	}{
		ErrorCode: code,
		ErrorText: message,
	}
	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, response)
}
