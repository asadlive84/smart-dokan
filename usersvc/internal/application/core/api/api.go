package api

import (
	"smart-dokan/usersvc/internal/application/core/domain"
	"smart-dokan/usersvc/internal/ports"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{db: db}
}

func (app *Application) CreateUser(user *domain.User) (*domain.User, error) {
	id, err := app.db.Insert(user)
	if err != nil {
		return nil, err
	}
	user.ID = id
	return user, nil
}

func (app *Application) GetUser(user *domain.User) (*domain.User, error) {
	_, err := app.db.Insert(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
