package data

import (
	"errors"
	"musiclab-be/features/auth"
	"musiclab-be/utils/consts"

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
		return auth.Core{}, errors.New(consts.QUERY_NotFound)
	}
	return (mentorToCore(userLogin)), nil
}

// LoginStudent implements auth.AuthData
func (aq *authQuery) LoginStudent(email string) (auth.Core, error) {
	userLogin := Student{}
	txSelect := aq.db.Where("email = ?", email).First(&userLogin)
	if txSelect.Error != nil {
		return auth.Core{}, errors.New(consts.QUERY_NotFound)
	}
	return (studentToCore(userLogin)), nil
}

// RegisterMentor implements auth.AuthData
func (aq *authQuery) RegisterMentor(newUser auth.Core) error {
	newUser.Avatar = "https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460_1280.png"
	dataModel := CoreToDataMentor(newUser)
	txInsert := aq.db.Create(&dataModel)
	if txInsert.Error != nil {
		return errors.New(consts.QUERY_ErrorInsertData)
	}
	if txInsert.RowsAffected == 0 {
		return errors.New(consts.QUERY_NoRowsAffected)
	}
	return nil
}

// RegisterStudent implements auth.AuthData
func (aq *authQuery) RegisterStudent(newUser auth.Core) error {
	newUser.Avatar = "https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460_1280.png"
	dataModel := CoreToDataStudent(newUser)
	txInsert := aq.db.Create(&dataModel)
	if txInsert.Error != nil {
		return errors.New(consts.QUERY_ErrorInsertData)
	}
	if txInsert.RowsAffected == 0 {
		return errors.New(consts.QUERY_NoRowsAffected)
	}
	return nil
}

func New(db *gorm.DB) auth.AuthData {
	return &authQuery{
		db: db,
	}
}
