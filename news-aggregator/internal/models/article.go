package models

type Article struct {
	Title       string `json:"title"`
	Content     string `json:"content"`
	Author      string `json:"author,omitempty"`
	URL         string `json:"url"`
	ImageURL    string `json:"image_url,omitempty"`
	Language    string `json:"language"`
	PublishedAt string `json:"published_at"`
	Username    string `json:"username"`
}

type PyArticle struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Author      string `json:"author"`
	URL         string `json:"url"`
	ImageURL    string `json:"image_url"`
	Language    string `json:"language"`
	PublishedAt string `json:"published_at"`
}
