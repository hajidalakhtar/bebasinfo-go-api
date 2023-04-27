package database

import (
	"62teknologi-senior-backend-test-muhammad-hajid-al-akhtar/domain"
	"62teknologi-senior-backend-test-muhammad-hajid-al-akhtar/pkg/database/seeder"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/url"
)

func NewPosgresqlDatabase(dbHost string, dbPort string, dbUser string, dbPass string, dbName string) *gorm.DB {
	connection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPass, dbName, dbPort)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s TimeZone=%s", connection, val.Encode())
	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
		log.Fatal(err)

	}

	dbConn.Migrator().DropTable(&domain.Business{}, &domain.Category{})
	dbConn.AutoMigrate(&domain.Business{}, &domain.Category{})
	seeder.Seed(dbConn)

	return dbConn
}
