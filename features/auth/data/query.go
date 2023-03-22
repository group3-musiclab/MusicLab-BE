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
	dupEmail := CoreToDataStudent(newUser)
	err := aq.db.Where("email = ?", newUser.Email).First(&dupEmail).Error
	if err == nil {
		log.Println("duplicated")
		return errors.New("email duplicated")
	}
	student := CoreToDataStudent(newUser)
	if err = aq.db.Create(&student).Error; err != nil {
		mentor := CoreToDataMentor(newUser)
		if err = aq.db.Create(&mentor).Error; err != nil {
			log.Println("query error", err.Error())
			return errors.New("server error")
		}
		return nil
	}
	return nil
}

// Login implements auth.AuthData
func (*authQuery) Login(email string) (auth.Core, error) {
	panic("unimplemented")
}
