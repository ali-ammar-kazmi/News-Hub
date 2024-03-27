package model

import "time"

type OpenWeatherRes struct {
	City string `json:"name"`
	Sys  struct {
		Country string `json:"country"`
	} `json:"sys"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Date DateData
}

type DateData struct {
	Day   int          `json:"day"`
	Month time.Month   `json:"month"`
	Year  int          `json:"year"`
	Week  time.Weekday `json:"week"`
}

type NewsRes struct {
	Articles []struct {
		Source struct {
			Name string `json:"name"`
		} `json:"source"`
		Title       string `json:"title"`
		Description string `json:"description"`
		URL         string `json:"url"`
		URLtoImg    string `json:"UrlToImage"`
		PublishedAt string `json:"publishedAt"`
	} `json:"articles"`
}
