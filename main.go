package main

import (
    "fmt"
    "html/template"
    "net/http"
)

func index(w http.ResponseWriter, r *http.Request) {

    tmpl := template.Must(template.ParseFiles("index.html"))

    tmpl.Execute(w, nil)
}

func login(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("login.html"))

    tmpl.Execute(w, nil)
}

func main() {
    http.HandleFunc("/", index)
    http.HandleFunc("/", login)

    fmt.Println("Listening...")
    http.ListenAndServe(":8080", nil)
}
