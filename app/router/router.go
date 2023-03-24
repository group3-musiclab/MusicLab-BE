package router

import (
	authData "musiclab-be/features/auth/data"
	authHdl "musiclab-be/features/auth/handler"
	authSrv "musiclab-be/features/auth/services"
	"musiclab-be/utils/helper"

	genreData "musiclab-be/features/genres/data"
	genreHdl "musiclab-be/features/genres/handler"
	genreSrv "musiclab-be/features/genres/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {

	aData := authData.New(db)
	aSrv := authSrv.New(aData)
	aHdl := authHdl.New(aSrv)

	gData := genreData.New(db)
	gSrv := genreSrv.New(gData)
	gHdl := genreHdl.New(gSrv)

	// AUTH
	e.POST("/login", aHdl.Login())
	e.POST("/register", aHdl.Register())

	// Mentor Genre
	e.POST("/mentors/genres", gHdl.AddMentorGenre(), helper.JWTMiddleware())

}
