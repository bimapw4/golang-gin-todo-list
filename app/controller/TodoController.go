package controller

import (
	"golang-gin-todo-list/app/entity"
	"golang-gin-todo-list/app/model"
	"golang-gin-todo-list/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTodo(ctx *gin.Context) {
	var Todo model.Todo
	ctx.BindJSON(&Todo)

	Todo.Priority = "very-high"
	config.DB.Create(&Todo)

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    Todo,
	})
}

func GetAllTodo(ctx *gin.Context) {
	var Todo []model.Todo
	// var ActivityOutput []entity.GetActivityResponse
	query := config.DB.Model(&Todo)

	if ctx.Query("activity_group_id") != "" {
		query.Where("activity_group_id", ctx.Query("activity_group_id"))
	}

	query.Find(&Todo)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    Todo,
	})
}

func GetOneTodo(ctx *gin.Context) {
	var Todo model.Todo
	config.DB.Model(&Todo).Where("id = ?", ctx.Param("id")).First(&Todo)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    Todo,
	})
}

func UpdateTodo(ctx *gin.Context) {
	var Todo model.Todo
	var req entity.UpdateTodo

	ctx.BindJSON(&req)
	config.DB.Model(&Todo).Where("id = ?", ctx.Param("id")).First(&Todo)

	Todo.Title = req.Title

	config.DB.Save(&Todo)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    Todo,
	})
}

func DeleteTodo(ctx *gin.Context) {
	var Todo model.Todo
	config.DB.Where("id = ?", ctx.Param("id")).Delete(&Todo)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    gin.H{},
	})
}
