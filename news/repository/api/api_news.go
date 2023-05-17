package rss

import (
	"bebasinfo/domain"
	"bebasinfo/news/repository/helper"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type apiNewsRepository struct {
	baseUrl string
	token   string
}

func NewAPINewsRepository(baseUrl string, token string) domain.APINewsRepository {
	return &apiNewsRepository{
		baseUrl: baseUrl,
		token:   token,
	}
}

func (a apiNewsRepository) GetFromAPI(ctx context.Context, category string) ([]domain.News, error) {
	var apiResp domain.NewsApiResp

	resp, err := http.Get(a.baseUrl + "/top-headlines?country=id&category=" + category + "&apiKey=" + a.token)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &apiResp)

	news := helper.ApiToNews(apiResp.Articles)

	return news, nil
}
