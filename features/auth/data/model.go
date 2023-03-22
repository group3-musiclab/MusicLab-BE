package data

import (
	"musiclab-be/features/auth"

	"gorm.io/gorm"
)

type Mentor struct {
	gorm.Model
	Name          string
	Avatar        string
	Email         string
	Password      string
	Role          string
	Sex           string
	Phone         string
	Address       string
	Instagram     string
	Price         float64
	AverageRating float64
}

type Student struct {
	gorm.Model
	Name     string
	Avatar   string
	Email    string
	Password string
	Role     string
	Sex      string
	Phone    string
	Address  string
}

func StudentToCore(data Student) auth.Core {
	return auth.Core{
		ID:       data.ID,
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		Role:     data.Role,
	}
}

func CoreToDataStudent(data auth.Core) Student {
	return Student{
		Model:    gorm.Model{ID: data.ID},
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		Role:     data.Role,
	}
}

func MentorToCore(data Mentor) auth.Core {
	return auth.Core{
		ID:       data.ID,
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		Role:     data.Role,
	}
}

func CoreToDataMentor(data auth.Core) Mentor {
	return Mentor{
		Model:    gorm.Model{ID: data.ID},
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		Role:     data.Role,
	}
}
