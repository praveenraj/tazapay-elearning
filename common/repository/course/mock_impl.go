package course

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

// GetAll fetch all the distinct courses.
// course, section, lesson should be active and content should be in saved state
func (m *Mock) GetAll(columns ...string) ([]*models.Course, error) {
	args := m.Called(columns)
	return args.Get(constants.IntZero).([]*models.Course), args.Error(constants.IntOne)
}
