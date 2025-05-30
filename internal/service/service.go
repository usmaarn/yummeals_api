package service

import (
	"github.com/usmaarn/yummeals_api/internal/repository"
)

type Service struct {
	AuthService  *AuthService
	UserService  *UserService
	TokenService *TokenService
}

func RegisterServices(r *repository.Repository) *Service {
	userService := NewUserService(r)
	tokenService := NewTokenService(r)
	authService := NewAuthService(userService, tokenService)

	return &Service{
		AuthService:  authService,
		UserService:  userService,
		TokenService: tokenService,
	}
}
