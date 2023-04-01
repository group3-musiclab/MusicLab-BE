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

// FindAccount implements auth.AuthData
func (aq *authQuery) FindAccount(email string) (auth.Core, error) {
	dataMentor := Mentor{}
	dataStudent := Student{}

	var chanMentor = make(chan Mentor)

	var findMentor = func(email string) {
		aq.db.Where("email = ?", email).First(&dataMentor)
		chanMentor <- dataMentor
	}

	var chanStudent = make(chan Student)

	var findStudent = func(email string) {
		aq.db.Where("email = ?", email).First(&dataStudent)
		chanStudent <- dataStudent
	}

	go findMentor(email)
	go findStudent(email)

	var mentor = <-chanMentor
	coreMentor := mentorToCore(mentor)
	if coreMentor.Name != "" {
		return coreMentor, nil
	}

	var student = <-chanStudent
	coreStudent := studentToCore(student)
	if coreStudent.Name != "" {
		return coreStudent, nil
	}

	return auth.Core{}, errors.New(consts.QUERY_NotFound)
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
