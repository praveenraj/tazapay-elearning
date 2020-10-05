package coursesectionlesson

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

// GetByID fetch lesson by ID, preload lesson content
func (m *Mock) GetByID(id int, columns ...string) (*models.CourseSectionLesson, error) {
	args := m.Called(id, columns)
	return args.Get(constants.IntZero).(*models.CourseSectionLesson), args.Error(constants.IntOne)
}
