package db

import (
	"database/sql"
	"log"
	"time"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func ConnectDB() {
	db, err := sql.Open("sqlite", "internal/db/finance.db")
	if err != nil {
		log.Fatalf("Can't connect to DB: %v", err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(10)
	db.SetConnMaxIdleTime(50 * time.Second)
	db.SetConnMaxLifetime(60 * time.Second)

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping DB: %v", err)
	}

	log.Println("DB connection opened")
	DB = db
}

func CloseDB() {
	if DB != nil {
		DB.Close()
		log.Println("DB connection closed")
	}
}
