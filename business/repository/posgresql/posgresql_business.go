package posgresql

import (
	"62teknologi-senior-backend-test-muhammad-hajid-al-akhtar/domain"
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type mysqlBusinessRepository struct {
	conn *gorm.DB
}

func NewMysqlBusinessRepository(conn *gorm.DB) domain.BusinessRepository {
	return &mysqlBusinessRepository{conn}
}

func (m mysqlBusinessRepository) Find(ctx context.Context, term string, sortBy string, limit int, offset int, openAt string) ([]domain.Business, error) {
	var businesses []domain.Business
	query := m.conn.Model(&domain.Business{})

	if limit != 0 {
		query = query.Limit(limit)
	}

	if offset != 0 {
		query = query.Offset(offset)
	}

	if term != "" {
		query = query.Where("name ILIKE ?", "%"+term+"%")
	}

	if openAt != "" {
		query = query.Where("openAt ILIKE ?", openAt)
	}

	switch sortBy {
	case "best_match":
		query = query.Order("ASC")
	case "rating":
		query = query.Order("rating DESC")
	case "review_count":
		query = query.Order("review_count DESC")
	case "distance":
		query = query.Order("distance ASC")
	}

	err := query.Find(&businesses).Error
	return businesses, err

}

func (m mysqlBusinessRepository) Store(ctx context.Context, bs *domain.Business) error {
	return m.conn.Create(&bs).Error
}

func (m mysqlBusinessRepository) Update(ctx context.Context, bs *domain.Business, id uuid.UUID) error {
	return m.conn.Where("id", id).Updates(&bs).Error
}

func (m mysqlBusinessRepository) Delete(ctx context.Context, id uuid.UUID) error {
	//INFO: HARD DELETE
	return m.conn.Unscoped().Delete(&domain.Business{}, id).Error

}
