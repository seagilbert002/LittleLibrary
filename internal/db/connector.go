package db

import (
	"database/sql"
	"os"

	"github.com/go-sql-driver/mysql"
)

// TODO: Add .env variable to pull from
// Initialization creates and returns a new database Connection
func InitializeDB() (*sql.DB, error) {
    // Connection properties
    cfg := mysql.Config{
        User:   os.Getenv("DB_USER"),
        Passwd: os.Getenv("DB_PASSWORD"),
        Net:    "tcp",
        Addr:   os.Getenv("DB_HOST"),
        DBName: os.Getenv("DB_NAME"),
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


