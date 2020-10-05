package controllers

import (
	"net/http"

	"tazapay.com/elearning/svc/course/constants"
	"tazapay.com/elearning/svc/course/logic"

	"tazapay.com/elearning/common/request"
	"tazapay.com/elearning/common/responses"
)

// API API gateway controller for the service
func API(payload *request.Payload) *responses.APIResponse {
	// resource path routing
	switch payload.GetAPI().Resource {
	case constants.PathCourses: // /courses
		return handleCourses(payload)

	case constants.PathLessonContent: // /courses/{course_id}/sections/{section_id}/lessons/{lesson_id}
		return handleLessonContent(payload)

	default: // invalid api path
		return responses.APIBadRequest(nil)
	}
}

// handleCourses handles /courses request
func handleCourses(payload *request.Payload) *responses.APIResponse {
	// method routing
	switch payload.GetAPI().HTTPMethod {
	case http.MethodGet: // GET
		return responses.API(logic.GetAllCourses(payload))

	default: // invalid http method
		return responses.APIMethodNotAllowed()
	}
}

// handleLessonContent handles /courses/{course_id}/sections/{section_id}/lessons/{lesson_id} request
func handleLessonContent(payload *request.Payload) *responses.APIResponse {
	// method routing
	switch payload.GetAPI().HTTPMethod {
	case http.MethodGet: // GET
		return responses.API(logic.GetLessonContent(payload))

	case http.MethodPost: // POST
		return responses.API(logic.UpdateLessonContent(payload))

	default: // invalid http method
		return responses.APIMethodNotAllowed()
	}
}
