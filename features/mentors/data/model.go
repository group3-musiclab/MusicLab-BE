package data

import (
	_modelChat "musiclab-be/features/chats/data"
	_modelClass "musiclab-be/features/classes/data"
	_modelMentorGenres "musiclab-be/features/genres/data"
	_modelMentorInstruments "musiclab-be/features/instruments/data"
	_modelReview "musiclab-be/features/reviews/data"
	_modelSchedule "musiclab-be/features/schedules/data"
	_modelTransaction "musiclab-be/features/transactions/data"

	"gorm.io/gorm"
)

type Mentor struct {
	gorm.Model
	Avatar            string
	Name              string `gorm:"type:varchar(50) not null"`
	Email             string `gorm:"not null;unique;type:varchar(50)"`
	Password          string
	Role              string `gorm:"type:varchar(25) not null default 'Mentor'"`
	Sex               string
	Phone             string `gorm:"type:varchar(12)"`
	Address           string `gorm:"type:varchar(100)"`
	Instagram         string
	Price             float64 `gorm:"type:float"`
	AvgRating         float32 `gorm:"type:float"`
	MentorInstruments []_modelMentorInstruments.MentorInstrument
	MentorGenres      []_modelMentorGenres.MentorGenre
	Credentials       []Credential
	Schedules         []_modelSchedule.Schedule
	Classes           []_modelClass.Class
	Chats             []_modelChat.Chat
	Reviews           []_modelReview.Review
	Transactions      []_modelTransaction.Transaction
}

type Credential struct {
	gorm.Model
	MentorID    uint
	Name        string
	Type        string `gorm:"enum('International','National')"`
	Certificate string
}
