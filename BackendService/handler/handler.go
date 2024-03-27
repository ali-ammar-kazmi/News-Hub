package handler

import (
	"encoding/json"
	"net/http"

	service "github.com/ali-ammar-kazmi/backend-service/service"
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
	}

	news := service.GetNewsData(category, page, pageSize)

	data, _ := json.Marshal(news)

	w.Write(data)
}
