package server

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title   string `gorm:"column:title"`
	Content string `gorm:"column:content"`
	Done    bool   `gorm:"column:done"`
}

func (Todo) Bla() {}
