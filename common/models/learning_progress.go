package models

import "time"

// LearningProgress Holds details about user's progress through a course.
type LearningProgress struct {
	ID                    int                  `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	EnrollmentID          int                  `gorm:"index:fk_learning_progress_enrollment_id_idx;column:enrollment_id;type:int(11);not null" json:"enrollment_id"`
	Enrollment            *Enrollment          `gorm:"association_foreignkey:enrollment_id;foreignkey:id" json:"enrollment_list"`                                                      // Stores user-course information
	CourseSectionLessonID int                  `gorm:"index:fk_learning_progress_lesson_id_idx;column:course_section_lesson_id;type:int(11);not null" json:"course_section_lesson_id"` // The ID for the section lesson related to this record
	CourseSectionLesson   *CourseSectionLesson `gorm:"association_foreignkey:course_section_lesson_id;foreignkey:id" json:"course_section_lesson_list"`                                // Holds lesson details for each section.
	Status                string               `gorm:"column:status;type:enum('P','C');not null" json:"status"`                                                                        // P - in progress, C - complete
	CreatedAt             time.Time            `gorm:"column:created_at;type:datetime;not null" json:"created_at"`                                                                     // The timestamp when the user starts the lesson.
	UpdatedAt             time.Time            `gorm:"column:updated_at;type:datetime;not null" json:"updated_at"`                                                                     // The timestamp when the user finishes the lesson recently (if a user goes through the lesson multiple times).
}
