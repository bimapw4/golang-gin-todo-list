package entity

type UpdateTodo struct {
	Title string `json:"title" validate:"required"`
}
