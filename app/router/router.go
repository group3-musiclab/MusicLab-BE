package router

import (
	authData "musiclab-be/features/auth/data"
	authHdl "musiclab-be/features/auth/handler"
	authSrv "musiclab-be/features/auth/services"

	genreData "musiclab-be/features/genres/data"
	genreHdl "musiclab-be/features/genres/handler"
	genreSrv "musiclab-be/features/genres/services"

	mentorData "musiclab-be/features/mentors/data"
	mentorHdl "musiclab-be/features/mentors/handler"
	mentorSrv "musiclab-be/features/mentors/services"

	studentData "musiclab-be/features/students/data"
	studentHdl "musiclab-be/features/students/handler"
	studentSrv "musiclab-be/features/students/services"

	instrumentData "musiclab-be/features/instruments/data"
	instrumentHdl "musiclab-be/features/instruments/handler"
	instrumentSrv "musiclab-be/features/instruments/services"

	"musiclab-be/utils/helper"

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

	mData := mentorData.New(db)
	mSrv := mentorSrv.New(mData)
	mHdl := mentorHdl.New(mSrv)

	sData := studentData.New(db)
	sSrv := studentSrv.New(sData)
	sHdl := studentHdl.New(sSrv)

	iData := instrumentData.New(db)
	iSrv := instrumentSrv.New(iData)
	iHdl := instrumentHdl.New(iSrv)

	// Auth
	e.POST("/login", aHdl.Login())
	e.POST("/register", aHdl.Register())

	// Mentors
	e.GET("/mentors/profile", mHdl.GetProfile(), helper.JWTMiddleware())
	e.GET("/mentors/:id", mHdl.GetProfileByIdParam())
	e.PUT("/mentors", mHdl.UpdateData(), helper.JWTMiddleware())
	e.PUT("/mentors/password", mHdl.UpdatePassword(), helper.JWTMiddleware())
	e.POST("/mentors/credentials", mHdl.AddCredential(), helper.JWTMiddleware())
	e.DELETE("/mentors", mHdl.Delete(), helper.JWTMiddleware())

	// Student
	e.GET("/students/profile", sHdl.GetProfile(), helper.JWTMiddleware())
	e.PUT("/students", sHdl.UpdateData(), helper.JWTMiddleware())
	e.DELETE("/students", sHdl.Delete(), helper.JWTMiddleware())
	e.PUT("/students/password", sHdl.UpdatePassword(), helper.JWTMiddleware())

	// Instrument
	e.GET("/instruments", iHdl.GetAll())
	e.GET("/mentors/instruments", iHdl.GetAllByMentorID(), helper.JWTMiddleware())
	e.POST("/mentors/instruments", iHdl.Add(), helper.JWTMiddleware())
	e.DELETE("/mentors/instruments/:id", iHdl.Delete(), helper.JWTMiddleware())

	// Mentor Genre
	e.POST("/mentors/genres", gHdl.AddMentorGenre(), helper.JWTMiddleware())
	e.GET("/genres", gHdl.GetGenre())
	e.GET("/mentors/genres", gHdl.GetMentorGenre(), helper.JWTMiddleware())
	e.DELETE("/mentors/genres/:genre_id", gHdl.Delete(), helper.JWTMiddleware())

}
