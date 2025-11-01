package db

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
    // Opening the csv file
    file, err := os.Open("books.csv")
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        return
    }
    defer file.Close();

    // Parsing the CSV file
    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        fmt.Printf("Error reading CSV file: %v\n", err)

        return
    }

    // Connecting to the database

    // Connection properties
    cfg := mysql.Config{
        User:   os.Getenv("DBUSER"),
        Passwd: os.Getenv("DBPASS"),
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "little_library",
    }

    // Database handle
    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }

    // Iterate through the CSV and insert into the database
    for i, record := range records {
        if i == 0 {
            continue
            // skips the header
        }
        //Parsing the record
        title := record[0]
        author := record[1]
        firstName := record[2]
        lastName := record[3]
        genre := record[4]
        series := record[5]
        description := record[6]
        publishDate := record[7]
        publisher := record[8]
        eanIsbn := record[9]
        upcIsbn := record[10]

        pages, err := strconv.Atoi(record[11])
        if err != nil {
            pages = 1
        }

        ddc := record[12]
        coverStyle := record[13]
        sprayedEdges := record[14]
        specialEd := record[15]
        firstEd := record[16]
        signed := record[17]
        location := record[18]

        // Insert the above into the database
        query := "INSERT INTO books (title, author, first_name, last_name, genre, series, description, publish_date, publisher, ean_isbn, upc_isbn, pages, ddc, cover_style, sprayed_edges, special_ed, first_ed, signed, location) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        _, err = db.Exec(query, title, author, firstName, lastName, genre, series, description,publishDate, publisher, eanIsbn, upcIsbn, pages, ddc, coverStyle, sprayedEdges, specialEd, firstEd, signed, location)
        if err != nil {
            fmt.Printf("Error inserting record (%v): %v\n", record, err)
            continue
        }
    }
    fmt.Println("CSV Successfully loaded")
}
