package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/url"
)

func NewPosgresqlDatabase(dbHost string, dbPort string, dbUser string, dbPass string, dbName string) *gorm.DB {
	fmt.Println(dbName)
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

	//dbConn.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	//dbConn.Migrator().DropTable(&domain.News{}, &domain.Image{}, &domain.User{})
	//dbConn.AutoMigrate(&domain.News{}, &domain.Image{}, &domain.User{})
	return dbConn
}
