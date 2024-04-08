package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/ali-ammar-kazmi/backend-service/model"
)

func GetImageData(title string) model.ImageRes {

	KEY := os.Getenv("PEXELS_API_KEY")
	URL := "https://api.pexels.com/v1/search"

	client := &http.Client{}
	image := model.ImageRes{}

	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		panic(err.Error())
	}

	req.Header.Add("Authorization", KEY)

	q := req.URL.Query()
	q.Add("query", title)
	q.Add("per_page", "1")
	q.Add("orientation", "portrait")
	q.Add("size", "small")
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

	json.Unmarshal(body, &image)

	return image
}
