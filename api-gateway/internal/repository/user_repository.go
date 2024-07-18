package repository

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// Add methods to interact with the user data source
