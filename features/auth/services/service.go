package services

import (
	"errors"
	"musiclab-be/features/auth"
	"musiclab-be/helper"
	"strings"
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

	err := auc.qry.Register(newUser)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "duplicated") {
			msg = "email already registered"
		} else {
			msg = "server error"
		}
		return errors.New(msg)
	}

	return nil
}

// Login implements auth.AuthService
func (*authUseCase) Login(email string, password string) (string, auth.Core, error) {
	panic("unimplemented")
}
