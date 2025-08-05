package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUser(name string, email string) *User {
	user := &User{
		Id: uuid.New().String(),
		Name:name, 
		Email:email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return user
}