package data

import (
	_modelChat "musiclab-be/features/chats/data"
	_modelReview "musiclab-be/features/reviews/data"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Avatar   string
	Name     string
	Email    string `gorm:"unique"`
	Password string
	Role     string `gorm:"type:varchar(25) not null default 'Student'"`
	Sex      string
	Phone    string `gorm:"type:varchar(12)"`
	Address  string `gorm:"type:varchar(100)"`
	Chats    []_modelChat.Chat
	Reviews  []_modelReview.Review
}
