package controller

import (
	"golang-gin-todo-list/app/entity"
	"golang-gin-todo-list/app/model"
	"golang-gin-todo-list/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateActivity(ctx *gin.Context) {
	var activity model.Activity
	ctx.BindJSON(&activity)

	config.DB.Create(&activity)

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  "Success",
		"message": "Success",
		"data": gin.H{
			"id":         activity.ID,
			"created_at": activity.CreatedAt,
			"updated_at": activity.UpdatedAt,
			"title":      activity.Title,
			"email":      activity.Email,
		},
	})
}

func GetAllActivity(ctx *gin.Context) {
	var Activity entity.Activity
	var ActivityOutput []entity.GetActivityResponse
	config.DB.Model(&Activity).Find(&ActivityOutput)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    ActivityOutput,
	})
}
