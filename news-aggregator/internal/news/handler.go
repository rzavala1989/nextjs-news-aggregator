package news

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"news-aggregator/internal/db"
	"news-aggregator/internal/models"
	"os"
	"strconv"
	"sync"
)

// SearchHandler handles the search requests and saves articles with user information.
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	username := r.URL.Query().Get("username")

	if query == "" || username == "" {
		http.Error(w, "Query and username are required", http.StatusBadRequest)
		return
	}

	apiURLs := []string{
		fmt.Sprintf("https://api.currentsapi.services/v1/search?keywords=%s&language=en&apiKey=%s", query, os.Getenv("CURRENTS_API_KEY")),
		fmt.Sprintf("https://api.thenewsapi.com/v1/news/top?api_token=%s&locale=us&limit=5&search=%s", os.Getenv("THENEWSAPI_API_KEY"), query),
		fmt.Sprintf("https://newsdata.io/api/1/latest?apikey=%s&q=%s", os.Getenv("NEWSDATAIO_API_KEY"), query),
		fmt.Sprintf("http://api.mediastack.com/v1/news?access_key=%s&keywords=%s", os.Getenv("MEDIASTACK_API_KEY"), query),
	}

	var wg sync.WaitGroup
	ch := make(chan []models.Article)

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

	convertedArticles := convertToPyArticles(allArticles)

	err := syncWithPythonService(convertedArticles)
	if err != nil {
		log.Println("Error syncing with Python service:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(allArticles); err != nil {
		log.Println("Error encoding articles to JSON:", err)
	}
}

// NewsHandler handles the GET request to fetch news articles.
func NewsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("NewsHandler invoked")

	pageParam := r.URL.Query().Get("page")
	limitParam := r.URL.Query().Get("limit")
	username := r.URL.Query().Get("username")

	if username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit < 1 {
		limit = 10
	}

	apiURLs := []string{
		"https://api.currentsapi.services/v1/search?keywords=news&language=en&apiKey=" + os.Getenv("CURRENTS_API_KEY"),
		"https://api.thenewsapi.com/v1/news/top?api_token=" + os.Getenv("THENEWSAPI_API_KEY") + "&locale=us&limit=5",
		"https://newsdata.io/api/1/latest?apikey=" + os.Getenv("NEWSDATAIO_API_KEY") + "&q=news",
		"http://api.mediastack.com/v1/news?access_key=" + os.Getenv("MEDIASTACK_API_KEY") + "&keywords=news",
	}

	var wg sync.WaitGroup
	ch := make(chan []models.Article)

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

	convertedArticles := convertToPyArticles(allArticles)

	// Sync with Python service
	err = syncWithPythonService(convertedArticles)
	if err != nil {
		log.Println("Error syncing with Python service:", err)
	}

	start := (page - 1) * limit
	end := start + limit
	if start > len(allArticles) {
		start = len(allArticles)
	}
	if end > len(allArticles) {
		end = len(allArticles)
	}

	paginatedArticles := allArticles[start:end]

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(paginatedArticles); err != nil {
		log.Println("Error encoding articles to JSON:", err)
	}
}

// RecommendationsHandler handles the GET request to fetch recommendations for a user.
func RecommendationsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("RecommendationsHandler invoked")

	pageParam := r.URL.Query().Get("page")
	limitParam := r.URL.Query().Get("limit")
	username, ok := r.Context().Value("username").(string)
	if !ok || username == "" {
		log.Println("Error: username not found in context")
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit < 1 {
		limit = 10
	}

	recommendations, err := db.GetRecommendations(username, page, limit)
	if err != nil {
		log.Println("Error fetching recommendations:", err)
		http.Error(w, "Error fetching recommendations", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(recommendations); err != nil {
		log.Println("Error encoding recommendations to JSON:", err)
	}
}
