package models

import "time"

/*
	Courses may be created by 1 or more authors working together.
	Each has his/her own access to the course (permission).
	Imagine a “git” like process where they would be able to merge their changes.
*/

// CourseAuthor Courses may be created by 1 or more authors working together. Each has his/her own access to the course.
type CourseAuthor struct {
	ID         int       `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	CourseID   int       `gorm:"index:fk_course_author_course_id_idx;column:course_id;type:int(11);not null" json:"course_id"`
	Course     *Course   `gorm:"association_foreignkey:course_id;foreignkey:id" json:"course_list"` // Holds basic data about courses
	AuthorID   int       `gorm:"index:fk_course_author_author_id_idx;column:author_id;type:int(11);not null" json:"author_id"`
	Author     *Author   `gorm:"association_foreignkey:author_id;foreignkey:id" json:"author_list"`                  // Information regarding individual instructors. Not everyone is allowed to register as an author, only authorized individuals can submit course content.
	Permission string    `gorm:"column:permission;type:enum('admin','maintain','write');not null" json:"permission"` // admin - add/view/edit/delete/merge/publish,maintain - view/edit/merge/publish,write - view/edit
	IsActive   *int      `gorm:"column:is_active;type:tinyint(4)" json:"is_active"`                                  // whether the author is authorize to handle the course. null/0 - disabled
	CreatedAt  time.Time `gorm:"column:created_at;type:datetime;not null" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at;type:datetime;not null" json:"updated_at"`
}
