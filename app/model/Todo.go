package model

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ActivityGroupId uint           `gorm:"size:11" json:"activity_group_id" validate:"required"`
	Title           string         `gorm:"size:100" json:"title" validate:"required"`
	Priority        string         `gorm:"size:15" json:"priority"`
	ID              uint           `json:"id"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at"`
}
