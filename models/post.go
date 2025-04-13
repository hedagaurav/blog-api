package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
	Tags    string `json:"tags"`
	Status  string `json:"status"`
	Draft   bool   `json:"draft" gorm:"default:false"`
}
