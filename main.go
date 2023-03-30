package main

import (
	"log"
	"musiclab-be/app/config"
	"musiclab-be/app/database"
	"musiclab-be/app/router"

	"github.com/jasonlvhit/gocron"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func CreateTableView(db *gorm.DB) error {
	errView := database.CreateTableView(db)
	if errView != nil {
		log.Println(errView.Error())
		return errView
	}
	log.Println("success create table view")
	return nil
}

func executeCronJob(db *gorm.DB) {
	gocron.Every(1).Monday().At("00:00").Do(CreateTableView, db)
	<-gocron.Start()
}

func main() {
	cfg := config.InitConfig()
	db := database.InitDB(*cfg)
	database.Migrate(db)

	go executeCronJob(db)

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
}
