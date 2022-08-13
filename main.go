package main

import (
	"golang-gin-todo-list/app/model"
	"golang-gin-todo-list/config"
	"golang-gin-todo-list/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.Connect()
	config.Migrate(
		&model.Activity{},
		&model.Todo{},
	)

	router := gin.Default()

	routes.Router(router)

	router.Run()
}
