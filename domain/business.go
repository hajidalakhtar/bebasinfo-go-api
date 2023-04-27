package domain

import (
	"context"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Business struct {
	ID           uuid.UUID      `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Alias        string         `json:"alias"`
	Name         string         `gorm:"unique;not null;type:varchar(100);default:null" json:"name"`
	ImageURL     string         `json:"image_url"`
	IsClosed     bool           `json:"is_closed"`
	URL          string         `json:"url"`
	ReviewCount  int            `json:"review_count"`
	Categories   []Category     `gorm:"many2many:business_categories;" json:"categories"`
	Rating       float32        `json:"rating"`
	Coordinates  Coordinates    `gorm:"embedded" json:"coordinates"`
	Transactions pq.StringArray `gorm:"type:text[]" json:"transactions"`
	Price        string         `json:"price"`
	Location     Location       `gorm:"embedded" json:"location"`
	Phone        string         `json:"phone"`
	DisplayPhone string         `json:"display_phone"`
	Distance     float32        `json:"distance"`
}

type Categories []Category

type Category struct {
	ID    uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"-"`
	Alias string    `json:"alias"`
	Title string    `json:"title"`
}

type Coordinates struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"-"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
}

type Location struct {
	ID             uuid.UUID      `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"-"`
	Address1       string         `json:"address1"`
	Address2       string         `json:"address2"`
	Address3       string         `json:"address3"`
	City           string         `json:"city"`
	ZipCode        string         `json:"zip_code"`
	Country        string         `json:"country"`
	State          string         `json:"state"`
	DisplayAddress pq.StringArray `gorm:"type:text[]" json:"display_address"`
}

type BusinessUsecase interface {
	Find(ctx context.Context, term string, sortBy string, limit int, offset int, openAt string) ([]Business, error)
	Store(ctx context.Context, bs *Business) error
	Update(ctx context.Context, bs *Business, id uuid.UUID) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type BusinessRepository interface {
	Find(ctx context.Context, term string, sortBy string, limit int, offset int, openAt string) ([]Business, error)
	Store(ctx context.Context, bs *Business) error
	Update(ctx context.Context, bs *Business, id uuid.UUID) error
	Delete(ctx context.Context, id uuid.UUID) error
}
