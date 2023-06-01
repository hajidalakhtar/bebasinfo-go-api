package api

import (
	"bebasinfo/domain"
	"bebasinfo/news/repository/helper"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type apiNewsDataRepository struct {
	baseUrl string
	token   string
}

func NewNewsDataRepository(baseUrl string, token string) domain.APINewsDataRepository {
	return &apiNewsDataRepository{
		baseUrl: baseUrl,
		token:   token,
	}
}

func (a apiNewsDataRepository) GetFromAPI(ctx context.Context, category string, page string) ([]domain.News, string, error) {
	var apiResp domain.NewsDataApiResp
	resp, err := http.Get(a.baseUrl + "/api/1/news?country=id&category=" + category + "&apikey=" + a.token + "&page=" + page)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &apiResp)

	news := helper.NewsDataApiToNews(apiResp.Results, category)
	return news, apiResp.NextPage, nil
}
