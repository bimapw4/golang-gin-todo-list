package routes

import (
	"golang-gin-todo-list/app/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusAccepted, gin.H{
			"message": "Hello World",
		})
	})

	ag := r.Group("activity-groups")
	ag.POST("/", controller.CreateActivity)
	ag.GET("/", controller.GetAllActivity)
}
