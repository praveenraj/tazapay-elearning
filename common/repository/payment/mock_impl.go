package payment

import "github.com/stretchr/testify/mock"

// Mock database interaction for unit testing
type Mock struct {
	mock.Mock
}

// SetObj set operable orm object
func (m *Mock) SetObj(i interface{}) {
	m.Called(i)
}
