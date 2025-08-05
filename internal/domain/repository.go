package domain

type UserRepository interface {
	CreateUser(user *User) error
	FindById(id string) (*User, error)
	FindAll() ([]*User, error)
	UpdateUser(user *User) error
	DeleteById(id string) error
}