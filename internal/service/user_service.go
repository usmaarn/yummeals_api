package service

import (
	"fmt"

	"github.com/usmaarn/yummeals_api/internal/dto"
	"github.com/usmaarn/yummeals_api/internal/model"
	"github.com/usmaarn/yummeals_api/internal/repository"

	"github.com/usmaarn/yummeals_api/pkg/utils/constants"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo}
}

func (s *UserService) FindByEmail(email string) (*model.User, error) {
	return s.repo.FindOneBy("email", email)
}

func (s *UserService) CreateUser(dto *dto.CreateUserRequest) (*model.User, error) {
	var newUser = model.User{
		Name:        fmt.Sprintf("%s %s", dto.FirstName, dto.LastName),
		Email:       dto.Email,
		PhoneNumber: dto.PhoneNumber,
		Password:    dto.Password,               // In a real application, ensure to hash the password
		Status:      constants.UserStatusActive, // Set the user status to active by default
	}

	err := s.repo.Save(&newUser)
	return &newUser, err
}
