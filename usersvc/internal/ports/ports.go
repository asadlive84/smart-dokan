package ports

import (
	"smart-dokan/usersvc/internal/application/core/domain"

	"github.com/google/uuid"
)

type APIPort interface {
	CreateUser(user *domain.User) (*domain.User, error)
	GetUser(user *domain.User) (*domain.User, error)
}

type DBPort interface {
	Insert(user *domain.User) (uuid.UUID, error)
}
