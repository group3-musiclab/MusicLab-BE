package data

import (
	"errors"
	"log"
	"musiclab-be/features/auth"

	"gorm.io/gorm"
)

type authQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) auth.AuthData {
	return &authQuery{
		db: db,
	}
}

func (aq *authQuery) Register(newUser auth.Core) error {
	dupEmail := CoreToData(newUser)
	err := aq.db.Where("email = ?", newUser.Email).First(&dupEmail).Error
	if err == nil {
		log.Println("duplicated")
		return errors.New("email duplicated")
	}

	newUser.Avatar = "https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460_1280.png"

	cnv := CoreToData(newUser)
	err = aq.db.Create(&cnv).Error
	if err != nil {
		log.Println("query error", err.Error())
		return errors.New("server error")
	}

	newUser.ID = cnv.ID
	return nil
}

// Login implements auth.AuthData
func (*authQuery) Login(email string) (auth.Core, error) {
	return auth.Core{}, errors.New("data not found")
}
