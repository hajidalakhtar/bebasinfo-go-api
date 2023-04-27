package main

import (
	_authsHttpDelivery "62teknologi-senior-backend-test-muhammad-hajid-al-akhtar/auth/delivery/http"
	_businessHttpDelivery "62teknologi-senior-backend-test-muhammad-hajid-al-akhtar/business/delivery/http"
	_businessRepo "62teknologi-senior-backend-test-muhammad-hajid-al-akhtar/business/repository/posgresql"
	_businessUsecase "62teknologi-senior-backend-test-muhammad-hajid-al-akhtar/business/usecase"
	"62teknologi-senior-backend-test-muhammad-hajid-al-akhtar/helper"
	"62teknologi-senior-backend-test-muhammad-hajid-al-akhtar/pkg/database"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	helper.PanicIfError(err)
}

func main() {

	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	portService := viper.GetString(`server.address`)
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	dbConn := database.NewPosgresqlDatabase(dbHost, dbPort, dbUser, dbPass, dbName)

	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "62teknologi-senior-backend-test-muhammad-hajid-al-akhtar",
	})

	// Init Repository

	br := _businessRepo.NewMysqlBusinessRepository(dbConn)

	// Init Usecase
	bu := _businessUsecase.NewBusinessUsecase(br, timeoutContext)

	// Init Delivery
	_businessHttpDelivery.NewBusinessHandler(app, bu)
	_authsHttpDelivery.NewAuthHandler(app)

	err := app.Listen(portService)
	helper.PanicIfError(err)
}
