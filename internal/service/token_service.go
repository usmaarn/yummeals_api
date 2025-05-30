package service

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/usmaarn/yummeals_api/internal/model"
	"github.com/usmaarn/yummeals_api/internal/repository"
)

type TokenService struct {
	repo *repository.Repository
}

func NewTokenService(repo *repository.Repository) *TokenService {
	return &TokenService{repo: repo}
}

func (s *TokenService) CreateToken(userID int64) (*model.Token, error) {
	token := model.Token{
		UserID:                userID,
		AccessToken:           uuid.NewString(),
		RefreshToken:          uuid.NewString(),
		AccessTokenExpiresAt:  time.Now().Add(15 * time.Minute),    // 15 minutes
		RefreshTokenExpiresAt: time.Now().Add(30 * 24 * time.Hour), // 30 days
	}

	err := s.repo.TokenRepo.Create(&token)
	if err != nil {
		fmt.Printf("Error creating token: %v\n", err)
	}
	return &token, err
}
