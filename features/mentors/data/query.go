package data

import (
	"musiclab-be/features/mentors"

	"gorm.io/gorm"
)

type mentorQuery struct {
	db *gorm.DB
}

// Login implements mentors.MentorData
func (mq *mentorQuery) Login(email string) (mentors.Core, error) {
	userLogin := Mentor{}
	txSelect := mq.db.Where("email = ?", email).First(&userLogin)
	if txSelect.Error != nil {
		return mentors.Core{}, txSelect.Error
	}
	return ToCore(userLogin), nil
}

// Register implements mentors.MentorData
func (mq *mentorQuery) Register(newUser mentors.Core) error {
	dataModel := CoreToData(newUser)
	txInsert := mq.db.Create(&dataModel)
	if txInsert.Error != nil {
		return txInsert.Error
	}
	if txInsert.RowsAffected == 0 {
		return txInsert.Error
	}
	return nil
}

func New(db *gorm.DB) mentors.MentorData {
	return &mentorQuery{
		db: db,
	}
}
