package models

import "time"

/*
	The actual content can be stored in some cloud storage bucket, but
	the metadata and the link to that content should be stored in the Course entity.
*/

/*
	The lesson can’t go to the saved state without going into the merged state first.
	Each author will be working in a local branch, again similar to the git process.
	To be ready to publish, the local branch should be merged with the master branch (status changes to merged)
	and on final confirmation/approval followed by save of master branch, status of the lesson changes to saved.
	All lessons need to be in a saved state for the course to be published.
*/

/*
	TimeRequiredInSec:
		The number of seconds estimated for users to complete the content.
		Time required to complete a section = Sum of the times required to complete each piece of related content.
		Time required to complete a course = Sum of the times required to complete each related section.

	ParentID:
		<= 0 - master branch
		> 0 - copy branch from master
*/

// CourseSectionLessonContent Maintains lesson contents, version, state, etc.
type CourseSectionLessonContent struct {
	ID                int       `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Link              string    `gorm:"column:link;type:varchar(128);not null" json:"link"`
	Version           string    `gorm:"column:version;type:varchar(8);not null" json:"version"`
	State             string    `gorm:"column:state;type:enum('created','draft','saved','merged');not null" json:"state"`
	TimeRequiredInSec int       `gorm:"column:time_required_in_sec;type:smallint(6);not null" json:"time_required_in_sec"` // The number of seconds estimated for users to complete the content.,Section Time - sum of time required to complete each related content.,Course Time - sum of time required to complete each related section.
	ParentID          int       `gorm:"column:parent_id;type:int(11);not null" json:"-"`                                   // 0 - master branch
	CreatedAt         time.Time `gorm:"column:created_at;type:datetime;not null" json:"-"`
	UpdatedAt         time.Time `gorm:"column:updated_at;type:datetime;not null" json:"-"`
}

// TableName return course_section_lesson_content database tablename
func (*CourseSectionLessonContent) TableName() string {
	return TableCourseSectionLessonContent
}

// state enum values
const (
	StateEnumCreated = "created"
	StateEnumDraft   = "draft"
	StateEnumSaved   = "saved"
	StateEnumMerged  = "merged"
)
