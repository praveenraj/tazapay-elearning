package coursesectionlessoncontent

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

// Create add new lesson content
func (g *Gorm) Create(content *models.CourseSectionLessonContent) error {
	return g.getObj().Create(content).Error
}

// GetByParentIDAndVersion fetch content by parent id & version
func (g *Gorm) GetByParentIDAndVersion(parentID int, version string, columns ...string) (*models.CourseSectionLessonContent, error) {
	var content models.CourseSectionLessonContent
	err := g.getObj().
		Where("parent_id=?", parentID).
		Where("version=?", version).
		Find(&content).Error
	return &content, err
}

// Update modify lesson content record with struct ID value
func (g *Gorm) Update(content *models.CourseSectionLessonContent, columns []string) error {
	return g.getObj().Model(content).Select(columns).Updates(content).Error
}
