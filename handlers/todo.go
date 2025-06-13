package handlers

import (
	"net/http"
	"todo-api-go/database"
	"todo-api-go/models"
	"todo-api-go/util"

	"github.com/gin-gonic/gin"
)

func CreateTodo(ctx *gin.Context) {
	var todo models.Todo

	err := ctx.ShouldBindJSON(&todo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Request Data",
		})
		return
	}

	result := database.DB.Create(&todo)

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Todo created!",
		"data":    todo,
	})

}

func GetTodo(ctx *gin.Context) {
	id := ctx.Param("id")
	var todo models.Todo
	result := database.DB.First(&todo, util.ParseInt(id))

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Todo not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    todo,
	})

}

func GetAllTodos(ctx *gin.Context) {
	var todos []models.Todo

	response := database.DB.Find(&todos)

	if response.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to get todos",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    todos,
	})
}

func UpdateTodo(ctx *gin.Context) {
	id := ctx.Param("id")
	var todo models.Todo

	result := database.DB.First(&todo, util.ParseInt(id))

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Todo does not exist",
		})
		return
	}

	var updateData models.Todo

	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request data",
		})
	}

	todo.Title = updateData.Title
	todo.Description = updateData.Description

	response := database.DB.Save(&todo)

	if response.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Error updating todo data",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    todo,
	})
}

func DeleteTodo(ctx *gin.Context) {
	id := ctx.Param("id")

	response := database.DB.Delete(&models.Todo{}, util.ParseInt(id))

	if response.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "An error occured while deleting Todo: " + id,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Todo with ID: " + id + " deleted!",
		"data":    nil,
	})
}
