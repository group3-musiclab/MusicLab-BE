package services

import (
	"errors"
	"log"
	"musiclab-be/features/auth"
	"musiclab-be/features/transactions"
	"musiclab-be/utils/consts"
	"musiclab-be/utils/helper"
	"time"

	"github.com/go-playground/validator/v10"
)

type authUseCase struct {
	qry       auth.AuthData
	qryTrans  transactions.TransactionData
	validate  *validator.Validate
	googleApi helper.GoogleAPI
}

// CreateEvent implements auth.AuthService
func (auc *authUseCase) CreateEvent(code, orderID string) error {

	token, err := auc.googleApi.GetToken(code)
	if err != nil {
		log.Println("get token in create event error: ", err)
		return errors.New("failed to create event in calendar")
	}

	startTime := time.Date(2023, 3, 29, 8, 0, 0, 0, time.UTC)
	endTime := time.Date(2023, 3, 29, 8, 5, 0, 0, time.UTC)

	startRFC := startTime.Format(time.RFC3339)
	endRFC := endTime.Format(time.RFC3339)

	detailCal := helper.CalendarDetail{
		Summary:  "Kelas Gitar",
		Location: "Lewo",
		Start:    startRFC,
		End:      endRFC,
		// nanti diisi email guest dan host
		Emails: []string{"blbla@gmail.com"}, // email guest
	}

	err = auc.googleApi.CreateCalendar(token, detailCal)
	if err != nil {
		log.Println("failed create event", err.Error())
		return errors.New("failed to create event in calendar")
	}

	return nil
}

// Register implements auth.AuthService
func (auc *authUseCase) Register(newUser auth.Core) error {
	errValidate := auc.validate.Struct(newUser)
	if errValidate != nil {
		return errors.New("validate: " + errValidate.Error())
	}

	// check duplicate email
	if newUser.Role == "Mentor" {
		_, errMentorLogin := auc.qry.LoginMentor(newUser.Email)
		if errMentorLogin == nil {
			return errors.New(consts.AUTH_DuplicateEmail)
		}
	}

	if newUser.Role == "Student" {
		_, errStudentLogin := auc.qry.LoginStudent(newUser.Email)
		if errStudentLogin == nil {
			return errors.New(consts.AUTH_DuplicateEmail)
		}
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

func New(ud auth.AuthData, ga helper.GoogleAPI, td transactions.TransactionData) auth.AuthService {
	return &authUseCase{
		qry:       ud,
		qryTrans:  td,
		validate:  validator.New(),
		googleApi: ga,
	}
}
