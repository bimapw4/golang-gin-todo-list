package main

import (
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

	router := gin.Default()

	routes.Router(router)

	router.Run()
}
