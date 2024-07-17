package database

import (
	"context"
	"database/sql"
	"fmt"

	// Register the SQLite3 driver
	_ "github.com/mattn/go-sqlite3"
)

// Initialize sets up the database connection and creates necessary tables
func Initialize() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./company.db")
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	createTables := `
    CREATE TABLE IF NOT EXISTS users (
        user_id TEXT PRIMARY KEY,
        username TEXT UNIQUE,
        password TEXT,
        firstname TEXT,
        lastname TEXT
    );
    CREATE TABLE IF NOT EXISTS company (
      	CompanyID TEXT PRIMARY KEY,
		name TEXT NOT NULL UNIQUE,
		description TEXT,
		amountofemployees INTEGER NOT NULL,
		registered BOOLEAN NOT NULL,
		type TEXT NOT NULL
    );`

	_, err = db.ExecContext(context.Background(), createTables)
	if err != nil {
		return nil, fmt.Errorf("failed to create tables: %w", err)
	}

	return db, nil
}
