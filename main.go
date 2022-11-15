package main

import (
    "log"
    "net/http"
)

func main() {
     http.Handle("/web/templates/", http.StripPrefix("/web/templates/", http.FileServer(http.Dir("./public"))))
}
