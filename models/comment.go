package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	PostId      uint   `json:"post_id" gorm:"not null"`
	AuthorId    uint   `json:"author_id" gorm:"not null"`
	ParentId    uint   `json:"parent_id" gorm:"default:0"`
	CommentBody string `json:"body" gorm:"not null" faker:"sentence"`

	Author User `json:"author" gorm:"foreignKey:AuthorId;references:ID"`
	Post   Post `json:"post" gorm:"foreignKey:PostId;references:ID"`
	// Parent    *Comment `json:"parent" gorm:"foreignKey:ParentId;references:ID"`
	// Children []Comment `json:"children" gorm:"foreignKey:ParentId;references:ID"`
}
