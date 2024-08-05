// internal/news/fetch.go

package news

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	"news-aggregator/internal/models"
)

// Define the response structures for the different APIs
type CurrentsAPIResponse struct {
	Status string            `json:"status"`
	News   []CurrentsArticle `json:"news"`
}

type CurrentsArticle struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
	URL         string `json:"url"`
	ImageURL    string `json:"image"`
	Language    string `json:"language"`
	PublishedAt string `json:"published"`
}

type NewsAPIResponse struct {
	Meta struct {
		Found    int `json:"found"`
		Returned int `json:"returned"`
		Limit    int `json:"limit"`
		Page     int `json:"page"`
	} `json:"meta"`
	Data []NewsAPIArticle `json:"data"`
}

type NewsAPIArticle struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
	URL         string `json:"url"`
	ImageURL    string `json:"image_url"`
	Language    string `json:"language"`
	PublishedAt string `json:"published_at"`
}

type NewsDataResponse struct {
	Status       string            `json:"status"`
	TotalResults int               `json:"totalResults"`
	Results      []NewsDataArticle `json:"results"`
}

type NewsDataArticle struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
	URL         string `json:"url"`
	ImageURL    string `json:"image_url"`
	Language    string `json:"language"`
	PublishedAt string `json:"published_at"`
}

type MediaStackResponse struct {
	Pagination struct {
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
		Count  int `json:"count"`
		Total  int `json:"total"`
	} `json:"pagination"`
	Data []MediaStackArticle `json:"data"`
}

type MediaStackArticle struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
	URL         string `json:"url"`
	ImageURL    string `json:"image_url"`
	Language    string `json:"language"`
	PublishedAt string `json:"published_at"`
}

func fetchNews(apiURL, username string, wg *sync.WaitGroup, ch chan []models.Article) {
	defer wg.Done()
	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error fetching news:", err)
		ch <- nil
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		ch <- nil
		return
	}

	var articles []models.Article

	switch {
	case apiURLContains(apiURL, "currentsapi"):
		var currentsResponse CurrentsAPIResponse
		if err := json.Unmarshal(body, &currentsResponse); err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			ch <- nil
			return
		}
		for _, news := range currentsResponse.News {
			articles = append(articles, models.Article{
				Title:       news.Title,
				Content:     news.Description,
				Author:      news.Author,
				URL:         news.URL,
				ImageURL:    news.ImageURL,
				Language:    news.Language,
				PublishedAt: news.PublishedAt,
				Username:    username,
			})
		}

	case apiURLContains(apiURL, "thenewsapi"):
		var newsAPIResponse NewsAPIResponse
		if err := json.Unmarshal(body, &newsAPIResponse); err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			ch <- nil
			return
		}
		for _, data := range newsAPIResponse.Data {
			articles = append(articles, models.Article{
				Title:       data.Title,
				Content:     data.Description,
				Author:      data.Author,
				URL:         data.URL,
				ImageURL:    data.ImageURL,
				Language:    data.Language,
				PublishedAt: data.PublishedAt,
				Username:    username,
			})
		}

	case apiURLContains(apiURL, "newsdata.io"):
		var newsDataResponse NewsDataResponse
		if err := json.Unmarshal(body, &newsDataResponse); err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			ch <- nil
			return
		}
		for _, result := range newsDataResponse.Results {
			articles = append(articles, models.Article{
				Title:       result.Title,
				Content:     result.Description,
				Author:      result.Author,
				URL:         result.URL,
				ImageURL:    result.ImageURL,
				Language:    result.Language,
				PublishedAt: result.PublishedAt,
				Username:    username,
			})
		}

	case apiURLContains(apiURL, "mediastack"):
		var mediaStackResponse MediaStackResponse
		if err := json.Unmarshal(body, &mediaStackResponse); err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			ch <- nil
			return
		}
		for _, data := range mediaStackResponse.Data {
			articles = append(articles, models.Article{
				Title:       data.Title,
				Content:     data.Description,
				Author:      data.Author,
				URL:         data.URL,
				ImageURL:    data.ImageURL,
				Language:    data.Language,
				PublishedAt: data.PublishedAt,
				Username:    username,
			})
		}
	}

	ch <- articles
}

func apiURLContains(apiURL, key string) bool {
	return stringContains(apiURL, key)
}

func stringContains(s, substr string) bool {
	return len(substr) == 0 || (len(s) >= len(substr) && (s[0:len(substr)] == substr || stringContains(s[1:], substr)))
}
