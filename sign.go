package main

import (
	"fmt"
	"time"
	"net/http"
)

type Entry struct {
	Timestamp time.Time
	Name      string
	Message   string
}

func sign(w http.ResponseWriter, r *http.Request) {

	// only except POST requests
	if r.Method != "POST" {
		http.NotFound(w, r)
		return
	}

	entry := &Entry{
		Timestamp: 	time.Now(),
		Name:		r.FormValue("name"),
		Message:	r.FormValue("message"),
	}

	fmt.Println(entry)	
	/* 
	// insert entry into db
	if err := db.Insert(entry); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	*/

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
