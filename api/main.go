package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", echoHello)
    fmt.Println("server start ....")
    http.ListenAndServe(":8080", nil)
}

func echoHello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Hello World</h1>")
}