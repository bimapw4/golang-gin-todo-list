package model

import "gorm.io/gorm"

type Activity struct {
	gorm.Model
	Email string `gorm:"size:50"`
	Title string `gorm:"size:100"`
}
