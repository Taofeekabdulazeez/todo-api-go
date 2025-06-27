package main

import (
	"net/http"
	"todo-api-go/configs"
	"todo-api-go/database"
	"todo-api-go/middlewares"
	"todo-api-go/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	configs.Init()

	router := gin.Default()

	router.Use(middlewares.Logger())

	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "home.html", nil)
	})

	routes.RegisterTodoRoutes(router)

	database.Connect()

	router.Run()
}
