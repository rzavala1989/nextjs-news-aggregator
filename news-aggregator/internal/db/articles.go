// internal/db/articles.go

package db

import (
	"news-aggregator/internal/models"
)

func StoreArticles(articles []models.Article) error {
	for _, article := range articles {
		_, err := db.Exec("INSERT INTO articles (title, content, author, url, image_url, language, published_at) VALUES (?, ?, ?, ?, ?, ?, ?)",
			article.Title, article.Content, article.Author, article.URL, article.ImageURL, article.Language, article.PublishedAt)
		if err != nil {
			return err
		}
	}
	return nil
}
