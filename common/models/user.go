package models

import "time"

// User The details users enter during registration. Users can register using an email/mobile, and they can enroll in any course.
type User struct {
	ID                    int       `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	FirstName             string    `gorm:"column:first_name;type:varchar(64);not null" json:"first_name"`
	LastName              *string   `gorm:"column:last_name;type:varchar(64)" json:"last_name"`
	Mobile                string    `gorm:"unique;column:mobile;type:varchar(16);not null" json:"mobile"`
	Email                 string    `gorm:"unique;column:email;type:varchar(64);not null" json:"email"`
	NumOfCoursesEnrolled  *int      `gorm:"column:num_of_courses_enrolled;type:tinyint(4)" json:"num_of_courses_enrolled"`
	NumOfCoursesCompleted *int      `gorm:"column:num_of_courses_completed;type:tinyint(4)" json:"num_of_courses_completed"`
	CreatedAt             time.Time `gorm:"column:created_at;type:datetime;not null" json:"created_at"` // user registration date can be derived from this value
	UpdatedAt             time.Time `gorm:"column:updated_at;type:datetime;not null" json:"updated_at"`
}

// TableName return user database tablename
func (*User) TableName() string {
	return TableUser
}
