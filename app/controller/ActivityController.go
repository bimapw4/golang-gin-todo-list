package controller

import (
	"golang-gin-todo-list/app/entity"
	"golang-gin-todo-list/app/model"
	"golang-gin-todo-list/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateActivity(ctx *gin.Context) {
	var activity model.Activity
	ctx.BindJSON(&activity)

	err := validator.New().Struct(activity)

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
	config.DB.Model(&Activity).Find(&Activity)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    Activity,
	})
}

func GetOneActivity(ctx *gin.Context) {
	var Activity model.Activity
	result := config.DB.Model(&Activity).Where("id = ?", ctx.Param("id")).First(&Activity)

	if result.RowsAffected <= 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": "Activity with ID " + ctx.Param("id") + " Not Found",
			"data":    gin.H{},
		})
		return
	}

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

	result := config.DB.Model(&Activity).Where("id = ?", ctx.Param("id")).First(&Activity)

	if result.RowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": "Activity with ID " + ctx.Param("id") + " Not Found",
			"data":    gin.H{},
		})
		return
	}

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
	result := config.DB.Where("id = ?", ctx.Param("id")).Delete(&Activity)

	if result.RowsAffected <= 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": "Activity with ID " + ctx.Param("id") + " Not Found",
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
