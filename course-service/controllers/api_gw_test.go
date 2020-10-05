package controllers_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"tazapay.com/elearning/common/driver"
	"tazapay.com/elearning/common/models"
	"tazapay.com/elearning/common/repository/course"
	"tazapay.com/elearning/common/request"
	"tazapay.com/elearning/common/responses"

	"tazapay.com/elearning/svc/course/controllers"
)

func init() {
	driver.ConnectMock()
}

func TestAPI(t *testing.T) {
	type args struct {
		payload *request.Payload
	}
	tests := []struct {
		name string
		args args
		want *responses.APIResponse
	}{
		{"invalidPath", args{payload: &request.Payload{API: &request.API{Resource: "/abc"}}}, &responses.APIResponse{StatusCode: 400, Headers: map[string]string{"content-type": "application/json"}, Body: "{\"status\":\"error\",\"message\":\"bad request\",\"error\":null}", Cookies: []string(nil), IsBase64Encoded: false}},
		{"coursesInvalidMethod", args{payload: &request.Payload{API: &request.API{Resource: "/courses", HTTPMethod: "HEAD"}}}, &responses.APIResponse{StatusCode: 405, Headers: map[string]string{"content-type": "application/json"}, Body: "{\"status\":\"error\",\"message\":\"method not allowed\"}", Cookies: []string(nil), IsBase64Encoded: false}},
		{"lessonContentInvalidMethod", args{payload: &request.Payload{API: &request.API{Resource: "/courses/{course_id}/sections/{section_id}/lessons/{lesson_id}", HTTPMethod: "HEAD"}}}, &responses.APIResponse{StatusCode: 405, Headers: map[string]string{"content-type": "application/json"}, Body: "{\"status\":\"error\",\"message\":\"method not allowed\"}", Cookies: []string(nil), IsBase64Encoded: false}},
		{"getAllCourses", args{payload: &request.Payload{API: &request.API{Resource: "/courses", HTTPMethod: "GET"}}}, &responses.APIResponse{StatusCode: 200, Headers: map[string]string{"content-type": "application/json"}, Body: "{\"status\":\"success\",\"data\":[{\"title\":\"Unit Test in Go\",\"introduction_brief\":\"intro brief\",\"fee\":149,\"language\":\"English\"}]}", Cookies: []string(nil), IsBase64Encoded: false}},
		{"getLessonContent", args{payload: &request.Payload{API: &request.API{Resource: "/courses/{course_id}/sections/{section_id}/lessons/{lesson_id}", HTTPMethod: "GET", PathParameters: make(map[string]string)}}}, &responses.APIResponse{StatusCode: 400, Headers: map[string]string{"content-type": "application/json"}, Body: "{\"status\":\"error\",\"message\":\"bad request\"}", Cookies: []string(nil), IsBase64Encoded: false}},
		{"updateLessonContent", args{payload: &request.Payload{API: &request.API{Resource: "/courses/{course_id}/sections/{section_id}/lessons/{lesson_id}", HTTPMethod: "POST", PathParameters: make(map[string]string)}}}, &responses.APIResponse{StatusCode: 400, Headers: map[string]string{"content-type": "application/json"}, Body: "{\"status\":\"error\",\"message\":\"bad request\"}", Cookies: []string(nil), IsBase64Encoded: false}},
	}
	for _, tt := range tests {
		if tt.name == "getAllCourses" {
			mockCourse := new(course.Mock)
			mockCourse.On("SetObj", mock.Anything).Return()
			mockCourse.On("GetAll", mock.Anything).Return(getTestCourses(), nil).Once()
			course.SetMock(mockCourse)
		}

		t.Run(tt.name, func(t *testing.T) {
			got := controllers.API(tt.args.payload)
			assert.Equal(t, tt.want, got)
		})
	}
}

func getTestCourses() []*models.Course {
	return []*models.Course{
		&models.Course{
			Title:             "Unit Test in Go",
			IntroductionBrief: "intro brief",
			Fee:               149,
			Language:          "English",
		},
	}
}
