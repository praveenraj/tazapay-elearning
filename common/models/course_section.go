package models

import "time"

// Course should have 1 or more sections

/*
	Order: Move the entire section with all its lessons up or down.
*/

// CourseSection Stores the following records for each section of each course.
type CourseSection struct {
	ID              int       `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	CourseID        int       `gorm:"index:fk_course_section_course_id_idx;column:course_id;type:int(11);not null" json:"course_id"`
	Course          *Course   `gorm:"association_foreignkey:course_id;foreignkey:id" json:"course_list"` // Holds basic data about courses
	Title           string    `gorm:"column:title;type:varchar(128);not null" json:"title"`
	Order           int       `gorm:"column:order;type:tinyint(4)" json:"order"`                         // The position where this section will be available for the course
	NumOfReading    *int      `gorm:"column:num_of_reading;type:tinyint(4)" json:"num_of_reading"`       // The number of reading activities included in the section
	NumOfVideo      *int      `gorm:"column:num_of_video;type:tinyint(4)" json:"num_of_video"`           // The number of videos included in the section.
	NumOfAssignment *int      `gorm:"column:num_of_assignment;type:tinyint(4)" json:"num_of_assignment"` // The number of assignments associated with the section.
	IsActive        *int      `gorm:"column:is_active;type:tinyint(4)" json:"is_active"`                 // whether the section is available for the course. null/0 - disabled
	CreatedAt       time.Time `gorm:"column:created_at;type:datetime;not null" json:"created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at;type:datetime;not null" json:"updated_at"`
}

// TableName return course_section database tablename
func (*CourseSection) TableName() string {
	return TableCourseSection
}
