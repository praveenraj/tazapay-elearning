package coursesectionlessoncontent

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

// Create add new lesson content
func (m *Mock) Create(content *models.CourseSectionLessonContent) error {
	args := m.Called(content)
	return args.Error(constants.IntZero)
}

// GetByParentIDAndVersion fetch content by parent id & version
func (m *Mock) GetByParentIDAndVersion(parentID int, version string, columns ...string) (*models.CourseSectionLessonContent, error) {
	args := m.Called(parentID, version, columns)
	if args.Get(constants.IntZero) == nil {
		return nil, args.Error(constants.IntOne)
	}
	return args.Get(constants.IntZero).(*models.CourseSectionLessonContent), args.Error(constants.IntOne)
}

// Update modify lesson content record with struct ID value
func (m *Mock) Update(content *models.CourseSectionLessonContent, columns []string) error {
	args := m.Called(content, columns)
	return args.Error(constants.IntZero)
}
