package data

import (
	"musiclab-be/features/auth"

	"gorm.io/gorm"
)

type Mentor struct {
	gorm.Model
	Avatar     string
	Name       string
	Email      string
	Password   string
	Role       string
	TokenOauth string
	Sex        string
	Phone      string
	Address    string
	Instagram  string
	AvgRating  float32
}

type Student struct {
	gorm.Model
	Avatar   string
	Name     string
	Email    string
	Password string
	Role     string
	Sex      string
	Phone    string
	Address  string
}

func mentorToCore(data Mentor) auth.Core {
	return auth.Core{
		ID:         data.ID,
		Avatar:     data.Avatar,
		Name:       data.Name,
		Email:      data.Email,
		Password:   data.Password,
		Role:       data.Role,
		TokenOauth: data.TokenOauth,
	}
}

func studentToCore(data Student) auth.Core {
	return auth.Core{
		ID:       data.ID,
		Avatar:   data.Avatar,
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		Role:     data.Role,
	}
}

func CoreToDataMentor(data auth.Core) Mentor {
	return Mentor{
		Avatar:     data.Avatar,
		Name:       data.Name,
		Email:      data.Email,
		Password:   data.Password,
		Role:       data.Role,
		TokenOauth: data.TokenOauth,
		AvgRating:  data.AvgRating,
	}
}

func CoreToDataStudent(data auth.Core) Student {
	return Student{
		Avatar:   data.Avatar,
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		Role:     data.Role,
	}
}
