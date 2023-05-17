package posgresql

import (
	"bebasinfo/domain"
	"bebasinfo/news/repository/helper"
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type posgresqlNewsRepository struct {
	conn *gorm.DB
}

func NewPosgresqlNewsRepository(conn *gorm.DB) domain.PosgresqlNewsRepository {
	return &posgresqlNewsRepository{conn}
}

func (m posgresqlNewsRepository) Store(ctx context.Context, ns domain.News) error {
	return m.conn.Create(&ns).Error
}

func (m posgresqlNewsRepository) FindByTitle(ctx context.Context, title string) (domain.News, error) {
	var news domain.News
	err := m.conn.Where("title = ?", title).First(&news).Error
	return news, err
}

func (m posgresqlNewsRepository) Find(ctx context.Context, id uuid.UUID, date string, source string, page int, limit int) ([]domain.News, int64, error) {

	var news []domain.News
	var count int64

	offset := (page - 1) * limit
	selectedSource := helper.GetSelectedSource(source)

	err := m.conn.Model(&domain.News{}).Where("source", selectedSource.Name).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	query := m.conn.Preload("Image").Model(&domain.News{}).Offset(offset).Limit(limit)

	if id != uuid.Nil {

		query = query.Where("id", id)

	}

	if date != "" {
		//query = query.Offset()
	}

	if source != "" {
		query = query.Where("source", selectedSource.Name)

	}

	err = query.Find(&news).Error

	return news, count, err

}
