package main

import (
    "database/sql"
    "fmt"
    

    _ "github.com/lib/pq"
)

func main() {
    dbURL := fmt.Sprintf("host=localhost port=5432 user=postgres password=Ujjawal@7613 dbname=file_sharing sslmode=disable")

    db, err := sql.Open("postgres", dbURL)
    if err != nil {
        fmt.Println("Database connection failed:", err)
        return
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        fmt.Println("Ping failed:", err)
    } else {
        fmt.Println("Connected to PostgreSQL successfully!")
    }
}
