package services

import (
	"errors"
	"musiclab-be/features/auth"
	"musiclab-be/features/classes"
	"musiclab-be/features/schedules"
	"musiclab-be/features/students"
	"musiclab-be/features/transactions"
	"musiclab-be/utils/consts"
	"musiclab-be/utils/helper"
	"time"

	"github.com/go-playground/validator/v10"
)

type authUseCase struct {
	qry         auth.AuthData
	qryTrans    transactions.TransactionData
	qryClass    classes.ClassData
	qryStudent  students.StudentData
	qrySchedule schedules.ScheduleData
	validate    *validator.Validate
	googleApi   helper.GoogleAPI
}

// LoginOauth implements auth.AuthService
func (*authUseCase) LoginOauth(code string) error {
	panic("unimplemented")
}

// CreateEvent implements auth.AuthService
func (auc *authUseCase) CreateEvent(code, orderID string) error {
	// get token oauth2
	token, errToken := auc.googleApi.GetToken(code)
	if errToken != nil {
		return errors.New("failed to create event in calendar")
	}

	// transaction detail
	coreTrans, errTrans := auc.qryTrans.SelectOne("ALTA-MusicLab-2-iJykzCx5x5U99hva8SnWf8")
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
		Summary:     coreClass.Name,
		Location:    coreStudent.Address,
		StartTime:   startRFC,
		EndTime:     endRFC,
		EndDate:     endDateRFC,
		DisplayName: coreStudent.Name,
		Email:       coreStudent.Email,
	}

	errCreateEvent := auc.googleApi.CreateCalendar(token, detailCal)
	if errCreateEvent != nil {
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

func New(ud auth.AuthData, ga helper.GoogleAPI, td transactions.TransactionData, cd classes.ClassData, sd students.StudentData, scd schedules.ScheduleData) auth.AuthService {
	return &authUseCase{
		qry:         ud,
		qryTrans:    td,
		qryClass:    cd,
		qryStudent:  sd,
		qrySchedule: scd,
		validate:    validator.New(),
		googleApi:   ga,
	}
}
