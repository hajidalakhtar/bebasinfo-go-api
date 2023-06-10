package main

import (
	_authsHttpDelivery "bebasinfo/authentication/delivery/http"
	_authRepoPG "bebasinfo/authentication/repository/posgresql"
	_authUsecase "bebasinfo/authentication/usecase"
	"bebasinfo/helper"
	_newsHttpDelivery "bebasinfo/news/delivery/http"
	_newsRepoAPI "bebasinfo/news/repository/api"
	_newsRepoPG "bebasinfo/news/repository/posgresql"
	_newsRepoRSS "bebasinfo/news/repository/rss"
	_newsUsecase "bebasinfo/news/usecase"
	"bebasinfo/pkg/database"
	"bebasinfo/pkg/utils"
	_userRepoPG "bebasinfo/user/repository/posgresql"
	"github.com/gofiber/fiber/v2"

	"github.com/spf13/viper"
	"time"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	helper.PanicIfError(err)
}

func main() {
	utils.InitLogger()
	logger := utils.GetLogger()
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)

	newsApiBaseUrl := viper.GetString(`newsapi.base_url`)
	newsApitoken := viper.GetString(`newsapi.token`)

	newsDataBaseUrl := viper.GetString(`newsdata.base_url`)
	newsDataToken := viper.GetString(`newsdata.token`)

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

	inr := _newsRepoAPI.NewAPINewsRepository(newsApiBaseUrl, newsApitoken)
	ndr := _newsRepoAPI.NewNewsDataRepository(newsDataBaseUrl, newsDataToken)

	uns := _userRepoPG.NewPosgresqlUserRepository(dbConn)
	anr := _authRepoPG.NewPosgresqlAuthRepository(dbConn)

	// Init Usecase
	bu := _newsUsecase.NewNewsUsecase(pnr, rnr, inr, ndr, timeoutContext)
	au := _authUsecase.NewAuthUsecase(anr, uns, timeoutContext)

	// Init Delivery
	_newsHttpDelivery.NewNewsHandler(app, bu, logger)
	_authsHttpDelivery.NewAuthHandler(app, au)

	err := app.Listen(portService)
	helper.PanicIfError(err)
}
