package data

import (
	_modelMentorGenres "musiclab-be/features/genres/data"
	_modelMentorInstruments "musiclab-be/features/instruments/data"
	_modelSchedule "musiclab-be/features/schedules/data"

	"gorm.io/gorm"
)

type Mentor struct {
	gorm.Model
	Avatar            string
	Name              string
	Email             string `gorm:"unique"`
	Password          string
	Role              string `gorm:"type:varchar(25) not null default 'Mentor'"`
	Sex               string
	Phone             string `gorm:"type:varchar(12)"`
	Address           string `gorm:"type:varchar(100)"`
	Instagram         string
	Price             float64 `gorm:"type:float not null"`
	AvgRating         float32 `gorm:"type:float not null"`
	MentorInstruments []_modelMentorInstruments.MentorInstrument
	MentorGenres      []_modelMentorGenres.MentorGenre
	Credentials       []Credential
	Schedules         []_modelSchedule.Schedule
}

type Credential struct {
	gorm.Model
	MentorID    uint
	Name        string
	Type        string `gorm:"enum('International','National')"`
	Certificate string
}
