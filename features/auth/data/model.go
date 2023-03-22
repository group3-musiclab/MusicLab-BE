package data

import "gorm.io/gorm"

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
