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
func (s *UserService) GetAllUsers() (map[string]string, error) {
	values, err := s.userRepository.FindAll()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)

	}
	return values, err
}
