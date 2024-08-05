package news

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"news-aggregator/internal/db"
	"news-aggregator/internal/models"
)

func SyncArticlesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Manual sync triggered")
	SyncArticlesAndUpdateRecommendations()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Articles and recommendations synced successfully"))
}

func SyncArticlesAndUpdateRecommendations() {
	articles, err := fetchAndConvertAllArticles()
	if err != nil {
		log.Println("Error fetching articles:", err)
		return
	}

	pyArticles := convertToPyArticles(articles)

	err = syncWithPythonService(pyArticles)
	if err != nil {
		log.Println("Error syncing articles with Python service:", err)
		return
	}

	err = storeRecommendationsForAllUsers(pyArticles)
	if err != nil {
		log.Println("Error storing recommendations:", err)
	}
}

func fetchAndConvertAllArticles() ([]models.Article, error) {
	apiURLs := []string{
		"https://api.currentsapi.services/v1/search?keywords=news&language=en&apiKey=" + os.Getenv("CURRENTS_API_KEY"),
		"https://api.thenewsapi.com/v1/news/top?api_token=" + os.Getenv("THENEWSAPI_API_KEY") + "&locale=us&limit=5",
		"https://newsdata.io/api/1/latest?apikey=" + os.Getenv("NEWSDATAIO_API_KEY") + "&q=news",
		"https://api.mediastack.com/v1/news?access_key=" + os.Getenv("MEDIASTACK_API_KEY") + "&keywords=news",
	}

	var wg sync.WaitGroup
	ch := make(chan []models.Article)
	username := "system" // Use a default username for system-fetch

	for _, url := range apiURLs {
		wg.Add(1)
		go fetchNews(url, username, &wg, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	var allArticles []models.Article
	for articles := range ch {
		if articles != nil {
			allArticles = append(allArticles, articles...)
		}
	}

	return allArticles, nil
}

func convertToPyArticles(articles []models.Article) []models.PyArticle {
	var pyArticles []models.PyArticle
	for _, article := range articles {
		pyArticles = append(pyArticles, models.PyArticle{
			Title:       article.Title,
			Content:     article.Content,
			Author:      article.Author,
			URL:         article.URL,
			ImageURL:    article.ImageURL,
			Language:    article.Language,
			PublishedAt: article.PublishedAt,
		})
	}
	return pyArticles
}

func syncWithPythonService(articles []models.PyArticle) error {
	pythonServiceURL := os.Getenv("PYTHON_RECOMMENDATION_SERVICE_URL") + "/sync"
	jsonData, err := json.Marshal(articles)
	if err != nil {
		return err
	}

	resp, err := http.Post(pythonServiceURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-OK response from Python service: %v", resp.Status)
	}

	return nil
}

func fetchRecommendationsFromPythonService(title string) ([]models.PyArticle, error) {
	pythonServiceURL := os.Getenv("PYTHON_RECOMMENDATION_SERVICE_URL") + "/recommend"
	jsonData, err := json.Marshal(map[string]string{"title": title})
	if err != nil {
		return nil, fmt.Errorf("error marshalling request: %v", err)
	}

	resp, err := http.Post(pythonServiceURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error fetching recommendations from Python service: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-OK response from Python service: %v", resp.Status)
	}

	var recommendations []models.PyArticle
	err = json.NewDecoder(resp.Body).Decode(&recommendations)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return recommendations, nil
}

func storeRecommendationsForAllUsers(articles []models.PyArticle) error {
	users, err := db.GetAllUsers()
	if err != nil {
		return fmt.Errorf("error fetching users: %v", err)
	}

	for _, user := range users {
		for _, article := range articles {
			recommendations, err := fetchRecommendationsFromPythonService(article.Title)
			if err != nil {
				log.Println("Error fetching recommendations for user:", user.Username, "error:", err)
				continue
			}

			err = db.StoreRecommendations(user.Username, recommendations)
			if err != nil {
				log.Println("Error storing recommendations for user:", user.Username, "error:", err)
			}
		}
	}

	return nil
}
