package services

import (
	"errors"
	"log"
	"mime/multipart"
	"musiclab-be/features/classes"
	"musiclab-be/utils/helper"
	"strings"
)

type classUseCase struct {
	qry classes.ClassData
}

// Delete implements classes.ClassService
func (*classUseCase) Delete(mentorID uint, classID uint) error {
	panic("unimplemented")
}

// GetMentorClass implements classes.ClassService
func (*classUseCase) GetMentorClass(mentorID uint) ([]classes.Core, error) {
	panic("unimplemented")
}

// GetMentorClassDetail implements classes.ClassService
func (*classUseCase) GetMentorClassDetail(classID uint) (classes.Core, error) {
	panic("unimplemented")
}

// Update implements classes.ClassService
func (*classUseCase) Update(mentorID uint, classID uint, updatedClass classes.Core) (classes.Core, error) {
	panic("unimplemented")
}

func New(cd classes.ClassData) classes.ClassService {
	return &classUseCase{
		qry: cd,
	}
}

func (cuc *classUseCase) PostClass(fileData multipart.FileHeader, newClass classes.Core) error {
	url, err := helper.GetUrlImagesFromAWS(fileData)
	if err != nil {
		return errors.New("validate: " + err.Error())
	}
	newClass.Image = url
	err = cuc.qry.PostClass(newClass)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data not found"
		} else {
			msg = "server problem"
		}
		log.Println("error add query in service: ", err.Error())
		return errors.New(msg)
	}
	return nil
}
