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

func (cuc *classUseCase) GetMentorClass(mentorID uint, page, limit int) ([]classes.Core, error) {
	offset := (page - 1) * limit
	res, err := cuc.qry.GetMentorClass(mentorID, limit, offset)

	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "class not found"
		} else {
			msg = "there is a problem with the server"
		}
		return []classes.Core{}, errors.New(msg)
	}

	return res, nil
}

func (cuc *classUseCase) GetMentorClassDetail(classID uint) (classes.Core, error) {
	res, err := cuc.qry.GetMentorClassDetail(classID)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return classes.Core{}, errors.New("data not found")
		} else {
			return classes.Core{}, errors.New("internal server error")
		}
	}

	return res, nil
}

func (cuc *classUseCase) Update(mentorID uint, classID uint, fileData multipart.FileHeader, updatedClass classes.Core) error {
	url, err := helper.GetUrlImagesFromAWS(fileData)
	if err != nil {
		return errors.New("validate: " + err.Error())
	}
	updatedClass.Image = url
	err = cuc.qry.Update(uint(mentorID), classID, updatedClass)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return errors.New("data not found")
		} else {
			return errors.New("internal server error")
		}
	}

	return nil
}

func (cuc *classUseCase) Delete(mentorID uint, classID uint) error {
	err := cuc.qry.Delete(mentorID, classID)

	if err != nil {
		log.Println("delete query error", err.Error())
		return errors.New("data not found")
	}

	return nil
}
