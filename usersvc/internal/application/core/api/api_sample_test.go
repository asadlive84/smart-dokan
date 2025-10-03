package api_test

import (
	"smart-dokan/usersvc/internal/application/core/api"
	"smart-dokan/usersvc/internal/application/core/domain"
	"smart-dokan/usersvc/internal/ports/ports_mocks"
	"testing"

	"github.com/google/uuid"
)

func TestCreateUser1(t *testing.T) {
	mockDB := &ports_mocks.MockDB{}

	mockDB.InsertFunc = func(user *domain.User) (uuid.UUID, error) {
		return uuid.New(), nil
	}

	app := api.NewApplication(mockDB)

	user := &domain.User{
		FirstName:   "asad",
		LastName:    "sohel",
		Email:       "asad@me.com",
		PhoneNumber: "1223",
		Password:    "asd",
	}

	createdUser, err := app.CreateUser(user)
	if err != nil {
		t.Fatal(err)
	}

	if createdUser.ID == uuid.Nil {
		t.Fatal("expeced valid UUID")
	}

}
