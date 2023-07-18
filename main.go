package main

import (
	"log"
	"net/http"
	yinyang "yinyang/functions"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", yinyang.Page)
	mux.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("templates"))))
	log.Println("Server is running on port http://localhost:8081/")
	err := http.ListenAndServe(":8081", mux)
	log.Fatal(err)
}
