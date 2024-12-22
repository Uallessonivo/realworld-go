package ports

import "Github.com/Uallessonivo/RealWorld/internal/core/domain/models"

type UserRepository interface {
	GetUserById(id int) (*models.User, error)
	CreateUser(model *models.User) error
	UpdateUser(model *models.User) error
	FindByEmail(email string) (*models.User, error)
}
