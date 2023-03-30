package main

import (
	"log"
	"musiclab-be/app/config"
	"musiclab-be/app/database"
	"musiclab-be/app/router"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.InitConfig()
	db := database.InitDB(*cfg)
	database.Migrate(db)

	for {
		errView := database.CreateTableView(db)
		if errView != nil {
			log.Println(errView.Error())
		}
		e := echo.New()
		e.Pre(middleware.RemoveTrailingSlash())
		e.Use(middleware.CORS())
		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
		}))

		router.InitRouter(db, e)

		if err := e.Start(":8000"); err != nil {
			log.Println(err.Error())
		}
		time.Sleep(168 * time.Hour)
	}
}
