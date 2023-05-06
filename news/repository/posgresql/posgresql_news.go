package posgresql

import (
	"bebasinfo/domain"
	"context"
	"fmt"
	"gorm.io/gorm"
)

type posgresqlNewsRepository struct {
	conn *gorm.DB
}

func NewPosgresqlNewsRepository(conn *gorm.DB) domain.PosgresqlNewsRepository {
	return &posgresqlNewsRepository{conn}
}

func (m posgresqlNewsRepository) Store(ctx context.Context, ns []domain.News) error {
	return m.conn.Create(&ns).Error
}

func (m posgresqlNewsRepository) Find(ctx context.Context, date string, source string) ([]domain.News, error) {
	var news []domain.News
	query := m.conn.Preload("Image").Model(&domain.News{})

	if date != "" {
		//query = query.Offset(offset)
	}
	if source != "" {
		//query = query.Offset(offset)
	}

	//switch sortBy {
	//case "best_match":
	//	query = query.Order("ASC")
	//case "rating":
	//	query = query.Order("rating DESC")
	//case "review_count":
	//	query = query.Order("review_count DESC")
	//case "distance":
	//	query = query.Order("distance ASC")
	//}

	err := query.Find(&news).Error
	fmt.Println(err)
	return news, err

}

//func (m mysqlBusinessRepository) Store(ctx context.Context, bs *domain.Business) error {
//	return m.conn.Create(&bs).Error
//}
//
//func (m mysqlBusinessRepository) Update(ctx context.Context, bs *domain.Business, id uuid.UUID) error {
//	return m.conn.Where("id", id).Updates(&bs).Error
//}
//
//func (m mysqlBusinessRepository) Delete(ctx context.Context, id uuid.UUID) error {
//	//INFO: HARD DELETE
//	return m.conn.Unscoped().Delete(&domain.Business{}, id).Error
//
//}
