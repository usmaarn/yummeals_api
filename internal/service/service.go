package service

import "github.com/usmaarn/yummeals_api/internal/storage"

type Service struct {
	AuthService  *AuthService
	UserService  *UserService
	TokenService *TokenService
}

func RegisterServices(s *storage.Storage) *Service {
	userService := NewUserService(s)
	tokenService := NewTokenService(s)
	authService := NewAuthService(userService, tokenService)

	return &Service{
		AuthService:  authService,
		UserService:  userService,
		TokenService: tokenService,
	}
}
