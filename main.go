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

    // creates the server
    server := &handlers.Server{DB: db}

    // Defines roots
    http.HandleFunc("/", server.IndexHandler)
    http.HandleFunc("/books", server.BooksHandler)

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


