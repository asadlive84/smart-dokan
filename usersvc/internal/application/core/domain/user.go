package domain

import "github.com/google/uuid"

type User struct {
	ID          uuid.UUID `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Password    string    `json:"password,omitempty"`
}

func NewUser(id uuid.UUID, username, email, password string) *User {
	return &User{
		ID:       id,
		Email:    email,
		Password: password,
	}
}
