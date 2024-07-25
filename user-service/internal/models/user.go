package models

type User struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	FamilyName string `json:"family_name"`
	Email      string `json:"email"`
}

func NewUser(name, familyName, email string) *User {
	return &User{
		Name:       name,
		FamilyName: familyName,
		Email:      email,
	}
}
