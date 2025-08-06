package repository

import (
	"api-echo/internal/domain"
	"database/sql"
	"time"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *domain.User) error {
	stmt, err := r.db.Prepare(`
		INSERT INTO users (id, name, email, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5)
	`)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		user.Id,
		user.Name,
		user.Email,
		user.CreatedAt,
		user.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) FindById(id string) (*domain.User, error) {
	var user domain.User
	var createdAt, updatedAt time.Time

	err := r.db.QueryRow(`
		SELECT id, name, email, created_at, updated_at
		FROM users
		WHERE id = $1
	`, id).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return nil, err
	}

	user.CreatedAt = createdAt
	user.UpdatedAt = updatedAt
	return &user, nil
}

func (r *UserRepository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	var createdAt, updatedAt time.Time

	err := r.db.QueryRow(`
		SELECT id, name, email, created_at, updated_at
		FROM users
		WHERE email = $1
	`, email).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return nil, err
	}

	user.CreatedAt = createdAt
	user.UpdatedAt = updatedAt
	return &user, nil
}

func (r *UserRepository) FindAll() ([]*domain.User, error) {
	rows, err := r.db.Query(`SELECT id, name, email, created_at, updated_at FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*domain.User

	for rows.Next() {
		var user domain.User
		var createdAt, updatedAt time.Time

		err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.Email,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, err
		}
		user.CreatedAt = createdAt
		user.UpdatedAt = updatedAt

		users = append(users, &user)
	}

	if err = rows.Err(); err != nil {
        return nil, err
    }

    return users, nil
}

func (r *UserRepository) UpdateUser(user *domain.User) error {
	_, err := r.db.Exec(`
		UPDATE users 
		SET name = $1, email = $2, updated_at = $3
		WHERE id = $4
	`,
		user.Name, 
		user.Email, 
		time.Now(),
		user.Id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) DeleteById(id string) error {
	_, err := r.db.Exec(`DELETE FROM users WHERE id = $1`, id)
	if err != nil {
		return err
	}

	return nil
}