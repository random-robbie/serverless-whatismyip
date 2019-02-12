package main

import (
    "fmt"
    "net/http"
    "strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, strings.Split(r.RemoteAddr,":")[0])
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
