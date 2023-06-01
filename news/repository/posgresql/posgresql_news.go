package posgresql

import (
	"bebasinfo/domain"
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

func (m posgresqlNewsRepository) Find(ctx context.Context, id uuid.UUID, date string, source []string, category []string, page int, limit int) ([]domain.News, int64, error) {

	var news []domain.News
	var count int64
	offset := (page - 1) * limit
	query := m.conn.Debug().Preload("Image").Model(&domain.News{})

	if id != uuid.Nil {
		query = query.Where("id", id)
	}

	if len(category) > 0 {
		query = query.Where("category IN ?", category)
	}

	if len(source) > 0 {
		sourceName := make([]string, 0)
		for _, item := range source {
			sourceName = append(sourceName, item)
		}
		query = query.Where("source IN ?", sourceName)
	}

	if date != "" {
		//query = query.Offset()
	}
	query = query.Order("id::uuid")

	err := query.Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Offset(offset).Limit(limit).Find(&news).Error
	return news, count, err

}
