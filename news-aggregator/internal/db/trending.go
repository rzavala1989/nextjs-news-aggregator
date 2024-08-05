package db

import (
	"log"
	"news-aggregator/internal/models"
	"time"
)

func StoreTrendingSearches(searches []models.TrendingSearch) error {
	log.Println("Storing trending searches in the database")
	for _, search := range searches {
		_, err := db.Exec("INSERT INTO trending_searches (query, date) VALUES (?, ?) ON CONFLICT(query) DO UPDATE SET date=?", search.Query, search.Date, search.Date)
		if err != nil {
			log.Println("Error storing trending search:", search.Query, err)
			return err
		}
	}
	return nil
}

func GetTrendingSearchesForDate(date string) ([]models.TrendingSearch, error) {
	rows, err := db.Query("SELECT id, query, date FROM trending_searches WHERE date = ?", date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var trendingSearches []models.TrendingSearch
	for rows.Next() {
		var search models.TrendingSearch
		if err := rows.Scan(&search.ID, &search.Query, &search.Date); err != nil {
			return nil, err
		}
		trendingSearches = append(trendingSearches, search)
	}

	log.Println("Trending searches from DB:", trendingSearches)
	return trendingSearches, nil
}

func GetTrendingSearches() ([]models.TrendingSearch, error) {
	log.Println("Fetching trending searches from the database")
	rows, err := db.Query("SELECT query, date FROM trending_searches WHERE date=?", time.Now().Format("20060102"))
	if err != nil {
		log.Println("Error fetching trending searches:", err)
		return nil, err
	}
	defer rows.Close()

	var searches []models.TrendingSearch
	for rows.Next() {
		var search models.TrendingSearch
		if err := rows.Scan(&search.Query, &search.Date); err != nil {
			log.Println("Error scanning trending search row:", err)
			return nil, err
		}
		searches = append(searches, search)
	}
	return searches, nil
}
