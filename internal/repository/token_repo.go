package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/usmaarn/yummeals_api/internal/model"
)

type TokenRepository struct {
	db *sqlx.DB
}

func NewTokenRepository(db *sqlx.DB) *TokenRepository {
	return &TokenRepository{db: db}
}

func (s *TokenRepository) Create(token *model.Token) error {
	row := s.db.QueryRow(`
	INSERT INTO tokens (user_id, access_token, refresh_token, access_token_expires_at, refresh_token_expires_at)
		VALUES ($1, $2, $3, $4, $5)
 	RETURNING id
	`,
		token.UserID, token.AccessToken, token.RefreshToken, token.AccessTokenExpiresAt, token.RefreshTokenExpiresAt)

	if row.Err() != nil {
		return row.Err()
	}

	err := row.Scan(&token.ID)
	return err
}

func (s *TokenRepository) FindByUserID(userID int64) (*model.Token, error) {
	var token model.Token
	err := s.db.Get(&token, "SELECT * FROM tokens WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (s *TokenRepository) FindByAccessToken(accessToken string) (*model.Token, error) {
	var token model.Token
	err := s.db.Get(&token, "SELECT * FROM tokens WHERE access_token = $1", accessToken)
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (s *TokenRepository) DeleteByUserID(userID int64) error {
	_, err := s.db.Exec("DELETE FROM tokens WHERE user_id = $1", userID)
	return err
}
