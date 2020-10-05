package models

import "time"

// Course Holds basic data about courses
type Course struct {
	ID                int       `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Title             string    `gorm:"index:title_idx;column:title;type:varchar(128);not null" json:"title"`
	IntroductionBrief string    `gorm:"column:introduction_brief;type:varchar(256);not null" json:"introduction_brief"` // A brief summary of an instructor’s educational background, work experience, skill sets, hobbies, etc.
	Fee               int       `gorm:"column:fee;type:smallint(6);not null" json:"fee"`
	Language          string    `gorm:"column:language;type:enum('English');not null" json:"language"`
	IsActive          *int      `gorm:"column:is_active;type:tinyint(4)" json:"-"` // whether the course is available for the user to enroll. null/0 - disabled
	CreatedAt         time.Time `gorm:"column:created_at;type:datetime;not null" json:"-"`
	UpdatedAt         time.Time `gorm:"column:updated_at;type:datetime;not null" json:"-"`
}

// TableName return course database tablename
func (*Course) TableName() string {
	return TableCourse
}
