package services

import (
	"Github.com/Uallessonivo/RealWorld/internal/core/domain/models"
	"Github.com/Uallessonivo/RealWorld/internal/core/domain/ports"
)

type UserService struct {
	repo ports.UserRepository
}

func NewUserService(u ports.UserRepository) ports.UserService {
	return &UserService{
		repo: u,
	}
}

// CreateUser implements ports.UserService.
func (u *UserService) CreateUser(model *models.UserRegistrationRequest) (*models.User, error) {
	panic("unimplemented")
}

// FindByEmail implements ports.UserService.
func (u *UserService) FindByEmail(email string) (*models.User, error) {
	panic("unimplemented")
}

// GetUserById implements ports.UserService.
func (u *UserService) GetUserById(id int) (*models.User, error) {
	panic("unimplemented")
}

// UpdateUser implements ports.UserService.
func (u *UserService) UpdateUser(model *models.UserUpdateRequest) (*models.User, error) {
	panic("unimplemented")
}
