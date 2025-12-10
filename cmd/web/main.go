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
	// ****** Database Connections ******
	// Loads in the environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("WARNING: Could not find .env file. Reading system env variables.")
	}

	// Starts the database connection
    db, err := db.InitializeDB()
    if err != nil {
        log.Fatalf("Could not connect to database: %v", err)
    }
	log.Println("Database Connection Successful")
    defer db.Close()


	// ****** Wiring up layers *****
	// ----Repositories----
	bookRepo := repository.NewSQLBookRepo(db)

	// ----Services----
	catalogService := services.NewCatalogService(bookRepo)

	// ----Handlers----
	genralHandler := handlers.NewGeneralHandler()
	bookHandler := handlers.NewBookHandler(catalogService)

	// ***** Define Routes ******
	http.HandleFunc("/", genralHandler.IndexHandler)
	// Book Routes
    http.HandleFunc("/books", bookHandler.BooksHanlder)
	http.HandleFunc("/display_book/", bookHandler.BookDisplayHandler)
	http.HandleFunc("/add_book", bookHandler.AddBookHandler)
	http.HandleFunc("/remove_book/", bookHandler.RemoveBookHandler)

	// RUN SERVER
    // Lets the admin know the server is running
	hostPort := os.Getenv("HOST_PORT")
    fmt.Println("Server running on http://localhost:" + hostPort)
    log.Fatal(http.ListenAndServe(":" + hostPort, nil))
}


