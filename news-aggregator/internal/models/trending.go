package models

type TrendingSearch struct {
	ID    int    `json:"id"`
	Query string `json:"query"`
	Date  string `json:"date"`
}
