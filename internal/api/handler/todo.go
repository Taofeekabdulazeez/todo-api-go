package handler

import (
	"net/http"
	"todo-api-go/internal/api/request"
	"todo-api-go/internal/core/model"
	"todo-api-go/pkg/database"
	"todo-api-go/pkg/utils"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct{}

func NewTodoHandler() *TodoHandler {
	return &TodoHandler{}
}

func (h *TodoHandler) CreateTodo(ctx *gin.Context) {
	var body request.CreateTodoRequest
	if validationErr := ctx.ShouldBindJSON(&body); validationErr != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid Request Data",
			"error":   validationErr.Error(),
		})
		return
	}

	user, _ := ctx.Get("user")
	todo := model.Todo{
		Title:       body.Title,
		Description: body.Description,
		Priority:    body.Priority,
		Completed:   body.Completed,
		DueDate:     body.DueDate,
		UserID:      user.(model.User).ID,
	}

	if result := database.DB.Create(&todo); result.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to create todo",
			"error":   result.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Todo created!",
		"data":    todo,
	})

}

func (h *TodoHandler) GetTodo(ctx *gin.Context) {
	id := ctx.Param("id")
	var todo model.Todo

	if result := database.DB.First(&todo, utils.ParseInt(id)); result.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Error getting Todo",
			"error":   result.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Todo fetched!",
		"data":    todo,
	})

}

func (h *TodoHandler) GetAllTodos(ctx *gin.Context) {
	var todos []model.Todo
	user, _ := ctx.Get("user")

	if result := database.DB.Find(&todos, "user_id = ? ", user.(model.User).ID); result.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Error getting all Todos",
			"error":   result.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "success",
		"data":    todos,
	})
}

func (h *TodoHandler) UpdateTodo(ctx *gin.Context) {
	id := ctx.Param("id")
	var todo model.Todo

	if result := database.DB.First(&todo, utils.ParseInt(id)); result.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Internal Server Error",
			"error":   result.Error.Error(),
		})
		return
	}

	var body request.UpdateTodoRequest

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Invalid request data",
			"error":   err.Error(),
		})
	}

	todo.Title = body.Title
	todo.Description = body.Description
	todo.Description = body.Description
	todo.Completed = body.Completed
	todo.DueDate = body.DueDate

	if result := database.DB.Save(&todo); result.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Error updating todo data",
			"error":   result.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Todo updated!",
		"data":    todo,
	})
}

func (h *TodoHandler) DeleteTodo(ctx *gin.Context) {
	id := ctx.Param("id")

	if result := database.DB.Delete(&model.Todo{}, utils.ParseInt(id)); result.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "An error occured while deleting Todo: " + id,
			"error":   result.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Todo with ID: " + id + " deleted!",
		"data":    nil,
	})
}
