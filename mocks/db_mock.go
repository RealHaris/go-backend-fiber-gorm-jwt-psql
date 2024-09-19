package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/RealHaris/go-fiber-backend/models"
)

type MockDB struct {
    mock.Mock
}

func (m *MockDB) Create(user *models.User) error {
    args := m.Called(user)
    return args.Error(0)
}

func (m *MockDB) FindByUsername(username string) (*models.User, error) {
    args := m.Called(username)
    return args.Get(0).(*models.User), args.Error(1)
}
