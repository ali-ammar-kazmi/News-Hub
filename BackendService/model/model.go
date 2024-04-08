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

type ImageRes struct {
	Photos []struct {
		Id  uint   `json:"id"`
		URL string `json:"url"`
		Src struct {
			Original string `json:"original"`
			Medium   string `json:"medium"`
			Portrait string `json:"portrait"`
		} `json:"src"`
		Alt string `json:"alt"`
	} `json:"photos"`
}
