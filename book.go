package main

import (
    "html/template"
    "net/http"
)

var index = template.Must(template.ParseFiles(
    "templates/base.html",
    "templates/index.html",
))

// book is a HandleFunc for the `/book` route, serving up
// our guestbook web page.
func book(w http.ResponseWriter, r *http.Request) {

    entries, err := Entries()    // get entries from db
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

    if err := index.Execute(w, entries); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
