package services

import (
	"errors"
	"log"
	"musiclab-be/app/config"
	"musiclab-be/features/auth"
	"musiclab-be/helper"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
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

func (auc *authUseCase) Login(email string, password string) (string, auth.Core, error) {
	res, err := auc.qry.Login(email)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data not found"
		} else {
			msg = "there is a problem with the server"
		}
		return "", auth.Core{}, errors.New(msg)
	}

	if err := helper.ComparePassword(res.Password, password); err != nil {
		log.Println("login compare", err.Error())
		return "", auth.Core{}, errors.New("password not matched")
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userID"] = res.ID
	claims["exp"] = time.Now().Add(time.Hour * 48).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	useToken, _ := token.SignedString([]byte(config.JWTKey))

	return useToken, res, nil
}
