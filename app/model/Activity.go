package model

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	// gorm.Model
	Email     string         `gorm:"size:50" json:"email"`
	Title     string         `gorm:"size:100" json:"title"`
	ID        uint           `json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
