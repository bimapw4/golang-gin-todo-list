package controller

import (
	"golang-gin-todo-list/app/entity"
	"golang-gin-todo-list/app/model"
	"golang-gin-todo-list/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateTodo(ctx *gin.Context) {
	var Todo model.Todo
	ctx.BindJSON(&Todo)

	err := validator.New().Struct(Todo)

	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status":  "bad request",
				"message": v.Field() + " cannot be null",
				"data":    gin.H{},
			})
			return
		}
	}

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

	result := config.DB.Model(&Todo).Where("id = ?", ctx.Param("id")).First(&Todo)

	if result.RowsAffected <= 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": "Todo with ID " + ctx.Param("id") + " Not Found",
			"data":    gin.H{},
		})
		return
	}

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

	err := validator.New().Struct(req)

	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status":  "bad request",
				"message": v.Field() + " cannot be null",
				"data":    gin.H{},
			})
			return
		}
	}

	result := config.DB.Model(&Todo).Where("id = ?", ctx.Param("id")).First(&Todo)

	if result.RowsAffected <= 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": "Todo with ID " + ctx.Param("id") + " Not Found",
			"data":    gin.H{},
		})
		return
	}

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
	result := config.DB.Where("id = ?", ctx.Param("id")).Delete(&Todo)

	if result.RowsAffected <= 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": "Todo with ID " + ctx.Param("id") + " Not Found",
			"data":    gin.H{},
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    gin.H{},
	})
}
