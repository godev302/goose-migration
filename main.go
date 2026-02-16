package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

const (
	dialect       = "postgres"
	defaultDSN    = "host=localhost port=5433 user=postgres password=postgres dbname=postgres sslmode=disable"
	migrationsDir = "migrations"
)

func main() {
	// 1. Open database connection
	db, err := sql.Open(dialect, defaultDSN)
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}
	defer db.Close()

	// 2. Set goose dialect
	if err := goose.SetDialect(dialect); err != nil {
		log.Fatalf("failed to set dialect: %v", err)
	}

	// 3. Uncomment the function you want to run:

	//showStatus(db)
	migrateUp(db)
	// migrateDown(db)
	// migrateUpTo(db, 2)
	//migrateDownTo(db, 0) // Rolls back all migrations
}

func migrateUp(db *sql.DB) {
	log.Println("Running: goose up")
	if err := goose.Up(db, migrationsDir); err != nil {
		log.Fatalf("goose up: %v", err)
	}
}

func migrateDown(db *sql.DB) {
	log.Println("Running: goose down")
	if err := goose.Down(db, migrationsDir); err != nil {
		log.Fatalf("goose down: %v", err)
	}
}

func migrateUpTo(db *sql.DB, version int64) {
	log.Printf("Running: goose up-to %d\n", version)
	if err := goose.UpTo(db, migrationsDir, version); err != nil {
		log.Fatalf("goose up-to: %v", err)
	}
}

func migrateDownTo(db *sql.DB, version int64) {
	log.Printf("Running: goose down-to %d\n", version)
	if err := goose.DownTo(db, migrationsDir, version); err != nil {
		log.Fatalf("goose down-to: %v", err)
	}
}

func showStatus(db *sql.DB) {
	log.Println("Running: goose status")
	if err := goose.Status(db, migrationsDir); err != nil {
		log.Fatalf("goose status: %v", err)
	}
}
