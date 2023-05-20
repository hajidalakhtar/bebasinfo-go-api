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

func (r rssNewsRepository) GetFromRSS(ctx context.Context, source []string) ([]domain.News, error) {

	newsArr := make([]domain.News, 0)
	selectedSource := helper.GetSelectedSource(source)
	fmt.Println(selectedSource[0])

	for _, source := range selectedSource {
		fmt.Println(source.Name)
		if source.Link == "" {
			return nil, fmt.Errorf("source not found")
		}
		fp := gofeed.NewParser()
		feed, err := fp.ParseURL(source.Link)
		if err != nil {
			fmt.Println("Error parsing RSS feed:", err)
		}

		news := helper.RSSToNews(feed, source.Name)
		newsArr = append(newsArr, news...)
	}
	return newsArr, nil
}
