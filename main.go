package main
 
import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "html/template"
)
 
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Bar!")
	})

	http.ListenAndServe(":3000", nil)
}
