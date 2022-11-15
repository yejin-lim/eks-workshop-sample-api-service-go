package main

import (
  "net/http"
  "html/template"
  "github.com/yejin-lim/eks-workshop-sample-api-service-go"
  "fmt"
)

func main() {
  r := mux.NewRouter()
  tmpl := template.Must(template.ParseFiles("index.html"))

  r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    tmpl.Execute(w, nil)
  }).Methods("GET")
  
  r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    var details ContactDetails
    json.NewDecoder(r.Body).Decode(&details) // read body, then decode
    json.NewEncoder(w).Encode(details) // encode, then send to user
  }).Methods("POST")
  
  http.ListenAndServe(":8080", r)
}
