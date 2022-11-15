package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParaseFiles("index.html")
	t.ExecuteTemplate(w, "index.html")
func main() {
	http.HandleFunc("/process", process)
	http.ListenAndServe(":8080", nil)
}
