package domain

import (
	"fmt"
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

func NewUser(user *User) (*User, error) {
	hasedPasswd, err := utility.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	if user.FirstName == "" || user.LastName == "" || user.PhoneNumber == "" || user.Email == ""  || user.Password == ""{
		return nil, fmt.Errorf("%+v", "field are missing, please check")
	}
	return &User{
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Password:    hasedPasswd,
	}, nil
}
