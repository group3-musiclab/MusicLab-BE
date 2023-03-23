package data

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Avatar   string
	Name     string
	Email    string `gorm:"unique"`
	Password string
	Role     string `gorm:"type:varchar(25) not null default 'student'"`
	Sex      string
	Phone    string `gorm:"type:varchar(12)"`
	Address  string `gorm:"type:varchar(100)"`
}
