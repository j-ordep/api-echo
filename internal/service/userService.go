package service

import (
	"api-echo/internal/domain"
)

type UserService struct {
	repository domain.UserRepository
}

func NewService(repository domain.UserRepository) *UserService {
	return &UserService{repository: repository}
}

func (s *UserService) CreateUser(name string, email string) (*domain.User, error) {

	newUser := domain.NewUser(name, email)

	err := s.repository.CreateUser(newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (s *UserService) FindById(id string) (*domain.User, error){
	user, err := s.repository.FindById(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) FindAll() ([]*domain.User, error) {
	users, err:= s.repository.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) UpdateUser(id string, name string, email string) (*domain.User, error) {
	user, err := s.repository.FindById(id)
	if err != nil {
		return nil, err
	}

	user.Name = name
	user.Email = email

	err = s.repository.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) DeleteById(id string) error{
	err := s.repository.DeleteById(id)
	if err != nil {
		return err
	}
	return nil
}