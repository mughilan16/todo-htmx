package main

import (
	"html/template"
	"log"
	"net/http"
)

type Todo struct {
	Text string
	id   int
}

func main() {
	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		todos := map[string][]Todo{
			"Todos": {
				{Text: "Learn HTMX", id: 1},
				{Text: "Learn Go http", id: 2},
			},
		}
		tmpl.Execute(w, todos)
	}
	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
