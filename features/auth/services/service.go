package services

import (
	"errors"
	"musiclab-be/features/auth"
	"musiclab-be/utils/consts"
	"musiclab-be/utils/helper"

	"github.com/go-playground/validator/v10"
)

type authUseCase struct {
	qry      auth.AuthData
	validate *validator.Validate
}

func New(ud auth.AuthData) auth.AuthService {
	return &authUseCase{
		qry:      ud,
		validate: validator.New(),
	}
}

// Register implements auth.AuthService
func (auc *authUseCase) Register(newUser auth.Core) error {
	errValidate := auc.validate.Struct(newUser)
	if errValidate != nil {
		return errors.New("validate: " + errValidate.Error())
	}

	// avatar value
	newUser.Avatar = "https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460_1280.png"

	hashed, errHash := helper.HashPassword(newUser.Password)
	if errHash != nil {
		return errors.New(consts.AUTH_ErrorHash)
	}
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
	} else {
		return errors.New(consts.AUTH_ErrorRole)
	}

	return nil
}

func (auc *authUseCase) Login(user auth.Core) (string, auth.Core, error) {
	errValidate := auc.validate.StructExcept(user, "Name")
	if errValidate != nil {
		return "", auth.Core{}, errors.New("validate: " + errValidate.Error())
	}
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
	} else {
		return "", auth.Core{}, errors.New(consts.AUTH_ErrorRole)
	}

	if !helper.CompareHashPassword(user.Password, res.Password) {
		return "", auth.Core{}, errors.New(consts.AUTH_ErrorComparePassword)
	}

	token, errToken := helper.CreateToken(res.ID, res.Role)
	if errToken != nil {
		return "", auth.Core{}, errors.New(consts.AUTH_ErrorCreateToken)
	}

	return token, res, nil
}
