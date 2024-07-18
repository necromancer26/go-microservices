package repository

type AuthRepository struct{}

func NewAuthRepository() *AuthRepository {
	return &AuthRepository{}
}

// Add methods to interact with the authentication data source
