package controllers

import (
	"net/http"
	"todo-api-go/models"

	"github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {
	todo := models.Todo{Id: "1", Title: "Assignment", Description: "Lorem ipsu"}

	c.JSON(http.StatusOK, gin.H{
		"message": "Todo created!",
		"data":    todo,
	})

}

func GetTodo(c *gin.Context) {
	todoId := c.Param("id")
	todo := models.Todo{Id: todoId}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    todo,
	})

}

func GetAllTodos(c *gin.Context) {
	todos := []models.Todo{
		{Id: "1", Title: "Assignment", Description: "CSC420 Assignment"},
		{Id: "1", Title: "Assignment", Description: "CSC420 Assignment"},
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    todos,
	})
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")

	todo := models.Todo{
		Id: id,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    todo,
	})
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"message": "Book with ID: " + id + " deleted!",
		"data":    nil,
	})
}
