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
	user := &models.User{
		Email:        model.Email,
		Username:     model.Username,
		PasswordHash: model.Password,
	}

	err := u.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// FindByEmail implements ports.UserService.
func (u *UserService) FindByEmail(email string) (*models.User, error) {
	user, err := u.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetUserById implements ports.UserService.
func (u *UserService) GetUserById(id int) (*models.User, error) {
	user, err := u.repo.GetUserById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUser implements ports.UserService.
func (u *UserService) UpdateUser(model *models.UserUpdateRequest) (*models.User, error) {
	user := &models.User{
		Email:        model.Email,
		Username:     model.Username,
		PasswordHash: model.Password,
		Image:        model.Image,
		Bio:          model.Bio,
	}

	err := u.repo.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
