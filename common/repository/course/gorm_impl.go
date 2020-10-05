package course

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

// GetAll fetch all the distinct courses.
// course, section, lesson should be active and content should be in saved state
func (g *Gorm) GetAll(columns ...string) ([]*models.Course, error) {
	var courses []*models.Course
	db := g.getObj().Joins("JOIN "+models.TableCourseSection+" section on "+models.TableCourse+".id=section.course_id").
		Joins("JOIN "+models.TableCourseSectionLesson+" lesson on section.id=lesson.course_section_id").
		Joins("JOIN "+models.TableCourseSectionLessonContent+" content on content.id = lesson.content_id").
		Where(models.TableCourse+".is_active=?", 1).
		Where("section.is_active=?", 1).
		Where("lesson.is_active=?", 1).
		Where("content.state=?", "saved").
		Group(models.TableCourse + ".id").
		Find(&courses)
	return courses, db.Error
}
