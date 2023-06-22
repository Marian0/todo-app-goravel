package dtos

import (
	"goravel/app/models"

	"github.com/google/uuid"
)

type UserDTO struct {
	ID uuid.UUID `json:"id"`

	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func UserToDTO(user models.User) UserDTO {
	return UserDTO{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.ToString(),
		UpdatedAt: user.UpdatedAt.ToString(),
	}
}

func UsersArrayToDTO(users []models.User) []UserDTO {
	var userDTOs []UserDTO
	for _, user := range users {
		todoDTO := UserToDTO(user)
		userDTOs = append(userDTOs, todoDTO)
	}

	return userDTOs
}
