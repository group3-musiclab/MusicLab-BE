package main

import (
	"log"
	"musiclab-be/app/config"
	"musiclab-be/app/database"
	"musiclab-be/app/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.InitConfig()
	db := database.InitDB(*cfg)
	// helper.ServerKey = cfg.SERVER_KEY_MIDTRANS
	database.Migrate(db)

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	router.InitRouter(db, e)

	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}
}
