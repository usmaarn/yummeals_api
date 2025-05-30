package service

import (
	"fmt"

	"github.com/usmaarn/yummeals_api/internal/dto"
	"github.com/usmaarn/yummeals_api/internal/model"
	"github.com/usmaarn/yummeals_api/internal/storage"
	"github.com/usmaarn/yummeals_api/pkg/utils/constants"
)

type UserService struct {
	storage *storage.Storage
}

func NewUserService(s *storage.Storage) *UserService {
	return &UserService{storage: s}
}

func (s *UserService) FindByEmail(email string) (*model.User, error) {
	return s.storage.UserStorage.FindOneBy("email", email)
}

func (s *UserService) CreateUser(dto *dto.CreateUserRequest) (*model.User, error) {
	var newUser = model.User{
		Name:        fmt.Sprintf("%s %s", dto.FirstName, dto.LastName),
		Email:       dto.Email,
		PhoneNumber: dto.PhoneNumber,
		Password:    dto.Password,               // In a real application, ensure to hash the password
		Status:      constants.UserStatusActive, // Set the user status to active by default
	}

	err := s.storage.UserStorage.Save(&newUser)
	return &newUser, err
}
