package dtos

import (
	"database/sql"
	"goravel/app/models"
	"time"

	"github.com/google/uuid"
)

type TodoDTO struct {
	ID uuid.UUID `json:"id"`

	Title       string       `json:"title"`
	CompletedAt sql.NullTime `json:"completed_at"`
	UserID      uuid.UUID    `json:"user_id"`
	CreatedAt   time.Time    `json:"created_at"`
}

func TodoToTodoDTO(todos []models.Todo) []TodoDTO {
	var todoDTOs []TodoDTO
	for _, todo := range todos {
		todoDTO := TodoDTO{
			ID:          todo.ID,
			Title:       todo.Title,
			CompletedAt: todo.CompletedAt,
			CreatedAt:   todo.CreatedAt,
		}
		todoDTOs = append(todoDTOs, todoDTO)
	}

	return todoDTOs
}
