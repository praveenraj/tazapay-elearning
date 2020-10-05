package models

import "time"

// Each section should have 1 or more lessons

/*
	Order: Move lesson blocks up or down
*/

// CourseSectionLesson Holds lesson details for each section.
type CourseSectionLesson struct {
	ID              int                         `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	CourseSectionID int                         `gorm:"index:fk_course_section_lesson_couser_section_id_idx;column:course_section_id;type:int(11);not null" json:"course_section_id"`
	CourseSection   *CourseSection              `gorm:"association_foreignkey:course_section_id;foreignkey:id" json:"course_section_list"`                       // Stores the following records for each section of each course.
	ContentID       int                         `gorm:"index:fk_course_section_lesson_content_id_idx;column:content_id;type:int(11);not null" json:"content_id"` // latest saved state content or admin author selected version(if 1 or more saved state contents) will be used
	Content         *CourseSectionLessonContent `gorm:"association_foreignkey:content_id;foreignkey:id" json:"content"`                                          // Maintains lesson contents, version, state, etc.
	Order           int                         `gorm:"column:order;type:tinyint(4)" json:"order"`                                                               // The position where this lesson should available for the section                                                          // The position where this section will be available for the course
	IsMandatory     *int                        `gorm:"column:is_mandatory;type:tinyint(4)" json:"is_mandatory"`                                                 // Whether completing the lesson is mandatory for the course. null/0 - non mandate
	IsOpenForFree   *int                        `gorm:"column:is_open_for_free;type:tinyint(4)" json:"is_open_for_free"`                                         // whether the lesson is available to users using a free enrollment. null/0 - paid content
	IsActive        *int                        `gorm:"column:is_active;type:tinyint(4)" json:"-"`                                                               // whether the lesson is available for the section. null/0 - disabled
	CreatedAt       time.Time                   `gorm:"column:created_at;type:datetime;not null" json:"-"`
	UpdatedAt       time.Time                   `gorm:"column:updated_at;type:datetime;not null" json:"-"`
}

// TableName return course_section_lesson database tablename
func (*CourseSectionLesson) TableName() string {
	return TableCourseSectionLesson
}
