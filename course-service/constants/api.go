package constants

// api endpoints
const (
	PathCourses       = "/courses"
	PathLessonContent = "/courses/{course_id}/sections/{section_id}/lessons/{lesson_id}"
)

// api path params
const (
	PathParamCourseID = "course_id"
	PathParamSection  = "section_id"
	PathParamLessonID = "lesson_id"
)

// api actions
const (
	ActionDraft = "draft"
	ActionMerge = "merge"
	ActionSave  = "save"
)
