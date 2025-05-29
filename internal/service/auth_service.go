package service

import "github.com/usmaarn/yummeals_api/internal/storage"

type AuthService struct {
	storage *storage.Storage
}

func NewAuthService(s *storage.Storage) *AuthService {
	return &AuthService{s}
}
