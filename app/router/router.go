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

	reviewData "musiclab-be/features/reviews/data"
	reviewHdl "musiclab-be/features/reviews/handler"
	reviewSrv "musiclab-be/features/reviews/services"

	classData "musiclab-be/features/classes/data"
	classHdl "musiclab-be/features/classes/handler"
	classSrv "musiclab-be/features/classes/services"

	scheduleData "musiclab-be/features/schedules/data"
	scheduleHdl "musiclab-be/features/schedules/handler"
	scheduleSrv "musiclab-be/features/schedules/services"

	chatData "musiclab-be/features/chats/data"
	chatHdl "musiclab-be/features/chats/handler"
	chatSrv "musiclab-be/features/chats/services"

	transactionData "musiclab-be/features/transactions/data"
	transactionHdl "musiclab-be/features/transactions/handler"
	transactionSrv "musiclab-be/features/transactions/services"

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

	rData := reviewData.New(db)
	rSrv := reviewSrv.New(rData)
	rHdl := reviewHdl.New(rSrv)

	cData := classData.New(db)
	cSrv := classSrv.New(cData)
	cHdl := classHdl.New(cSrv)

	schData := scheduleData.New(db)
	schSrv := scheduleSrv.New(schData, cData)
	schHdl := scheduleHdl.New(schSrv)

	chtData := chatData.New(db)
	chtSrv := chatSrv.New(chtData, mData, sData)
	chtHdl := chatHdl.New(chtSrv)

	transData := transactionData.New(db)
	transSrv := transactionSrv.New(transData)
	transHdl := transactionHdl.New(transSrv)

	// Auth
	e.POST("/login", aHdl.Login())
	e.POST("/register", aHdl.Register())

	// Mentors
	e.GET("/mentors", mHdl.GetAll())
	e.GET("/mentors/:id", mHdl.GetProfileByIdParam())
	e.GET("/mentors/topweek", mHdl.GetByRating())
	e.GET("/mentors/profile", mHdl.GetProfile(), helper.JWTMiddleware())
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
	e.GET("/mentors/:id/instruments", iHdl.GetAllByMentorID())
	e.POST("/mentors/instruments", iHdl.Add(), helper.JWTMiddleware())
	e.DELETE("/mentors/instruments/:id", iHdl.Delete(), helper.JWTMiddleware())

	// Mentor Genres
	e.POST("/mentors/genres", gHdl.AddMentorGenre(), helper.JWTMiddleware())
	e.GET("/genres", gHdl.GetGenre())
	e.GET("/mentors/:id/genres", gHdl.GetMentorGenre())
	e.DELETE("/mentors/genres/:genre_id", gHdl.Delete(), helper.JWTMiddleware())

	// Mentor Review
	e.POST("/mentors/:mentor_id/reviews", rHdl.PostMentorReview(), helper.JWTMiddleware())
	e.GET("/mentors/:mentor_id/reviews", rHdl.GetMentorReview())

	// Mentor Class
	e.POST("/mentors/classes", cHdl.PostClass(), helper.JWTMiddleware())
	e.GET("/mentors/:mentor_id/class", cHdl.GetMentorClass(), helper.JWTMiddleware())
	e.GET("/class/:class_id", cHdl.GetMentorClassDetail(), helper.JWTMiddleware())
	e.PUT("/class/:class_id", cHdl.Update(), helper.JWTMiddleware())
	e.DELETE("/class/:class_id", cHdl.Delete(), helper.JWTMiddleware())

	// Mentor Schedule
	e.POST("/mentors/schedules", schHdl.PostSchedule(), helper.JWTMiddleware())
	e.GET("/mentors/:mentor_id/schedules", schHdl.GetMentorSchedule(), helper.JWTMiddleware())
	e.DELETE("/schedules/:schedule_id", schHdl.Delete(), helper.JWTMiddleware())
	e.POST("/schedules/check", schHdl.CheckSchedule(), helper.JWTMiddleware())

	// Chats
	e.POST("/chats", chtHdl.Add(), helper.JWTMiddleware())
	e.GET("/chats", chtHdl.GetAll(), helper.JWTMiddleware())
	e.GET("/inbox", chtHdl.GetByStudent(), helper.JWTMiddleware())

	// Transaction
	e.POST("/transactions", transHdl.MakeTransaction(), helper.JWTMiddleware())

}
