package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.New("index").Parse("<h1>Hello, Go!<h1>")
	})
	http.ListenAndServe(":8080", nil)
}

type response struct {
	Message string   `json:"message"`
}
