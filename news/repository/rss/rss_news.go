package rss

import (
	"bebasinfo/domain"
	"bebasinfo/news/repository/helper"
	"context"
	"fmt"
	"github.com/mmcdole/gofeed"
)

type rssNewsRepository struct {
}

func NewRSSNewsRepository() domain.RSSNewsRepository {
	return &rssNewsRepository{}
}

func (r rssNewsRepository) GetFromRSS(ctx context.Context, source string) ([]domain.News, error) {

	selectedSource := helper.GetSelectedSource(source)
	if selectedSource.Link == "" {
		return nil, fmt.Errorf("source not found")
	}
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(selectedSource.Link)
	if err != nil {
		fmt.Println("Error parsing RSS feed:", err)
	}

	news := helper.RSSToNews(feed, selectedSource.Name)
	return news, nil
}
