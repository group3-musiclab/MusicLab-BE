package data

import (
	"musiclab-be/features/auth"

	"gorm.io/gorm"
)

type authQuery struct {
	db *gorm.DB
}

// LoginMentor implements auth.AuthData
func (aq *authQuery) LoginMentor(email string) (auth.Core, error) {
	userLogin := Mentor{}
	txSelect := aq.db.Where("email = ?", email).First(&userLogin)
	if txSelect.Error != nil {
		return auth.Core{}, txSelect.Error
	}
	return (mentorToCore(userLogin)), nil
}

// LoginStudent implements auth.AuthData
func (aq *authQuery) LoginStudent(email string) (auth.Core, error) {
	userLogin := Student{}
	txSelect := aq.db.Where("email = ?", email).First(&userLogin)
	if txSelect.Error != nil {
		return auth.Core{}, txSelect.Error
	}
	return (studentToCore(userLogin)), nil
}

// RegisterMentor implements auth.AuthData
func (aq *authQuery) RegisterMentor(newUser auth.Core) error {
	dataModel := CoreToDataMentor(newUser)
	txInsert := aq.db.Create(&dataModel)
	if txInsert.Error != nil {
		return txInsert.Error
	}
	if txInsert.RowsAffected == 0 {
		return txInsert.Error
	}
	return nil
}

// RegisterStudent implements auth.AuthData
func (aq *authQuery) RegisterStudent(newUser auth.Core) error {
	dataModel := CoreToDataStudent(newUser)
	txInsert := aq.db.Create(&dataModel)
	if txInsert.Error != nil {
		return txInsert.Error
	}
	if txInsert.RowsAffected == 0 {
		return txInsert.Error
	}
	return nil
}

func New(db *gorm.DB) auth.AuthData {
	return &authQuery{
		db: db,
	}
}
