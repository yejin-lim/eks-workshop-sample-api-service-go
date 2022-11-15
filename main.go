package main

import (
    "fmt"
    "html/template"
    "net/http"
)

func main() {
	// public 경로를 serving하는 파일서버
	http.Handle("/", http.FileServer(http.Dir("web/templates/")))
	http.ListenAndServe(":8080", nil)
}
