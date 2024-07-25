package repository

import "github.com/necromancer26/go-microservices/user-service/internal/models"

type UserRepository interface {
	FindAll() ([]*models.User, error)
	FindByID(id int) (*models.User, error)
	FindByName(name string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Save(user *models.User) error
	Update(user *models.User) error
	Delete(id int) error
}
