package data

import (
	_modelChat "musiclab-be/features/chats/data"
	_modelClass "musiclab-be/features/classes/data"
	_modelMentorGenres "musiclab-be/features/genres/data"
	"musiclab-be/features/instruments"
	_modelMentorInstruments "musiclab-be/features/instruments/data"
	"musiclab-be/features/mentors"
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
	TokenOauth        string
	Sex               string `gorm:"type:enum('','Male','Female')"`
	Phone             string `gorm:"type:varchar(12)"`
	Address           string `gorm:"type:varchar(100)"`
	Instagram         string
	About             string
	AvgRating         float32 `gorm:"type:float"`
	CountReviews      int64
	MentorInstrument  _modelMentorInstruments.MentorInstrument
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
	Type        string `gorm:"type:enum('International','National')"`
	Certificate string
}

func ModelToCore(data Mentor) mentors.Core {
	return mentors.Core{
		ID:           data.ID,
		Avatar:       data.Avatar,
		Name:         data.Name,
		Email:        data.Email,
		Password:     data.Password,
		Role:         data.Role,
		Sex:          data.Sex,
		Phone:        data.Phone,
		Address:      data.Address,
		Instagram:    data.Instagram,
		About:        data.About,
		AvgRating:    data.AvgRating,
		CountReviews: data.CountReviews,
		MentorInstrument: instruments.MentorInstrumentCore{
			Instrument: instruments.Core{
				Name: data.MentorInstrument.Instrument.Name,
			},
		},
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func CoreToModel(data mentors.Core) Mentor {
	return Mentor{
		Avatar:       data.Avatar,
		Name:         data.Name,
		Email:        data.Email,
		Password:     data.Password,
		Role:         data.Role,
		Sex:          data.Sex,
		Phone:        data.Phone,
		Address:      data.Address,
		Instagram:    data.Instagram,
		About:        data.About,
		AvgRating:    data.AvgRating,
		CountReviews: data.CountReviews,
	}
}

func ListModelToCore(dataModel []Mentor) []mentors.Core {
	var dataCore []mentors.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, ModelToCore(v))
	}
	return dataCore
}

func CredentialModelToCore(data Credential) mentors.CredentialCore {
	return mentors.CredentialCore{
		ID:          data.ID,
		MentorID:    data.MentorID,
		Name:        data.Name,
		Type:        data.Type,
		Certificate: data.Certificate,
	}
}

func CredentialCoreToModel(data mentors.CredentialCore) Credential {
	return Credential{
		MentorID:    data.MentorID,
		Name:        data.Name,
		Type:        data.Type,
		Certificate: data.Certificate,
	}
}
