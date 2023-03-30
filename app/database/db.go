package database

import (
	"errors"
	"fmt"
	"log"
	"musiclab-be/app/config"
	_modelChat "musiclab-be/features/chats/data"
	_modelClass "musiclab-be/features/classes/data"
	_modelGenre "musiclab-be/features/genres/data"
	_modelInstrument "musiclab-be/features/instruments/data"
	_modelMentor "musiclab-be/features/mentors/data"
	_modelReview "musiclab-be/features/reviews/data"
	_modelSchedule "musiclab-be/features/schedules/data"
	_modelStudent "musiclab-be/features/students/data"
	_modelTransaction "musiclab-be/features/transactions/data"

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

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&_modelMentor.Mentor{},
		&_modelMentor.Credential{},
		&_modelInstrument.Instrument{},
		&_modelInstrument.MentorInstrument{},
		&_modelGenre.Genre{},
		&_modelGenre.MentorGenre{},
		&_modelStudent.Student{},
		&_modelSchedule.Schedule{},
		&_modelClass.Class{},
		&_modelChat.Chat{},
		&_modelTransaction.Transaction{},
		&_modelReview.Review{},
	)
}

func CreateTableView(db *gorm.DB) error {
	query := `CREATE OR REPLACE VIEW top_mentor_week AS
			SELECT m.id , m.name, ROUND(AVG(r.rating), 1) AS avg_rating, m.about, m.instagram FROM mentors m 
			LEFT JOIN reviews r ON m.id = r.mentor_id
			WHERE r.created_at > date_sub(now(), interval 1 week)
			GROUP BY m.id`

	tx := db.Exec(query)
	if tx.Error != nil {
		return errors.New(tx.Error.Error())
	}
	return nil
}
