package dtos

import (
	"goravel/app/models"

	"github.com/google/uuid"
)

type TodoDTO struct {
	ID uuid.UUID `json:"id"`

	Title       string    `json:"title"`
	CompletedAt string    `json:"completed_at"`
	UserID      uuid.UUID `json:"user_id"`
	CreatedAt   string    `json:"created_at"`
}

func TodoToDTO(todo models.Todo) TodoDTO {
	completedAt := ""
	if todo.CompletedAt != nil {
		completedAt = todo.CompletedAt.ToString()
	}

	return TodoDTO{
		ID:          todo.ID,
		Title:       todo.Title,
		UserID:      todo.UserID,
		CompletedAt: completedAt,
		CreatedAt:   todo.CreatedAt.ToString(),
	}
}

func TodoArrayToDTO(todos []models.Todo) []TodoDTO {
	var todoDTOs []TodoDTO
	for _, todo := range todos {
		todoDTO := TodoToDTO(todo)
		todoDTOs = append(todoDTOs, todoDTO)
	}

	return todoDTOs
}
