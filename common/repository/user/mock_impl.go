package user

import (
	"github.com/stretchr/testify/mock"

	"tazapay.com/elearning/common/constants"
	"tazapay.com/elearning/common/models"
)

// Mock database interaction for unit testing
type Mock struct {
	mock.Mock
}

// SetObj set operable orm object
func (m *Mock) SetObj(i interface{}) {
	m.Called(i)
}

// Create add new user
func (m *Mock) Create(user *models.User) error {
	args := m.Called(user)
	return args.Error(constants.IntZero)
}
