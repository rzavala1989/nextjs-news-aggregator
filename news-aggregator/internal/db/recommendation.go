// internal/db/recommendation.go

package db

import (
	"database/sql"
	"fmt"
	"log"
	"news-aggregator/internal/models"
	"os"
)

// db/recommendations.go
func StoreRecommendations(username string, recommendations []models.PyArticle) error {
	db, err := sql.Open("sqlite3", os.Getenv("DATABASE_PATH"))
	if err != nil {
		return err
	}
	defer db.Close()

	for _, rec := range recommendations {
		// Log only key information
		log.Printf("Storing recommendation for user '%s': Title: '%s', Author: '%s'", username, rec.Title, rec.Author)
		_, err = db.Exec(`
            INSERT INTO recommendations (username, title, content, author, url, image_url, language, published_at)
            VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
			username, rec.Title, truncateString(rec.Content, 200), rec.Author, rec.URL, rec.ImageURL, rec.Language, rec.PublishedAt)
		if err != nil {
			log.Println("Error storing recommendation:", err)
			continue
		}
	}
	return nil
}

// truncateString is a helper function to truncate long strings
func truncateString(str string, num int) string {
	if len(str) > num {
		return str[0:num] + "..."
	}
	return str
}

func GetRecommendations(username string, page, limit int) ([]models.PyArticle, error) {
	offset := (page - 1) * limit
	query := `
		SELECT id, title, content, author, url, image_url, language, published_at
		FROM recommendations
		WHERE username = ?
		LIMIT ? OFFSET ?
	`
	rows, err := db.Query(query, username, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("error querying recommendations: %w", err)
	}
	defer rows.Close()

	var recommendations []models.PyArticle
	for rows.Next() {
		var rec models.PyArticle
		if err := rows.Scan(&rec.ID, &rec.Title, &rec.Content, &rec.Author, &rec.URL, &rec.ImageURL, &rec.Language, &rec.PublishedAt); err != nil {
			return nil, fmt.Errorf("error scanning recommendation: %w", err)
		}
		recommendations = append(recommendations, rec)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return recommendations, nil
}
