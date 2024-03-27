package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/ali-ammar-kazmi/backend-service/model"
)

func GetNewsData(category, page, pageSize string) model.NewsRes {

	KEY := os.Getenv("NEWS_API_KEY")
	URL := "https://newsapi.org/v2/top-headlines"

	client := &http.Client{}
	news := model.NewsRes{}

	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		panic(err.Error())
	}

	q := req.URL.Query()
	q.Add("country", "in")
	q.Add("category", category)
	q.Add("apiKey", KEY)
	q.Add("page", page)
	q.Add("pageSize", pageSize)
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

	json.Unmarshal(body, &news)

	return news
}
