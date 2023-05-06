package rss

import (
	"bebasinfo/domain"
	"context"
	"fmt"
	"github.com/mmcdole/gofeed"
)

type rssNewsRepository struct {
}

func NewRSSNewsRepository() *rssNewsRepository {
	return &rssNewsRepository{}
}

func (r rssNewsRepository) GetFromRSS(ctx context.Context, source string) ([]domain.News, error) {

	selectedSource := GetSelectedSource(source)
	if selectedSource.Link == "" {
		return nil, fmt.Errorf("source not found")
	}
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(selectedSource.Link)
	if err != nil {
		fmt.Println("Error parsing RSS feed:", err)
	}

	news := RSSToNews(feed, selectedSource.Name)
	return news, nil
}
