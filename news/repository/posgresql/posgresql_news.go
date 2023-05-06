package posgresql

import (
	"bebasinfo/domain"
	"bebasinfo/news/repository/helper"
	"context"
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

func (m posgresqlNewsRepository) Find(ctx context.Context, date string, source string) ([]domain.News, error) {
	var news []domain.News
	query := m.conn.Preload("Image").Model(&domain.News{})
	selectedSource := helper.GetSelectedSource(source)

	if date != "" {
		//query = query.Offset()
	}
	if source != "" {
		query = query.Where("source", selectedSource.Name)

	}

	err := query.Find(&news).Error
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
