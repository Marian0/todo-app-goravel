package dtos

import (
	"goravel/app/helpers"
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
	return TodoDTO{
		ID:          todo.ID,
		Title:       todo.Title,
		UserID:      todo.UserID,
		CompletedAt: helpers.FormatNullTimeToISO(todo.CompletedAt),
		CreatedAt:   helpers.FormatTimeToISO(todo.CreatedAt),
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
