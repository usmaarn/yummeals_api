package storage

import "github.com/jmoiron/sqlx"

type Storage struct {
	db           *sqlx.DB
	UserStorage  *UserStorage
	TokenStorage *TokenStorage
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{
		db:           db,
		UserStorage:  NewUserStorage(db),
		TokenStorage: NewTokenStorage(db),
	}
}
