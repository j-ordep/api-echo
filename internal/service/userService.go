package service

import (
	"api-echo/internal/domain"
)

type UserService struct {
	repository domain.UserRepository
}

func NewUserService(repository domain.UserRepository) *UserService {
	return &UserService{repository: repository}
}

func (s *UserService) CreateUser(user *domain.User) (*domain.User, error) {

	existingUser, err := s.repository.FindByEmail(user.Email)
   	if err != nil && err != domain.ErrUserNotFound {
    	return nil, err
    }

	if existingUser != nil {
		return nil, domain.ErrUserAlreadyExists
	}
	
	newUser := domain.NewUser(user.Name, user.Email)

	err = s.repository.CreateUser(newUser)
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

func (s *UserService) UpdateUser(id string, user *domain.User) (*domain.User, error) {

	newUser, err := s.repository.FindById(id)
	if err != nil {
		return nil, err
	}

	newUser.Name = user.Name
	newUser.Email = user.Email

	err = s.repository.UpdateUser(newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (s *UserService) DeleteById(id string) error {
	err := s.repository.DeleteById(id)
	if err != nil {
		return err
	}
	return nil
}