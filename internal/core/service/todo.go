package service

import (
	"todo-api-go/internal/api/request"
	"todo-api-go/internal/core/model"
	"todo-api-go/pkg/database"
)

type TodosService struct{}

func (s *TodosService) CreateTodo(data request.CreateTodoRequest) (*model.Todo, error) {

	todo := model.Todo{
		Title:       data.Title,
		Description: data.Description,
	}

	result := database.DB.Create(&todo)

	return &todo, result.Error
}
