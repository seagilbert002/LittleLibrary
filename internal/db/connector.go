package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// Initialization creates and returns a new database Connection
func InitializeDB() (*sql.DB, error) {
	// Load in the .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("NOTE: Error loading .env file. Assuming environment variables are set.")
	}

    // Connection properties
    cfg := mysql.Config{
        User:   os.Getenv("DB_USER"),
        Passwd: os.Getenv("DB_PASSWORD"),
        Net:    "tcp",
        Addr:   os.Getenv("DB_HOST"),
        DBName: os.Getenv("DB_NAME"),
		AllowNativePasswords: true,
    }

    // Database handle
    db, err := sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        return nil, err
    }

    //Double checks that data base is connected.
    pingErr := db.Ping()
    if pingErr != nil {
        return nil, pingErr
    }

	log.Println("NOTE: Connected to LittleLibrary db")	
    
    return db, nil
}


