package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

var db *sql.DB

func SetupDatabase() *sql.DB {
	// Use env to get the database file path
	dbPath := os.Getenv("DATABASE_PATH")
	if dbPath == "" {
		dbPath = "./app.db" // Default path if not set
	}
	var err error
	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}

	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);
	`

	createTrendingSearchesTable := `
	CREATE TABLE IF NOT EXISTS trending_searches (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		query TEXT NOT NULL UNIQUE,
		date TEXT NOT NULL
	);
	`

	createRecommendationsTable := `
	CREATE TABLE IF NOT EXISTS recommendations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL,
		title TEXT NOT NULL,
		content TEXT,
		author TEXT,
		url TEXT NOT NULL,
		image_url TEXT,
		language TEXT DEFAULT 'en',
		published_at TEXT NOT NULL,
		FOREIGN KEY (username) REFERENCES users(username),
		UNIQUE(username, title)
	);`

	createArticlesTable := `
	CREATE TABLE IF NOT EXISTS articles (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL UNIQUE,
		content TEXT,
		author TEXT,
		url TEXT NOT NULL,
		image_url TEXT,
		language TEXT,
		published_at TEXT NOT NULL,
		username TEXT NOT NULL
	);
	`

	_, err = db.Exec(createUserTable)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(createTrendingSearchesTable)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(createRecommendationsTable)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(createArticlesTable)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func GetDB() *sql.DB {
	return db
}
