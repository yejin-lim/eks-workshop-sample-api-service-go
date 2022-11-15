package main

import (
	"html/template"
	"net/http"
	"github.com/yejin-lim/eks-workshop-sample-api-service-go"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParaseFiles("index.html")
	t.ExecuteTemplate(w, "index.html")
func main() {
	http.HandleFunc("/process", process)
	http.ListenAndServe(":8080", nil)
}
