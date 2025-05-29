package service

import "github.com/usmaarn/yummeals_api/internal/storage"

type Service struct {
	Auth *AuthService
}

func RegisterServices(s *storage.Storage) *Service {
	return &Service{
		Auth: NewAuthService(s),
	}
}
