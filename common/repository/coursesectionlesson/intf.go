package coursesectionlesson

import (
	"tazapay.com/elearning/common/driver"
	"tazapay.com/elearning/common/models"
)

// DAO abstract layer for course_section_lesson table interactions
type DAO interface {
	SetObj(i interface{})
	GetByID(id int, columns ...string) (*models.CourseSectionLesson, error)
}

// GetImpl return dao implemented structs
func GetImpl() map[string]DAO {
	return map[string]DAO{
		driver.GOrm: new(Gorm),
		driver.Mock: getMock(),
	}
}

// getter & setter for mock object
var _m *Mock

func getMock() *Mock {
	return _m
}

// SetMock set mocked functions object
func SetMock(m *Mock) {
	_m = m
}
