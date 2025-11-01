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


