package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", HelloServer)
    http.HandleFunc("/200", returnCode200)
    http.HandleFunc("/500", returnCode500)
    http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
 func returnCode200(w http.ResponseWriter, r *http.Request) {
  // see http://golang.org/pkg/net/http/#pkg-constants
  w.WriteHeader(http.StatusOK)
  w.Write([]byte("HTTP status code 200 returned!"))
 }

 func returnCode500(w http.ResponseWriter, r *http.Request) {
  // see http://golang.org/pkg/net/http/#pkg-constants
  w.WriteHeader(http.StatusInternalServerError)
  w.Write([]byte("HTTP status code 500 returned!"))
 }
