package services

import (
	"todo-api-go/database"
	"todo-api-go/models"
	"todo-api-go/requests"
)

type TodosService struct{}

func (s *TodosService) CreateTodo(data requests.CreateTodoRequest) (*models.Todo, error) {

	todo := models.Todo{
		Title:       data.Title,
		Description: data.Description,
	}

	result := database.DB.Create(&todo)

	return &todo, result.Error
}
