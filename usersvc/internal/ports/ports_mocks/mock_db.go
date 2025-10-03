package ports_mocks

import (
	"smart-dokan/usersvc/internal/application/core/domain"

	"github.com/google/uuid"
)

type MockDB struct {
	InsertFunc func(user *domain.User) (uuid.UUID, error)
}

func (m *MockDB) Insert(user *domain.User) (uuid.UUID, error) {
	if m.InsertFunc != nil {
		return m.InsertFunc(user)
	}
	return uuid.New(), nil
}
