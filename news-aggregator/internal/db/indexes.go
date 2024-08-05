package db

import (
	"log"
)

func AddIndexes() {
	_, err := db.Exec(`
		CREATE INDEX IF NOT EXISTS idx_articles_title ON articles (title);
		CREATE INDEX IF NOT EXISTS idx_articles_published_at ON articles (published_at);
		CREATE INDEX IF NOT EXISTS idx_recommendations_username ON recommendations (username);
	`)
	if err != nil {
		log.Fatal("Error creating indexes:", err)
	}
}
