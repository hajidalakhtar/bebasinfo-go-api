package domain

import (
	"context"
	"github.com/google/uuid"
)

type News struct {
	ID      uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Title   string    ` json:"title" gorm:"unique"`
	Link    string    ` json:"link"`
	Content string    ` json:"content"`
	Date    string    ` json:"date"`
	Source  string    ` json:"source"`
	Image   []Image   `gorm:"many2many" json:"image"`
}

type Image struct {
	ID     uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	URL    string    `json:"url,omitempty"`
	Length string    `json:"length,omitempty"`
	Type   string    `json:"type,omitempty"`
}

// API struct for News API
type NewsApiResp struct {
	Status       string           `json:"status"`
	TotalResults int              `json:"totalResults"`
	Articles     []NewsApiArticle `json:"articles"`
}

type NewsApiArticle struct {
	Source      NewsApiSource `json:"source"`
	Author      string        `json:"author"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Url         string        `json:"url"`
	UrlToImage  string        `json:"urlToImage"`
	PublishedAt string        `json:"publishedAt"`
	Content     string        `json:"content"`
}

type NewsApiSource struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type NewsUsecase interface {
	Search(ctx context.Context, date string, source string, page int, limit int) ([]News, PaginatedResponse, error)
	Find(ctx context.Context, newsId uuid.UUID) ([]News, error)

	Store(ctx context.Context, newsResource string, category string, source string) ([]News, error)
}

type PosgresqlNewsRepository interface {
	Find(ctx context.Context, id uuid.UUID, date string, source string, page int, limit int) ([]News, int64, error)
	Store(ctx context.Context, ns News) error
	FindByTitle(ctx context.Context, title string) (News, error)
}

type RSSNewsRepository interface {
	GetFromRSS(ctx context.Context, source string) ([]News, error)
}

type APINewsRepository interface {
	GetFromAPI(ctx context.Context, category string) ([]News, error)
}
