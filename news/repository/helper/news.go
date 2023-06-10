package helper

import (
	"bebasinfo/domain"
	"fmt"
	"github.com/mmcdole/gofeed"
	"strings"
)

func RSSToNews(rss *gofeed.Feed, source domain.Source) []domain.News {
	var news []domain.News
	for _, item := range rss.Items {
		var image []domain.Image
		if item.Enclosures != nil {
			image = EnclosuresToImages(item.Enclosures)
		}

		news = append(news, domain.News{
			Title:    item.Title,
			Link:     item.Link,
			Content:  item.Content,
			Date:     item.Published,
			Image:    image,
			Category: source.Category,
			Source:   source.Name,
		})
	}
	return news
}

func ApiToNews(news []domain.NewsApiArticle) []domain.News {
	var newsDomain []domain.News

	for _, newsItem := range news {
		newsDomain = append(newsDomain, domain.News{
			Title:   newsItem.Title,
			Link:    newsItem.Url,
			Content: newsItem.Content,
			Date:    newsItem.PublishedAt,
			Image: []domain.Image{
				{
					URL:    newsItem.UrlToImage,
					Length: "10",
					Type:   "",
				},
			},
			Source: newsItem.Source.Name,
		})
	}
	return newsDomain
}

func NewsDataApiToNews(news []domain.NewsApiResults, category string) []domain.News {
	var newsDomain []domain.News

	for _, newsItem := range news {
		if newsItem.ImageURL == "" {
			fmt.Println(newsItem.ImageURL)
			continue
		}
		if strings.Contains(newsItem.Title, "VIDEO:") {
			fmt.Println(newsItem.Title)
			continue
		}
		newsDomain = append(newsDomain, domain.News{
			Title:    newsItem.Title,
			Link:     newsItem.Link,
			Content:  newsItem.Content,
			Category: category,
			Date:     newsItem.PubDate,
			Image: []domain.Image{
				{
					URL:    newsItem.ImageURL,
					Length: "10",
					Type:   "",
				},
			},
			Source: newsItem.SourceID,
		})
	}
	return newsDomain
}

func EnclosuresToImages(es []*gofeed.Enclosure) []domain.Image {
	var images []domain.Image
	for _, enclosure := range es {
		images = append(images, domain.Image{
			URL:    enclosure.URL,
			Length: enclosure.Length,
			Type:   enclosure.Type,
		})
	}
	return images
}

func GetSelectedSource(sources []string) []domain.Source {
	sourcesArr := make([]domain.Source, 0)
	for _, item := range sources {
		switch item {
		case "suara":
			sourcesArr = append(sourcesArr, domain.Suara)
		case "cnn":
			sourcesArr = append(sourcesArr, domain.CNN)
		case "cnbc":
			sourcesArr = append(sourcesArr, domain.CNBC)
		case "republika":
			sourcesArr = append(sourcesArr, domain.Republika)
		case "tempo":
			sourcesArr = append(sourcesArr, domain.Tempo)
		case "antara":
			sourcesArr = append(sourcesArr, domain.Antara)
		case "kumparan":
			sourcesArr = append(sourcesArr, domain.Kumparan)
		case "okezone":
			sourcesArr = append(sourcesArr, domain.Okezone)
		case "bbc":
			sourcesArr = append(sourcesArr, domain.BBC)
		case "vice":
			sourcesArr = append(sourcesArr, domain.Vice)
		case "voa":
			sourcesArr = append(sourcesArr, domain.VOA)
		}
	}
	return sourcesArr
}
