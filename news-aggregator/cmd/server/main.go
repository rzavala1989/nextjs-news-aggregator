// main.go
package main

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"news-aggregator/internal/auth"
	"news-aggregator/internal/db"
	"news-aggregator/internal/news"
	"news-aggregator/internal/user"
)

func main() {
	log.Println("Starting server...")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println(".env file loaded successfully")

	dbInstance := db.SetupDatabase()
	defer dbInstance.Close()

	// Add indexes to improve performance
	db.AddIndexes()

	// Populate the articles.csv file with sample data
	createArticlesCSV()

	r := mux.NewRouter()

	// Auth routes
	r.HandleFunc("/login", auth.LoginHandler).Methods("POST")
	r.HandleFunc("/register", auth.RegisterHandler).Methods("POST")

	// Secure routes
	s := r.PathPrefix("/secure").Subrouter()
	s.Use(auth.Authenticate)
	s.HandleFunc("/news", news.NewsHandler).Methods("GET")
	s.HandleFunc("/user", user.UserHandler).Methods("GET")
	s.HandleFunc("/sync", news.SyncArticlesHandler).Methods("POST")
	s.HandleFunc("/recommendations", news.RecommendationsHandler).Methods("GET") // New route for recommendations

	// Public routes
	r.HandleFunc("/search", news.SearchHandler).Methods("GET")
	r.HandleFunc("/trending", handleTrendingSearches).Methods("GET")

	// Apply CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	// Update trending searches daily
	go func() {
		for {
			log.Println("Updating trending searches...")
			news.UpdateTrendingSearches(os.Getenv("SERPAPI_API_KEY"))
			time.Sleep(24 * time.Hour)
		}
	}()

	// Sync articles and update recommendations periodically
	go func() {
		for {
			log.Println("Syncing articles and updating recommendations...")
			news.SyncArticlesAndUpdateRecommendations()
			time.Sleep(12 * time.Hour) // Sync every 12 hours, adjust as needed
		}
	}()

	// Archive old articles daily
	go news.ArchiveOldArticles()

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func handleTrendingSearches(w http.ResponseWriter, r *http.Request) {
	date := time.Now().Format("2006-01-02")
	trending, err := db.GetTrendingSearchesForDate(date)
	if err != nil {
		log.Println("Error fetching trending searches:", err)
		http.Error(w, "Error fetching trending searches", http.StatusInternalServerError)
		return
	}
	if len(trending) == 0 {
		log.Println("No trending searches for today. Updating...")
		news.UpdateTrendingSearches(os.Getenv("SERPAPI_API_KEY"))
		trending, err = news.GetTrendingSearchesFromDB()
		if err != nil {
			log.Println("Error fetching trending searches after update:", err)
			http.Error(w, "Error fetching trending searches", http.StatusInternalServerError)
			return
		}
	}
	err = json.NewEncoder(w).Encode(trending)
	if err != nil {
		log.Println("Error encoding trending searches to JSON:", err)
	}
}

func createArticlesCSV() {
	// Create the target directory if it doesn't exist
	dataDir := filepath.Join("..", "news-aggregator", "data")
	err := os.MkdirAll(dataDir, os.ModePerm)
	if err != nil {
		log.Fatalf("Error creating data directory: %v", err)
	}

	// Create and open the articles.csv file
	csvFilePath := filepath.Join(dataDir, "articles.csv")
	file, err := os.Create(csvFilePath)
	if err != nil {
		log.Fatalf("Error creating articles.csv file: %v", err)
	}
	defer file.Close()

	// Write sample data to articles.csv
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV headers
	headers := []string{"username", "title", "content", "author", "url", "image_url", "language", "published_at"}
	if err := writer.Write(headers); err != nil {
		log.Fatalf("Error writing headers to articles.csv: %v", err)
	}

	// Write sample data
	data := [][]string{
		{"testuser", "Election Fraud in the US", "There have been claims of election fraud in the US...", "Derek Hunter", "http://example.com/first-article", "https://images.pexels.com/photos/1550337/pexels-photo-1550337.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=2", "en", "2024-07-30"},
		{"testuser", "Climate Change Impacts", "Climate change is having a significant impact on the environment...", "Mr. Clean", "http://example.com/second-article", "https://images.pexels.com/photos/60013/desert-drought-dehydrated-clay-soil-60013.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=2", "en", "2024-07-31"},
		{"testuser", "Not Like Us", "Your best friend has a secret and a very weird case. He's not like us...", "Muhammed Beshear Smith, PhD.", "http://example.com/third-article", "https://images.pexels.com/photos/3824771/pexels-photo-3824771.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=2", "en", "2024-06-01"},
	}

	for _, record := range data {
		if err := writer.Write(record); err != nil {
			log.Fatalf("Error writing record to articles.csv: %v", err)
		}
	}

	log.Println("articles.csv created and populated successfully")
}
