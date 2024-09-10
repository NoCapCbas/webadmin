package data

import (
  "log"
  "database/sql"
)

func SeedDatabase(db *sql.DB) {
    _, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            email VARCHAR(100) NOT NULL
        );
        
        INSERT INTO users (email) VALUES
            ('alice@example.com'),
            ('bob@example.com');
    `)
    if err != nil {
        log.Fatalf("Error seeding database: %v", err)
    }
}
