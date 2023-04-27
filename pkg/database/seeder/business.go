package seeder

import (
	"62teknologi-senior-backend-test-muhammad-hajid-al-akhtar/domain"
	"gorm.io/gorm"
)

func CreateBusiness(db *gorm.DB, b domain.Business) error {
	return db.Create(&b).Error
}
func CreateCategory(db *gorm.DB, c domain.Category) error {
	return db.Create(&c).Error
}
