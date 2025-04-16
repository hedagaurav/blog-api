package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;not null" faker:"username"`
	Password string `json:"password" gorm:"not null;default:test@123"`
	Email    string `json:"email" gorm:"unique;not null" faker:"email"`

	Posts []Post `json:"posts" gorm:"foreignKey:AuthorId;references:ID"`
}
