package helper

import (
	"bebasinfo/domain"
	"github.com/mmcdole/gofeed"
)

func RSSToNews(rss *gofeed.Feed, source string) []domain.News {
	var news []domain.News
	for _, item := range rss.Items {
		var image []domain.Image
		if item.Enclosures != nil {
			image = EnclosuresToImages(item.Enclosures)
		}

		news = append(news, domain.News{
			Title:   item.Title,
			Link:    item.Link,
			Content: item.Content,
			Date:    item.Published,
			Image:   image,
			Source:  source,
		})
	}
	return news
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

func GetSelectedSource(source string) domain.Source {
	switch source {
	case "suara":
		return domain.Suara
	case "cnn":
		return domain.CNN
	case "cnbc":
		return domain.CNBC
	case "republika":
		return domain.Republika
	case "tempo":
		return domain.Tempo
	case "antara":
		return domain.Antara
	case "kumparan":
		return domain.Kumparan
	case "okezone":
		return domain.Okezone
	case "liputan6":
		return domain.Liputan6
	case "bbc":
		return domain.BBC
	case "vice":
		return domain.Vice
	case "voa":
		return domain.VOA
	default:
		return domain.Source{}
	}
}
