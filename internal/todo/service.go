package todo

import "todo-api-go/pkg/database"

type TodosService struct{}

func (s *TodosService) CreateTodo(data CreateTodoRequest) (*Todo, error) {

	todo := Todo{
		Title:       data.Title,
		Description: data.Description,
	}

	result := database.DB.Create(&todo)

	return &todo, result.Error
}
