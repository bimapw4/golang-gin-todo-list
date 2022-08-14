package entity

import (
	"time"
)

type Activity struct {
	Email string `json:"email" validate:"required"`
	Title string `json:"title" validate:"required"`
}

type UpdateActivity struct {
	Title string `json:"title" validate:"required"`
}

type GetActivityResponse struct {
	Id        int       `json:"id"`
	Email     string    `json:"email"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}
