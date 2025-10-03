package api_test

import (
	"smart-dokan/usersvc/internal/application/core/api"
	"smart-dokan/usersvc/internal/application/core/domain"
	"smart-dokan/usersvc/internal/ports/ports_mocks"
	"testing"

	"github.com/google/uuid"
)

func TestCreateUser(t *testing.T) {

	newMockDB := func(insertFunc func(*domain.User) (uuid.UUID, error)) *api.Application {
		a := &ports_mocks.MockDB{
			InsertFunc: insertFunc,
		}
		return api.NewApplication(a)
	}

	t.Run("create_test", func(t *testing.T) {
		insertFunc := func(*domain.User) (uuid.UUID, error) {
			return uuid.New(), nil
		}
		app := newMockDB(insertFunc)

		u1 := &domain.User{
			FirstName:   "asad",
			LastName:    "sohel",
			Email:       "asad@me.com",
			PhoneNumber: "888",
			Password:    "123",
		}

		createdUser, err := app.CreateUser(u1)
		if err != nil {
			t.Fatal(err)
		}

		if createdUser.ID == uuid.Nil {
			t.Fatal("expedted a valid uuid")
		}
	})

	t.Run("requried_fields", func(t *testing.T) {
		insertFunc := func(i *domain.User) (uuid.UUID, error) {
			return uuid.New(), nil
		}

		app := newMockDB(insertFunc)

		u1 := &domain.User{
			FirstName:   "asad",
			LastName:    "sohel",
			Email:       "asad@me.com",
			PhoneNumber: "888",
			Password:    "",
		}

		_, err := app.CreateUser(u1)
		if err == nil {
			t.Fatal(err)
		}
	})
}
