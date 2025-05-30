package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/usmaarn/yummeals_api/internal/model"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db}
}

func (s *UserRepository) FindOneBy(column string, value any) (*model.User, error) {
	user := &model.User{}
	query := fmt.Sprintf("SELECT * FROM users WHERE %s = $1", column)
	err := s.db.Get(user, query, value)
	if err != nil {
		return nil, fmt.Errorf("error finding user by %s: %w", column, err)
	}
	return user, nil
}

func (s *UserRepository) Delete(id int64) error {
	// Implement the logic to delete a user by ID from the database
	// This might involve executing an SQL DELETE statement
	return nil
}

func (s *UserRepository) FindAll(users []*model.User) error {
	return s.db.Select(users, "SELECT * FROM users WHERE deleted_at IS NULL")
}

func (s *UserRepository) Count() (int64, error) {
	var count int64
	err := s.db.Get(&count, "SELECT COUNT(*) FROM users")
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (s *UserRepository) Exists(email string) (bool, error) {
	if count, err := s.Count(); err != nil {
		return false, err
	} else {
		return count > 0, nil
	}
}

func (s *UserRepository) Save(user *model.User) error {
	var err error

	if user.ID == 0 {
		err = s.db.Get(&user, "INSERT INTO users (name, email, phone_number, password, status) VALUES ($1, $2, $3, $4, %5) RETURNING *",
			user.Name, user.Email, user.PhoneNumber, user.Password, user.Status)
	} else {
		err = s.db.Get(&user, "UPDATE users SET name = $1, email = $2, phone_number = $3 WHERE id = $4 RETURNING *",
			user.Name, user.Email, user.PhoneNumber, user.ID)
	}
	if err != nil {
		return fmt.Errorf("error saving user: %w", err)
	}
	return nil
}

func (s *UserRepository) SoftDelete(id int64) error {
	// Implement the logic to soft delete a user by ID from the database
	_, err := s.db.Exec("UPDATE users SET deleted_at = NOW() WHERE id = $1", id)
	return err
}
