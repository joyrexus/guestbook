package main

import (
    "fmt"
    "time"
    "net/http"
)

// sign handles the POST requests coming from the guestbook (`/book`)
// when someone "signs" it, i.e., fills out the guestbook form and
// submits a new entry.  We persist the new entry and redirect 
// back to the guestbook.
func sign(w http.ResponseWriter, r *http.Request) {

    // only accept POST requests
    if r.Method != "POST" {
        http.NotFound(w, r)
        return
    }

    entry := &Entry{
        Timestamp:  time.Now(),
        Name:       r.FormValue("name"),
        Message:    r.FormValue("message"),
    }

    fmt.Printf("Saving entry for %s\n", entry.Name)
    if err := entry.save(); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    http.Redirect(w, r, "/book", http.StatusTemporaryRedirect)
}
