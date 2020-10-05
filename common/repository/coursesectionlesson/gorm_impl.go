package coursesectionlesson

import (
	"github.com/jinzhu/gorm"

	"tazapay.com/elearning/common/models"
)

// Gorm database interaction with gorm
type Gorm struct {
	db *gorm.DB
}

// getObj return operable orm object
func (g *Gorm) getObj() *gorm.DB {
	return g.db
}

// SetObj set operable orm object
func (g *Gorm) SetObj(i interface{}) {
	g.db = i.(*gorm.DB)
}

// GetByID fetch lesson by ID, preload lesson content
func (g *Gorm) GetByID(id int, columns ...string) (*models.CourseSectionLesson, error) {
	var lesson models.CourseSectionLesson
	err := g.getObj().Preload("Content").
		Where("id=?", id).
		Find(&lesson).Error
	return &lesson, err
}
