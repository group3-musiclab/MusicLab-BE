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

func ToCore(data Student) auth.Core {
	return auth.Core{
		ID:       data.ID,
		Avatar:   data.Avatar,
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		Role:     data.Role,
	}
}

func CoreToData(data auth.Core) Student {
	return Student{
		Model:    gorm.Model{ID: data.ID},
		Avatar:   data.Avatar,
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		Role:     data.Role,
	}
}
