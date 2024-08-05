package news

import (
	"log"
	"news-aggregator/internal/db"
	"time"
)

// ArchiveOldArticles archives articles older than 30 days
func ArchiveOldArticles() {
	for {
		log.Println("Archiving old articles...")
		_, err := db.GetDB().Exec("DELETE FROM articles WHERE published_at < date('now', '-30 days')")
		if err != nil {
			log.Println("Error archiving old articles:", err)
		}
		time.Sleep(24 * time.Hour) // Run the archive job daily
	}
}
