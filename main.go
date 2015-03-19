package main

import "net/http"

func main() {
    OpenDB("./data.db")
    defer CloseDB()

    http.HandleFunc("/book", book)
    http.HandleFunc("/sign", sign)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
