package services

import (
	"errors"
	"musiclab-be/features/auth"
	"musiclab-be/features/classes"
	"musiclab-be/features/mentors"
	"musiclab-be/features/schedules"
	"musiclab-be/features/students"
	"musiclab-be/features/transactions"
	"musiclab-be/utils/consts"
	"musiclab-be/utils/helper"
	"time"

	"github.com/go-playground/validator/v10"
	"golang.org/x/oauth2"
)

type authUseCase struct {
	qry         auth.AuthData
	qryTrans    transactions.TransactionData
	qryClass    classes.ClassData
	qryStudent  students.StudentData
	qrySchedule schedules.ScheduleData
	qryMentor   mentors.MentorData
	validate    *validator.Validate
	googleApi   helper.GoogleAPI
}

// LoginOauth implements auth.AuthService
func (auc *authUseCase) LoginOauth(input auth.Core) (string, auth.Core, error) {
	// validation
	if input.TokenOauth == "" {
		return "", auth.Core{}, errors.New("token oauth cannot empty")
	}

	// get user info with token
	tokenOauth := &oauth2.Token{
		AccessToken: input.TokenOauth,
	}

	coreGoogle, errUserInfo := auc.googleApi.GetUserInfo(tokenOauth)
	if errUserInfo != nil {
		return "", auth.Core{}, errUserInfo
	}

	// find data by email sequential
	res := auth.Core{}
	coreMentor, _ := auc.qry.LoginMentor(coreGoogle.Email)
	res = coreMentor

	if res.Name == "" {
		coreStudent, errStudent := auc.qry.LoginStudent(coreGoogle.Email)
		if errStudent != nil {
			return "", auth.Core{}, errStudent
		}
		res = coreStudent
	}

	// insert token oauth to data mentor or student
	if res.Role == "Mentor" {
		inputMentor := mentors.Core{}
		inputMentor.TokenOauth = input.TokenOauth

		errUpdateMentor := auc.qryMentor.UpdateData(res.ID, inputMentor)
		if errUpdateMentor != nil {
			return "", auth.Core{}, errUpdateMentor
		}
	} else if res.Role == "Student" {
		inputStudent := students.Core{}
		inputStudent.TokenOauth = input.TokenOauth

		errUpdateStudent := auc.qryStudent.UpdateData(res.ID, inputStudent)
		if errUpdateStudent != nil {
			return "", auth.Core{}, errUpdateStudent
		}
	}

	// generate token jwt
	tokenJWT, errTokenJWT := helper.CreateToken(res.ID, res.Role)
	if errTokenJWT != nil {
		return "", auth.Core{}, errors.New(consts.AUTH_ErrorCreateToken)
	}

	return tokenJWT, res, nil
}

// LoginOauth implements auth.AuthService
func (auc *authUseCase) RedirectGoogleCallback(code string) error {
	// get token oauth2
	token, errToken := auc.googleApi.GetToken(code)
	if errToken != nil {
		return errors.New("failed to get google oauth2 token")
	}

	// get user info with token
	coreGoogle, errUserInfo := auc.googleApi.GetUserInfo(token)
	if errUserInfo != nil {
		return errUserInfo
	}

	coreMentor, errMentor := auc.qry.LoginMentor(coreGoogle.Email)
	if errMentor != nil {
		return errMentor
	}

	// insert tokent oauth2 to mentor data
	inputMentor := mentors.Core{}
	inputMentor.TokenOauth = token.AccessToken

	errUpdateMentor := auc.qryMentor.UpdateData(coreMentor.ID, inputMentor)
	if errUpdateMentor != nil {
		return errUpdateMentor
	}

	return nil
}

// CreateEvent implements auth.AuthService
func (auc *authUseCase) CreateEvent(input auth.Core) error {
	// validation
	errValidate := auc.validate.StructExcept(input, "Name", "Email", "Password", "Role")
	if errValidate != nil {
		return errors.New("validate: " + errValidate.Error())
	}

	// get token oauth2 from mentor data
	coreMentor, errMentor := auc.qryMentor.SelectProfile(input.ID)
	if errMentor != nil {
		return errMentor
	}

	// token oauth2 validation
	if coreMentor.TokenOauth == "" {
		return errors.New("token oauth not generated yet, please login with google account first")
	}

	// transaction detail
	coreTrans, errTrans := auc.qryTrans.SelectOne(input.TransactionID)
	if errTrans != nil {
		return errTrans
	}

	// class detail
	coreClass, errClass := auc.qryClass.GetMentorClassDetail(coreTrans.ClassID)
	if errClass != nil {
		return errClass
	}

	// student detail
	coreStudent, errStudent := auc.qryStudent.SelectProfile(coreTrans.StudentID)
	if errStudent != nil {
		return errStudent
	}

	// schedule detail
	coreSchedule, errSchedule := auc.qrySchedule.DetailSchedule(coreTrans.ScheduleID)
	if errSchedule != nil {
		return errSchedule
	}

	mapDay := map[time.Time]string{}

	start := coreTrans.StartDate
	end := coreTrans.EndDate
	for d := start; !d.After(end); d = d.AddDate(0, 0, 1) {
		dayString := d.Format("Monday")
		if dayString == coreSchedule.Day {
			mapDay[d] = dayString
			break
		}
	}

	var firstDay time.Time
	for k := range mapDay {
		firstDay = k
	}

	firstDayStr := firstDay.Format("2006-01-02 ") + coreSchedule.StartTime
	startTime, err := time.Parse("2006-01-02 15:04", firstDayStr)
	if err != nil {
		return errors.New("invalid start date format must YYYY-MM-DD")
	}

	endDayStr := firstDay.Format("2006-01-02 ") + coreSchedule.EndTime
	endTime, err := time.Parse("2006-01-02 15:04", endDayStr)
	if err != nil {
		return errors.New("invalid start date format must YYYY-MM-DD")
	}

	endDate := coreTrans.EndDate

	startRFC := startTime.Format(time.RFC3339)
	endRFC := endTime.Format(time.RFC3339)
	endDateRFC := endDate.Format(time.RFC3339)

	detailCal := helper.CalendarDetail{
		Summary:             coreClass.Name,
		Location:            coreStudent.Address,
		StartTime:           startRFC,
		EndTime:             endRFC,
		EndDate:             endDateRFC,
		CreatorDisplayName:  coreMentor.Name,
		CreatorEmail:        coreMentor.Email,
		AttendeeDisplayName: coreStudent.Name,
		AttendeeEmail:       coreStudent.Email,
	}

	token := &oauth2.Token{
		AccessToken: coreMentor.TokenOauth,
	}

	errCreateEvent := auc.googleApi.CreateCalendar(token, detailCal)
	if errCreateEvent != nil {
		return errCreateEvent
	}

	return nil
}

// Register implements auth.AuthService
func (auc *authUseCase) Register(newUser auth.Core) error {
	errValidate := auc.validate.StructExcept(newUser, "TransactionID")
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
	errValidate := auc.validate.StructExcept(user, "Name", "TransactionID")
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

func New(ud auth.AuthData, ga helper.GoogleAPI, td transactions.TransactionData, cd classes.ClassData, sd students.StudentData, scd schedules.ScheduleData, md mentors.MentorData) auth.AuthService {
	return &authUseCase{
		qry:         ud,
		qryTrans:    td,
		qryClass:    cd,
		qryStudent:  sd,
		qrySchedule: scd,
		qryMentor:   md,
		validate:    validator.New(),
		googleApi:   ga,
	}
}
