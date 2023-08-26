package newsapiscrape

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type source struct {
	ID   string `json:'id'`
	Name string `json:'name'`
}

type article struct {
	src           source `json:'source'`
	Author        string `json:'author'`
	Title         string `json:'title'`
	Description   string `json:'description'`
	Url           string `json:'url'`
	UrlToImage    string `json:'urlToImage'`
	TimePublished string `json:'publishedAt'`
	Content       string `json:'content'`
}

type news struct {
	Status       string    `json:'status'`
	TotalResults int       `json:'totalResults'`
	Articles     []article `json:'articles'`
}

// Non-exported functions (private)

// exported functions (public)
type Search struct {
	key   string
	value news
}

func (s *Search) SetKey(apikey string) {
	s.key = apikey
	s.value = news{}
}

func (s Search) GetNews() *news {
	return &s.value
}

func (s Search) Search() (*news, error) {
	client := &http.Client{}
	getRequestUrl := "https://newsapi.org/v2/top-headlines?country=us&apiKey=" + s.key
	response, err := client.Get(getRequestUrl)
	if err != nil {
		return &news{}, err
	}
	resBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &news{}, err
	}
	err = json.Unmarshal(resBody, &s.value)
	if err != nil {
		return &news{}, err
	}

	return &s.value, err
}

func (n news) GetStatus() string {
	return n.Status
}

func (n news) GetResultCount() int {
	return n.TotalResults
}

func (n news) GetResultTuple() (string, int) {
	return n.Status, n.TotalResults
}
