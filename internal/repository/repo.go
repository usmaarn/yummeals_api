package repository

import "github.com/jmoiron/sqlx"

type Repository struct {
	db        *sqlx.DB
	UserRepo  *UserRepository
	TokenRepo *TokenRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db:        db,
		UserRepo:  NewUserRepository(db),
		TokenRepo: NewTokenRepository(db),
	}
}
