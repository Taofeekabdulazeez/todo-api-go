package request

import "time"

type CreateTodoRequest struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Priority    string     `json:"priority"`
	Completed   bool       `json:"completed"`
	DueDate     *time.Time `json:"dueDate"`
}

type UpdateTodoRequest struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Priority    string     `json:"priority"`
	Completed   bool       `json:"completed"`
	DueDate     *time.Time `json:"dueDate"`
}
