package main

import (
    "fmt"
    "time"
    "encoding/json"

    "github.com/boltdb/bolt"
)

var db *bolt.DB
var open bool

func OpenDB(dbfile string) (*bolt.DB, error) {
    var err error
    config := &bolt.Options{Timeout: 1 * time.Second}
    db, err = bolt.Open(dbfile, 0600, config)
    if err != nil {
        return nil, fmt.Errorf("couldn't open %s: %s", dbfile, err)
    }
    open = true
    return db, nil
}

func CloseDB() {
    if db != nil {
        db.Close()
    }
    open = false
}

type Entry struct {
    Timestamp   time.Time
    Name        string
    Message     string
}

func (e *Entry) Date() string {
    return e.Timestamp.Format(time.RFC822)
}

func (e *Entry) save() error {
    if !open {
        return fmt.Errorf("db must be opened before saving!")
    }
    err := db.Update(func(tx *bolt.Tx) error {
        entries, err := tx.CreateBucketIfNotExists([]byte("entries"))
        if err != nil {
            return fmt.Errorf("problem creating bucket: %s", err)
        }
        enc, err := e.encode()
        if err != nil {
            return fmt.Errorf("problem encoding entry for %s: %v", e.Name, err)
        }
        ts, err := e.Timestamp.MarshalBinary()
        if err != nil {
            return fmt.Errorf("problem converting timestamp to bytes %s: %v", 
                e.Timestamp, 
                err)
        }
        err = entries.Put(ts, enc)
        return err
    })
    return err
}

func (e *Entry) encode() ([]byte, error) {
    enc, err := json.Marshal(e)
    if err != nil {
        return nil, err
    }
    return enc, nil
}

func decode(data []byte) (*Entry, error) {
    var e *Entry
    err := json.Unmarshal(data, &e)
    if err != nil {
        return nil, err
    }
    return e, nil
}

func Entries() ([]*Entry, error) {
    var entries []*Entry
    err := db.View(func(tx *bolt.Tx) error {
        c := tx.Bucket([]byte("entries")).Cursor()
        for k, v := c.First(); k != nil; k, v = c.Next() {
            // fmt.Printf("key=%s, value=%s\n", k, v)
            dec, err := decode(v)
            if err != nil {
                return fmt.Errorf("problem decoding %s: %s", k, err)
            }
            entries = append(entries, dec)
        }
        return nil
    })
    return entries, err
}

/*
func main() {
    import "log"

    _, err := OpenDB("data.db")
    if err != nil {
        log.Fatal(err)
    }
    defer CloseDB()

    entries := []*Entry{
        {time.Now().Add(time.Minute * 1), "Bill Joy", "Sun Micro Forever!"},
        {time.Now().Add(time.Minute * 2), "Peter Norvig", "AIMA Forever!"},
        {time.Now().Add(time.Minute * 3), "Donald Knuth", "TAOCP Forever!"},
        {time.Now().Add(time.Minute * 4), "Rob Pike", "CSP Forever!"},
        {time.Now().Add(time.Minute * 5), "Brian Kernighan", "Long live C!"},
        {time.Now().Add(time.Minute * 6), "Ken Thompson", "Long live Unix!"},
    }

    // Persist entries in the database.
    for _, e := range entries {
        e.save()
    }

    entries, err := Entries()
    if err != nil {
        log.Fatal(err)
    }
    for _, e := range entries {
        fmt.Println(e)
    }
}
*/
