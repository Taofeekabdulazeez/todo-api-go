package handlers

import (
	"fmt"
	"net/http"
	"todo-api-go/database"
	"todo-api-go/models"
	"todo-api-go/util"

	"github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {
	var todo models.Todo

	err := c.ShouldBindJSON(&todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Request Data",
		})
		return
	}

	result := database.DB.Create(&todo)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Todo created!",
		"data":    todo,
	})

}

func GetTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo
	result := database.DB.First(&todo, util.ParseUInt(id))
	fmt.Println("Todo ID: ", util.ParseUInt(id))

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Todo not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    todo,
	})

}

func GetAllTodos(c *gin.Context) {
	var todos []models.Todo

	response := database.DB.Find(&todos)

	if response.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to get todos",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    todos,
	})
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo

	result := database.DB.First(&todo, util.ParseUInt(id))

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Todo does not exist",
		})
		return
	}

	var updateData models.Todo

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request data",
		})
	}

	todo.Title = updateData.Title
	todo.Description = updateData.Description

	response := database.DB.Save(&todo)

	if response.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error updating todo data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    todo,
	})
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	response := database.DB.Delete(&models.Todo{}, util.ParseUInt(id))

	if response.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "An error occured while deleting Todo: " + id,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Todo with ID: " + id + " deleted!",
		"data":    nil,
	})
}
