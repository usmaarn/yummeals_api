package storage

import (
	"github.com/jmoiron/sqlx"
	"github.com/usmaarn/yummeals_api/internal/model"
)

type TokenStorage struct {
	db *sqlx.DB
}

func NewTokenStorage(db *sqlx.DB) *TokenStorage {
	return &TokenStorage{db: db}
}

func (s *TokenStorage) Create(token *model.Token) error {
	return s.db.Get(&token, `
	INSERT INTO tokens (user_id, access_token, refresh_token, access_token_expires_at, refresh_token_expires_at)
		VALUES ($1, $2, $3, $4, $5)
 	RETURNING *
	`,
		token.UserID, token.AccessToken, token.RefreshToken, token.AccessTokenExpiresAt, token.RefreshTokenExpiresAt)
}

func (s *TokenStorage) FindByUserID(userID int64) (*model.Token, error) {
	var token model.Token
	err := s.db.Get(&token, "SELECT * FROM tokens WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (s *TokenStorage) FindByAccessToken(accessToken string) (*model.Token, error) {
	var token model.Token
	err := s.db.Get(&token, "SELECT * FROM tokens WHERE access_token = $1", accessToken)
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (s *TokenStorage) DeleteByUserID(userID int64) error {
	_, err := s.db.Exec("DELETE FROM tokens WHERE user_id = $1", userID)
	return err
}
