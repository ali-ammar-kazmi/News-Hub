package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/ali-ammar-kazmi/backend-service/model"
)

func GetWeatherData(city string) model.OpenWeatherRes {

	KEY := os.Getenv("WEATHER_API_KEY")
	URL := "https://api.openweathermap.org/data/2.5/weather"
	client := &http.Client{}
	weather := model.OpenWeatherRes{}

	year, month, day := time.Now().Date()
	week := time.Now().Weekday()

	weather.Date = model.DateData{Day: day, Month: month, Year: year, Week: week}

	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		panic(err.Error())
	}

	q := req.URL.Query()
	q.Add("q", city)
	q.Add("appid", KEY)
	req.URL.RawQuery = q.Encode()

	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		panic(err.Error())
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	json.Unmarshal(body, &weather)

	return weather
}
