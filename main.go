package main

import (
  "net/http"
  "html/template"
  "github.com/yejin-lim/eks-workshop-sample-api-service-go"
  "fmt"
)

func main(){
    http.Handle("/", http.FileServer(http.Dir("/template")))
    
    fmt.Println("Serving ...")
    http.ListenAndServe(":8080", nil)
}
