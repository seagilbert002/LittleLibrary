package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/seagilbert002/LittleLibrary/internal/handlers"

	"github.com/seagilbert002/LittleLibrary/internal/db"
)


func main() {
    // initialize Database
	// TODO: use .env variables for more secure connections
    hostPort := "8080"
    db, err := db.InitializeDB()
    if err != nil {
        log.Fatalf("Could not connect to database: %v", err)
    }
    defer db.Close()

    // Defines roots
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){handlers.IndexHandler(db, w, r)})
    http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request){handlers.BooksHandler(db, w, r)})
    http.HandleFunc("/book/add", func(w http.ResponseWriter, r *http.Request){handlers.AddBookHandler(db, w, r)})
    http.HandleFunc("/display_book/", func(w http.ResponseWriter, r *http.Request){handlers.BookDisplayHandler(db, w, r)})

    // Lets the admin know the server is running
    fmt.Println("Server running on http://localhost:" + hostPort)
    log.Fatal(http.ListenAndServe(":" + hostPort, nil))
}


