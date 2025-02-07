package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "\"Hello World!\"")
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":7777", nil)
}
