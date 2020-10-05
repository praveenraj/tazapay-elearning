package models

import "time"

/*
	Avoiding a direct connection with the course or author table, we minimize fake feedback submissions.
	This way, only users can submit feedback, and only after they complete the course.
*/

// Feedback Stores feedback and reviews submitted by user for the enrolled course
type Feedback struct {
	ID           int         `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	EnrollmentID int         `gorm:"index:fk_feedback_enrollment_id_idx;column:enrollment_id;type:int(11);not null" json:"enrollment_id"`
	Enrollment   *Enrollment `gorm:"association_foreignkey:enrollment_id;foreignkey:id" json:"enrollment_list"` // Stores user-course information
	RatingScore  int         `gorm:"column:rating_score;type:tinyint(4);not null" json:"rating_score"`          // This value ranges from 1 (worst) to 5 (best).
	Message      *string     `gorm:"column:message;type:varchar(256)" json:"message"`
	CreatedAt    time.Time   `gorm:"column:created_at;type:datetime;not null" json:"created_at"`
}
