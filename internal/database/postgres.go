package database

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
)

var DB *sql.DB

func InitPostgresDB(connStr string) error {
    var err error
    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        return fmt.Errorf("failed to open database connection: %w", err)
    }

    if err = DB.Ping(); err != nil {
        return fmt.Errorf("failed to ping database: %w", err)
    }

    fmt.Println("Successfully connected to PostgreSQL!")
    return nil
}

func ClosePostgresDB() {
    if DB != nil {
        DB.Close()
    }
}