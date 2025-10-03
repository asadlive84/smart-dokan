package domain

import (
	"smart-dokan/usersvc/internal/utility"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Password    string    `json:"password,omitempty"`
}

func NewUser(user *User) *User {
	hasedPasswd, err := utility.HashPassword(user.Password)
	if err != nil {
		return nil
	}
	return &User{
		Email:       user.Email,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		PhoneNumber: user.PhoneNumber,
		Password:    hasedPasswd,
	}
}
