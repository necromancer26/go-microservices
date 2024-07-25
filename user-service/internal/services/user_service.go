package services

import (
	"log"

	"github.com/necromancer26/go-microservices/user-service/internal/models"
	"github.com/necromancer26/go-microservices/user-service/internal/repository"
)

type UserService struct {
	userRepository *repository.RedisUserRepository
}

func NewUserService() *UserService {
	return &UserService{
		userRepository: repository.NewRedisUserRepository(),
	}
}
func (s *UserService) CreateUser(user *models.User) error {
	err := s.userRepository.Save(user)
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	return nil
}
func (s *UserService) GetAllUsers() ([]*models.User, error) {
	values, err := s.userRepository.FindAll()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)

	}
	return values, err
}
func (s *UserService) GetUserByName(name string) (*models.User, error) {
	return s.userRepository.FindByName(name)
}
func (s *UserService) DeleteUserById(id int) error {
	err := s.userRepository.Delete(id)
	if err != nil {
		log.Fatalf("Could not delete by id: %v", err)
		return err
	}
	return nil
}
func (s *UserService) UpdateUser(user *models.User) error {
	return s.userRepository.Update(user)
}
