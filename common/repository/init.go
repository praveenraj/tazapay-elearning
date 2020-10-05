package repository

import (
	"tazapay.com/elearning/common/driver"
	"tazapay.com/elearning/common/repository/author"
	"tazapay.com/elearning/common/repository/config"
	"tazapay.com/elearning/common/repository/course"
	"tazapay.com/elearning/common/repository/courseauthor"
	"tazapay.com/elearning/common/repository/coursesection"
	"tazapay.com/elearning/common/repository/coursesectionlesson"
	"tazapay.com/elearning/common/repository/coursesectionlessoncontent"
	"tazapay.com/elearning/common/repository/enrollment"
	"tazapay.com/elearning/common/repository/feedback"
	"tazapay.com/elearning/common/repository/learningprogress"
	"tazapay.com/elearning/common/repository/payment"
	"tazapay.com/elearning/common/repository/user"
)

// NewAuthorDAO initialize author dao layer
func NewAuthorDAO() author.DAO {
	obj := getImplObj()
	dao := author.GetImpl()[driver.GetDB().GetOrm()]
	dao.SetObj(obj)
	return dao
}

// NewConfigDAO initialize config dao layer
func NewConfigDAO() config.DAO {
	obj := getImplObj()
	dao := config.GetImpl()[driver.GetDB().GetOrm()]
	dao.SetObj(obj)
	return dao
}

// NewCourseDAO initialize course dao layer
func NewCourseDAO() course.DAO {
	obj := getImplObj()
	dao := course.GetImpl()[driver.GetDB().GetOrm()]
	dao.SetObj(obj)
	return dao
}

// NewCourseAuthorDAO initialize course_author dao layer
func NewCourseAuthorDAO() courseauthor.DAO {
	obj := getImplObj()
	dao := courseauthor.GetImpl()[driver.GetDB().GetOrm()]
	dao.SetObj(obj)
	return dao
}

// NewCourseSectionDAO initialize course_section dao layer
func NewCourseSectionDAO() coursesection.DAO {
	obj := getImplObj()
	dao := coursesection.GetImpl()[driver.GetDB().GetOrm()]
	dao.SetObj(obj)
	return dao
}

// NewCourseSectionLessonDAO initialize course_section_lesson dao layer
func NewCourseSectionLessonDAO() coursesectionlesson.DAO {
	obj := getImplObj()
	dao := coursesectionlesson.GetImpl()[driver.GetDB().GetOrm()]
	dao.SetObj(obj)
	return dao
}

// NewCourseSectionLessonContentDAO initialize course_section_lesson_content dao layer
func NewCourseSectionLessonContentDAO() coursesectionlessoncontent.DAO {
	obj := getImplObj()
	dao := coursesectionlessoncontent.GetImpl()[driver.GetDB().GetOrm()]
	dao.SetObj(obj)
	return dao
}

// NewEnrollmentDAO initialize enrollment dao layer
func NewEnrollmentDAO() enrollment.DAO {
	obj := getImplObj()
	dao := enrollment.GetImpl()[driver.GetDB().GetOrm()]
	dao.SetObj(obj)
	return dao
}

// NewFeedbackDAO initialize feedback dao layer
func NewFeedbackDAO() feedback.DAO {
	obj := getImplObj()
	dao := feedback.GetImpl()[driver.GetDB().GetOrm()]
	dao.SetObj(obj)
	return dao
}

// NewLearningProgressDAO initialize learning_progress dao layer
func NewLearningProgressDAO() learningprogress.DAO {
	obj := getImplObj()
	dao := learningprogress.GetImpl()[driver.GetDB().GetOrm()]
	dao.SetObj(obj)
	return dao
}

// NewPaymentDAO initialize payment dao layer
func NewPaymentDAO() payment.DAO {
	obj := getImplObj()
	dao := payment.GetImpl()[driver.GetDB().GetOrm()]
	dao.SetObj(obj)
	return dao
}

// NewUserDAO initialize user dao layer
func NewUserDAO() user.DAO {
	obj := getImplObj()
	dao := user.GetImpl()[driver.GetDB().GetOrm()]
	dao.SetObj(obj)
	return dao
}

// getImplObj return orm or any relavant implementation obj
func getImplObj() interface{} {
	// switch w.r.t global driver's DB orm value
	switch driver.GetDB().GetOrm() {
	case driver.GOrm: // gorm
		return driver.GetDB().GetGOrm()

	default: // mock
		return driver.GetDB()
	}
}
