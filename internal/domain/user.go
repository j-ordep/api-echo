package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
    Id        string    `json:"id" gorm:"type:uuid;primaryKey"`
    Name      string    `json:"name" gorm:"type:varchar(100);not null"`
    Email     string    `json:"email" gorm:"type:varchar(100);unique;not null"`
    CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
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