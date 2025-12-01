package services

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

// InitDatabase initializes a SQLite database connection
func InitDatabase(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}

// CreateProductsTable creates the products table if it doesn't exist
func CreateProductsTable(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS products (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT,
			price REAL NOT NULL,
			stock INTEGER NOT NULL DEFAULT 0
		);
	`

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create products table: %w", err)
	}

	return nil
}
