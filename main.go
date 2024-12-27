package main

import (
    "fmt"
    "database/sql"
    "log"
    "net/http"
    "os"

    "github.com/seagilbert002/LittleLibrary/handlers"
    "github.com/go-sql-driver/mysql"
)


func main() {
    // initialize Database
    db, err := initializeDB()
    if err != nil {
        log.Fatalf("Could not connect to database: %v", err)
    }
    defer db.Close()

    // Defines roots
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){handlers.IndexHandler(db, w, r)})
    http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request){handlers.BooksHandler(db, w, r)})
    http.HandleFunc("/book/add", func(w http.ResponseWriter, r *http.Request){handlers.AddBookHandler(db, w, r)})

    // Lets the admin know the server is running
    fmt.Println("Server running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

// Initialization creates and returns a new database Connection
func initializeDB() (*sql.DB, error) {
    // Connection properties
    cfg := mysql.Config{
        User:   os.Getenv("DBUSER"),
        Passwd: os.Getenv("DBPASS"),
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "little_library",
    }

    // Database handle
    db, err := sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        return nil, err
    }

    //Double checks that data base is connected.
    pingErr := db.Ping()
    if pingErr != nil {
        return nil, err
    }
    
    return db, nil
}


