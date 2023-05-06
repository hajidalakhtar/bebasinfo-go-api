package domain

import (
	"context"
	"github.com/google/uuid"
)

type News struct {
	ID      uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Title   string    ` json:"title"`
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

type NewsUsecase interface {
	Find(ctx context.Context, date string, source string) ([]News, error)
	Store(ctx context.Context, source string) ([]News, error)
	//Update(ctx context.Context, bs *Business, id uuid.UUID) error
	//Delete(ctx context.Context, id uuid.UUID) error
}

type PosgresqlNewsRepository interface {
	Find(ctx context.Context, date string, source string) ([]News, error)
	Store(ctx context.Context, ns []News) error
}

type RSSNewsRepository interface {
	GetFromRSS(ctx context.Context, source string) ([]News, error)
}
