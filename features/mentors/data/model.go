package data

import "gorm.io/gorm"

type Mentor struct {
	gorm.Model
	Avatar    string
	Name      string
	Email     string `gorm:"unique"`
	Password  string
	Role      string `gorm:"type:varchar(25) not null default 'Mentor'"`
	Sex       string
	Phone     string `gorm:"type:varchar(12)"`
	Address   string `gorm:"type:varchar(100)"`
	Instagram string
	Price     float64 `gorm:"type:float not null"`
	AvgRating float32 `gorm:"type:float not null"`
}
