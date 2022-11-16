package main

import (
  "log"
  "net/http"
  "net/http/cgi"
)

func main() {
  http.Handle("/", &cgi.Handler{
    Path: "/usr/bin/php-cgi",
    Dir:  "./",
    Env: []string{
      "SCRIPT_FILENAME=F2_TBDEV3054.php5",
      "REDIRECT_STATUS=trash",
    },
  })

  // Assumes ../app/app is running in ../app directory
  //http.Handle("/go-fcgi", &fcgi.Handler{
  //  Dialer: &fcgi.NetDialer{
  //    Network: "unix",
  //    Address: "../app.socket",
  //  },
  //})
  //http.Handle("/go-cgi", &cgi.Handler{
  //  Path: "../app/app",
  //  Dir:  "../app/",
  //  Args: []string{"cgi"},
  //})
  //
  //// Assumes ../app/index.py is running in ../app directory
  //http.Handle("/py-fcgi", &fcgi.Handler{
  //  Dialer: &fcgi.NetDialer{
  //    Network: "unix",
  //    Address: "../app-py.socket",
  //  },
  //})
  //http.Handle("/py-cgi", &cgi.Handler{
  //  Path: "../app/index.py",
  //  Dir:  "../app/",
  //  Args: []string{"cgi"},
  //})

  log.Println("port:8080")
  err := http.ListenAndServe(":8080", nil)
  if err != nil {
    log.Fatal(err.Error())
  }
}
