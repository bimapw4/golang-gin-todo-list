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
	ag.GET("/:id", controller.GetOneActivity)
	ag.DELETE("/:id", controller.DeleteActivity)
	ag.PATCH("/:id", controller.UpdateActivity)

	td := r.Group("todo-items")
	td.POST("/", controller.CreateTodo)
	td.GET("/", controller.GetAllTodo)
	td.GET("/:id", controller.GetOneTodo)
	td.PATCH("/:id", controller.UpdateTodo)
	td.DELETE("/:id", controller.DeleteTodo)

}
