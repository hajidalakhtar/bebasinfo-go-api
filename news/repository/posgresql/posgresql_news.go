package posgresql

import (
	"bebasinfo/domain"
	"bebasinfo/news/repository/helper"
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
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

func (m posgresqlNewsRepository) Find(ctx context.Context, id uuid.UUID, date string, source []string, page int, limit int) ([]domain.News, int64, error) {

	now := time.Now()

	// Calculate the number of days to subtract to get to the start of the week (Monday)
	weekday := now.Weekday()
	var daysToSubtract int
	if weekday == time.Sunday {
		daysToSubtract = 6
	} else {
		daysToSubtract = int(weekday - time.Monday)
	}
	startOfWeek := now.AddDate(0, 0, -daysToSubtract)

	var news []domain.News
	var count int64
	offset := (page - 1) * limit
	query := m.conn.Debug().Preload("Image").Where("date::timestamp >= ?", startOfWeek.Format("2006-01-02 15:04:05")).Model(&domain.News{})

	if id != uuid.Nil {
		query = query.Where("id", id)
	}

	if len(source) > 0 {
		selectedSource := helper.GetSelectedSource(source)
		sourceName := make([]string, 0)
		for _, item := range selectedSource {
			sourceName = append(sourceName, item.Name)
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
