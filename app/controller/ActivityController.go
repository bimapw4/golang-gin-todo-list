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
	var Activity []model.Activity
	// var ActivityOutput []entity.GetActivityResponse
	config.DB.Model(&Activity).Find(&Activity)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    Activity,
	})
}

func GetOneActivity(ctx *gin.Context) {
	var Activity model.Activity
	// var ActivityOutput entity.GetActivityResponse
	config.DB.Model(&Activity).Where("id = ?", ctx.Param("id")).First(&Activity)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    Activity,
	})
}

func UpdateActivity(ctx *gin.Context) {
	var Activity model.Activity
	var req entity.UpdateActivity

	ctx.BindJSON(&req)
	config.DB.Model(&Activity).Where("id = ?", ctx.Param("id")).First(&Activity)

	Activity.Title = req.Title

	config.DB.Save(&Activity)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    Activity,
	})
}

func DeleteActivity(ctx *gin.Context) {
	var Activity model.Activity
	config.DB.Where("id = ?", ctx.Param("id")).Delete(&Activity)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    gin.H{},
	})
}
