package news

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"news-aggregator/internal/db"
	"news-aggregator/internal/models"
	"time"
)

func FetchTrendingSearches(apiKey, date string) ([]models.TrendingSearch, error) {
	// Get existing trending searches for the given date
	existingSearches, err := db.GetTrendingSearchesForDate(date)
	if err != nil {
		log.Println("Error fetching existing trending searches:", err)
		return nil, err
	}

	// Convert existing searches to a map for quick lookup
	existingSearchesMap := make(map[string]bool)
	for _, search := range existingSearches {
		existingSearchesMap[search.Query] = true
	}

	url := fmt.Sprintf("https://serpapi.com/search.json?engine=google_trends_trending_now&frequency=daily&date=%s&api_key=%s", date, apiKey)
	log.Println("Fetching trending searches from URL:", url)

	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error making GET request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return nil, err
	}

	log.Println("Response body read successfully")

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Println("Error unmarshalling JSON:", err)
		return nil, err
	}

	log.Println("JSON unmarshalled successfully!")

	var searches []models.TrendingSearch
	if dailySearches, ok := result["daily_searches"].([]interface{}); ok {
		for _, daily := range dailySearches {
			if dailyMap, ok := daily.(map[string]interface{}); ok {
				if searchesList, ok := dailyMap["searches"].([]interface{}); ok {
					for _, search := range searchesList {
						if searchMap, ok := search.(map[string]interface{}); ok {
							query := searchMap["query"].(string)
							// Only append the search to the searches slice and log it if it doesn't already exist in the database
							if !existingSearchesMap[query] {
								searches = append(searches, models.TrendingSearch{
									Query: query,
									Date:  dailyMap["date"].(string),
								})
								log.Println("New trending search:", query)
							}
						}
					}
				}
			}
		}
	}

	log.Println("Trending searches fetched successfully:", searches)
	return searches, nil
}

func UpdateTrendingSearches(apiKey string) {
	date := time.Now().Format("20060102")
	trendingSearches, err := FetchTrendingSearches(apiKey, date)
	if err != nil {
		log.Println("Error fetching trending searches:", err)
		return
	}

	// Get existing trending searches for today
	existingSearches, err := db.GetTrendingSearchesForDate(date)
	if err != nil {
		log.Println("Error fetching existing trending searches:", err)
		return
	}

	// Convert existing searches to a map for quick lookup
	existingSearchesMap := make(map[string]bool)
	for _, search := range existingSearches {
		existingSearchesMap[search.Query] = true
	}

	for _, search := range trendingSearches {
		// Only log the search if it doesn't already exist in the database
		if !existingSearchesMap[search.Query] {
			log.Println("New trending search:", search.Query)
		}
	}

	if len(existingSearches) == 0 {
		log.Println("No existing trending searches for today. Inserting new ones.")
		err = db.StoreTrendingSearches(trendingSearches)
		if err != nil {
			log.Println("Error storing trending searches:", err)
		}
	} else {
		log.Println("Trending searches for today already exist.")
	}
}

func GetTrendingSearchesFromDB() ([]models.TrendingSearch, error) {
	log.Println("Fetching trending searches from database")
	return db.GetTrendingSearches()
}
