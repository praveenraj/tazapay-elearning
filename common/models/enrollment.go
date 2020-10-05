package models

import "time"

// Enrollment Stores user-course information
type Enrollment struct {
	ID                 int       `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	UserID             int       `gorm:"index:fk_enrollment_user_id_idx;column:user_id;type:int(11);not null" json:"user_id"`
	User               *User     `gorm:"association_foreignkey:user_id;foreignkey:id" json:"user_list"` // The details users enter during registration. Users can register using an email/mobile, and they can enroll in any course.
	CourseID           int       `gorm:"index:fk_enrollment_course_id_idx;column:course_id;type:int(11);not null" json:"course_id"`
	Course             *Course   `gorm:"association_foreignkey:course_id;foreignkey:id" json:"course_list"`       // Holds basic data about courses
	IsPaidSubscription *int      `gorm:"column:is_paid_subscription;type:tinyint(4)" json:"is_paid_subscription"` // Whether the enrollment is free or paid. null/0 - free
	CreatedAt          time.Time `gorm:"column:created_at;type:datetime;not null" json:"created_at"`              // The date when the enrollment took place.
}
