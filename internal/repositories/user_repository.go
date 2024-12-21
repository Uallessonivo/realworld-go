package repositories

import (
	"Github.com/Uallessonivo/RealWorld/internal/core/domain/models"
	"Github.com/Uallessonivo/RealWorld/internal/core/domain/ports"
	"gorm.io/gorm"
)

type UserRepositoryDb struct {
	Db *gorm.DB
}

func NewUserRepository(Db *gorm.DB) ports.UserRepository {
	return &UserRepositoryDb{Db}
}

// CreateUser implements ports.UserRepository.
func (u *UserRepositoryDb) CreateUser(model *models.User) error {
	panic("unimplemented")
}

// FindByEmail implements ports.UserRepository.
func (u *UserRepositoryDb) FindByEmail(email string) (*models.User, error) {
	panic("unimplemented")
}

// GetUser implements ports.UserRepository.
func (u *UserRepositoryDb) GetUser(id int) (*models.User, error) {
	panic("unimplemented")
}

// UpdateUser implements ports.UserRepository.
func (u *UserRepositoryDb) UpdateUser(model *models.User) error {
	panic("unimplemented")
}
