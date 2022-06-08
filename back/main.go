package main

import (
	"net/http"
	"text/template"
)

func main() {
	tmpl := template.Must(template.ParseFiles("../dist/index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		tmpl.Execute(w, nil)
	})

	//static assets
	fs := http.FileServer(http.Dir("../dist/assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}
