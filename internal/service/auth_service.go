package service

import (
	"errors"
	"fmt"

	"github.com/usmaarn/yummeals_api/internal/dto"
	"github.com/usmaarn/yummeals_api/internal/model"
	"github.com/usmaarn/yummeals_api/pkg/utils/constants"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userService  *UserService
	tokenService *TokenService
}

func NewAuthService(userService *UserService, tokenService *TokenService) *AuthService {
	return &AuthService{userService, tokenService}
}

func (s *AuthService) Register(dto *dto.CreateUserRequest) (*model.Token, error) {
	if dto == nil {
		return nil, fmt.Errorf("user data cannot be nil")
	}
	user, err := s.userService.CreateUser(dto)
	if err != nil {
		return nil, err
	}

	return s.tokenService.CreateToken(user.ID)
}
func (s *AuthService) Login(dto dto.LoginUserRequest) (*model.Token, error) {
	user, err := s.userService.FindByEmail(dto.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if user.Status != constants.UserStatusActive {
		return nil, fmt.Errorf("user is not active")
	}

	return s.tokenService.CreateToken(user.ID)
}
