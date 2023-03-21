package database

import (
	"fmt"
	"log"
	"musiclab-be/app/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(dc config.DBConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dc.DBUser, dc.DBPass, dc.DBHost, dc.DBPort, dc.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("database connection error : ", err.Error())
		return nil
	}

	return db
}

// func Migrate(db *gorm.DB) {
// 	db.AutoMigrate(
// 		user.User{},
// 		rooms.Room{},
// 		reservations.Reservation{},
// 		feedback.Feedback{},
// 	)
// }
