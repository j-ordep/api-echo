package repository

import (
	"api-echo/internal/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *domain.User) error {
    return r.db.Create(user).Error
}

func (r *UserRepository) FindById(id string) (*domain.User, error) {
    var user domain.User
    result := r.db.First(&user, "id = ?", id)
    if result.Error != nil {
        return nil, result.Error
    }
    return &user, nil
}

func (r *UserRepository) FindByEmail(email string) (*domain.User, error) {
    var user domain.User
    result := r.db.First(&user, "email = ?", email)
    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            return nil, domain.ErrUserNotFound
        }
        return nil, result.Error
    }
    return &user, nil
}

func (r *UserRepository) FindAll() ([]*domain.User, error) {
    var users []*domain.User
    result := r.db.Find(&users)
    if result.Error != nil {
        return nil, result.Error
    }
    return users, nil
}

func (r *UserRepository) UpdateUser(user *domain.User) error {
    return r.db.Save(user).Error
}

func (r *UserRepository) DeleteById(id string) error {
    result := r.db.Delete(&domain.User{}, "id = ?", id)
    return result.Error
}