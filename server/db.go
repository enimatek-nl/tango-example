package server

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title   string `gorm:"column:title"`
	Content string `gorm:"column:content"`
}
