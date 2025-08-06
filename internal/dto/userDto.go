package dto

import (
	"api-echo/internal/domain"
	"time"
)

// Request
type InputDto struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
}

// Response
type OutputDto struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToDomain(input InputDto) *domain.User {
	return domain.NewUser(input.Name, input.Email)
}

func ToDto(user *domain.User) *OutputDto {
	return &OutputDto{
		Id: user.Id,
		Name: user.Name,
		Email: user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}