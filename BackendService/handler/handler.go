package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ali-ammar-kazmi/backend-service/service"
)

func GetWeather(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	city := r.FormValue("cityName")
	if city == "" {
		city = "lucknow"
	}

	weather := service.GetWeatherData(city)

	data, _ := json.Marshal(weather)

	w.Write(data)
}

func GetNews(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	category := r.FormValue("category")
	page := r.FormValue("page")
	pageSize := r.FormValue("pageSize")
	if category == "" {
		category = "general"
		page = "1"
		pageSize = "6"
	}

	news := service.GetNewsData(category, page, pageSize)

	data, _ := json.Marshal(news)

	w.Write(data)
}

func GetImage(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	title := r.FormValue("title")
	if title == "" {
		title = "landscape"
	}

	image := service.GetImageData(title)

	data, _ := json.Marshal(image)

	w.Write(data)
}
