package models

import "time"

/*
	Keep aggregated columns like num_of_enrolled_students and num_of_reviews separate.
	When there are millions of records in a db, it is always better to store aggregated values within the database itself.
	But it is not necessary to keep updates in real time.
	These are just columns to hold aggregated values, and these values may be computed once a day and stored per convenience, i.e. each night.
*/

// Author Information regarding individual instructors. Not everyone is allowed to register as an author, only authorized individuals can submit course content.
type Author struct {
	ID                    int       `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	FirstName             string    `gorm:"column:first_name;type:varchar(64);not null" json:"first_name"`
	LastName              *string   `gorm:"column:last_name;type:varchar(64)" json:"last_name"`
	Mobile                string    `gorm:"unique;column:mobile;type:varchar(16);not null" json:"mobile"`
	Email                 string    `gorm:"unique;column:email;type:varchar(64);not null" json:"email"`
	Qualification         string    `gorm:"column:qualification;type:enum('Ph.D','MSCE','B.E','B.SC');not null" json:"qualification"`
	IntroductionBrief     string    `gorm:"column:introduction_brief;type:varchar(256);not null" json:"introduction_brief"` // A brief summary of an instructor’s educational background, work experience, skill sets, hobbies, etc.
	Image                 string    `gorm:"column:image;type:varchar(128);not null" json:"image"`                           // image url where it is originally saved
	NumOfPublishedCourses *int      `gorm:"column:num_of_published_courses;type:tinyint(4)" json:"num_of_published_courses"`
	NumOfEnrolledStudents *int      `gorm:"column:num_of_enrolled_students;type:smallint(6)" json:"num_of_enrolled_students"` // Not necessary to keep updates in real time. Run a scheduler job to find once a day.
	AverageReviewRating   *float64  `gorm:"column:average_review_rating;type:float(3,1)" json:"average_review_rating"`        // Not necessary to keep updates in real time. Run a scheduler job to find once a day.
	NumOfReviews          *int      `gorm:"column:num_of_reviews;type:smallint(6)" json:"num_of_reviews"`                     // Not necessary to keep updates in real time. Run a scheduler job to find once a day.
	CreatedAt             time.Time `gorm:"column:created_at;type:datetime;not null" json:"created_at"`                       // author registration date can be derived from this value
	UpdatedAt             time.Time `gorm:"column:updated_at;type:datetime;not null" json:"updated_at"`
}
