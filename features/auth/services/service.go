package services

import (
	"errors"
	"log"
	"musiclab-be/features/auth"
	"musiclab-be/helper"
)

type authUseCase struct {
	qry auth.AuthData
}

func New(ud auth.AuthData) auth.AuthService {
	return &authUseCase{
		qry: ud,
	}
}

// Register implements auth.AuthService
func (auc *authUseCase) Register(newUser auth.Core) error {
	if len(newUser.Password) != 0 {
		//validation
		err := helper.RegistrationValidate(newUser)
		if err != nil {
			return errors.New("validate: " + err.Error())
		}
	}
	hashed := helper.GeneratePassword(newUser.Password)
	newUser.Password = string(hashed)

	if newUser.Role == "Mentor" {
		errMentor := auc.qry.RegisterMentor(newUser)
		if errMentor != nil {
			return errMentor
		}
	} else if newUser.Role == "Student" {
		errStudent := auc.qry.RegisterStudent(newUser)
		if errStudent != nil {
			return errStudent
		}
	}

	return nil
}

func (auc *authUseCase) Login(user auth.Core) (string, auth.Core, error) {
	res := auth.Core{}
	if user.Role == "Mentor" {
		var err error
		res, err = auc.qry.LoginMentor(user.Email)
		if err != nil {
			return "", auth.Core{}, err
		}
	} else if user.Role == "Student" {
		var err error
		res, err = auc.qry.LoginStudent(user.Email)
		if err != nil {
			return "", auth.Core{}, err
		}
	}

	if err := helper.ComparePassword(res.Password, user.Password); err != nil {
		log.Println("login compare", err.Error())
		return "", auth.Core{}, errors.New("password not matched")
	}

	token, errToken := helper.CreateToken(user.ID, user.Role)
	if errToken != nil {
		return "", auth.Core{}, errToken
	}

	return token, res, nil
}
