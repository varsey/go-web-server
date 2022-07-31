package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
)

func main() {

    fmt.Println("Starting server...")

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
        fmt.Println("Hello")
    })

    http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hi")
        fmt.Println("Hi")
    })

    log.Fatal(http.ListenAndServe(":8081", nil))

}

