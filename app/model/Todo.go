package model

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	ActivityGroupId string `gorm:"size:11"`
	Title           string `gorm:"size:100"`
	Priority        string `gorm:"size:15"`
}
