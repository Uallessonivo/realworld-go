package ports

import "Github.com/Uallessonivo/RealWorld/internal/core/domain/models"

type UserRepository interface {
	GetUserById(id int) (*models.User, error)
	CreateUser(model *models.User) error
	UpdateUser(model *models.User) error
	FindByEmail(email string) (*models.User, error)
}

type UserService interface {
	GetUserById(id int) (*models.User, error)
	CreateUser(model *models.UserRegistrationRequest) (*models.User, error)
	UpdateUser(model *models.UserUpdateRequest) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
}
