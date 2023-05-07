package main

import (
	"bebasinfo/helper"
	_newsHttpDelivery "bebasinfo/news/delivery/http"
	_newsRepoPG "bebasinfo/news/repository/posgresql"
	_newsRepoRSS "bebasinfo/news/repository/rss"
	_newsUsecase "bebasinfo/news/usecase"

	_authsHttpDelivery "bebasinfo/authentication/delivery/http"
	_authRepoPG "bebasinfo/authentication/repository/posgresql"
	_authUsecase "bebasinfo/authentication/usecase"

	"bebasinfo/pkg/database"
	_userRepoPG "bebasinfo/user/repository/posgresql"
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
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "bebasinfo",
	})

	// Init Repository
	pnr := _newsRepoPG.NewPosgresqlNewsRepository(dbConn)
	rnr := _newsRepoRSS.NewRSSNewsRepository()

	uns := _userRepoPG.NewPosgresqlUserRepository(dbConn)
	anr := _authRepoPG.NewPosgresqlAuthRepository(dbConn)

	// Init Usecase
	bu := _newsUsecase.NewNewsUsecase(pnr, rnr, timeoutContext)
	au := _authUsecase.NewAuthUsecase(anr, uns, timeoutContext)

	// Init Delivery
	_newsHttpDelivery.NewNewsHandler(app, bu)
	_authsHttpDelivery.NewAuthHandler(app, au)

	err := app.Listen(portService)
	helper.PanicIfError(err)
}
