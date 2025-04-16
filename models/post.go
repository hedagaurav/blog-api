package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title    string `json:"title" faker:"sentence"`
	Content  string `json:"content" faker:"paragraph"`
	AuthorId uint   `json:"author_id" gorm:"not null"`
	Slug     string `json:"slug" gorm:"unique;not null" faker:"slug"`
	Tags     string `json:"tags"`
	Status   string `json:"status" gorm:"default:0"`
	Draft    bool   `json:"draft" gorm:"default:false" faker:"bool"`

	Author User `json:"author" gorm:"foreignKey:AuthorId;references:ID"`
}
