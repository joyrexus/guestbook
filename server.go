package main

import (
	// "fmt"
	"html/template"
	"net/http"
)

var index = template.Must(template.ParseFiles(
	"templates/base.html",
	"templates/index.html",
))

// guestbook handle function
func book(w http.ResponseWriter, r *http.Request) {

	var entries []Entry
	// get entries from db

	if err := index.Execute(w, entries); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/book", book)
	http.HandleFunc("/sign", sign)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
