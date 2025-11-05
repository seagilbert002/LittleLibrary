package main

import (
	"os"
    "log"
	"fmt"
    "net/http"

	"github.com/joho/godotenv"

    "github.com/seagilbert002/LittleLibrary/internal/handlers"
	"github.com/seagilbert002/LittleLibrary/internal/db"
	"github.com/seagilbert002/LittleLibrary/internal/repository"
	"github.com/seagilbert002/LittleLibrary/internal/services"
)


func main() {
    // initialize Database
	if err := godotenv.Load(); err != nil {
		log.Println("WARNING: Could not find .env file. Reading system env variables.")
	}

    db, err := db.InitializeDB()
    if err != nil {
        log.Fatalf("Could not connect to database: %v", err)
    }
    defer db.Close()

	// Wiring up layers
	bookRepo := repository.NewSQLBookRepo(db)

	catalogService := services.NewCatalogService(bookRepo)

	bookHandler := handlers.NewBookHandler(catalogService)

	// Define Routes
    http.HandleFunc("/books", bookHandler.BooksHanlder)

    // Defines roots
//    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){handlers.IndexHandler(db, w, r)})
  //  http.HandleFunc("/book/add", func(w http.ResponseWriter, r *http.Request){handlers.AddBookHandler(db, w, r)})
    //http.HandleFunc("/display_book/", func(w http.ResponseWriter, r *http.Request){handlers.BookDisplayHandler(db, w, r)})

	// RUN SERVER
    // Lets the admin know the server is running
	hostPort := os.Getenv("HOST_PORT")
    fmt.Println("Server running on http://localhost:" + hostPort)
    log.Fatal(http.ListenAndServe(":" + hostPort, nil))
}


