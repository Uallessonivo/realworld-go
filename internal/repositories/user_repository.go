package repositories

import (
	"strings"

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
	if err := u.Db.Create(&model).Error; err != nil {
		return err
	}
	return nil
}

// FindByEmail implements ports.UserRepository.
func (u *UserRepositoryDb) FindByEmail(email string) (*models.User, error) {
	var model models.User
	if err := u.Db.Where("LOWER(email) = ?", strings.ToLower(email)).First(&model).Error; err != nil {
		return nil, err
	}
	return &model, nil
}

// GetUserById implements ports.UserRepository.
func (u *UserRepositoryDb) GetUserById(id int) (*models.User, error) {
	var model models.User
	if err := u.Db.Where("ID = ?", id).First(&model).Error; err != nil {
		return nil, err
	}
	return &model, nil
}

// UpdateUser implements ports.UserRepository.
func (u *UserRepositoryDb) UpdateUser(model *models.User) error {
	err := u.Db.Model(models.User{}).Where("ID = ?", model.ID).UpdateColumns(model).Error
	if err != nil {
		return err
	}
	return nil
}
