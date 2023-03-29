package main

import (
	"log"
	"musiclab-be/app/config"
	"musiclab-be/app/database"
	"musiclab-be/app/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// func testsometing(dbTest *gorm.DB) {
// 	// gmt, _ := time.LoadLocation("Asia/Jakarta")
// 	s := gocron.NewScheduler()
// 	s.Every(1).Seconds().Do(
// 		func() {
// 			err := dbTest.Raw("INSERT into genres (name) values ('Reage')")
// 			log.Println(err)
// 			if err != nil {
// 				log.Println(err.Error)
// 			}
// 		},
// 	)

// }

func main() {
	cfg := config.InitConfig()
	db := database.InitDB(*cfg)
	// helper.ServerKey = cfg.SERVER_KEY_MIDTRANS
	database.Migrate(db)

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	router.InitRouter(db, e)
	//test scheduling(belum work)
	// gmt, _ := time.LoadLocation("Asia/Jakarta")
	// s := gocron.NewScheduler(gmt)
	// s.Every(1).Day().At("00:00").Do(
	// 	func() {
	// 		err := db.Raw("INSERT into genres (name) values ('Rocks')")
	// 		log.Println(err)
	// 		if err != nil {
	// 			log.Println(err.Error)
	// 		}
	// 	},
	// )
	// s.StartBlocking()

	// // s := gocron.NewScheduler()
	// // s.Every(1).Seconds().Do(testsometing, db)
	// // <-s.Start()

	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}

}
