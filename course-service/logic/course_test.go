package logic_test

import (
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"tazapay.com/elearning/common/driver"
	"tazapay.com/elearning/common/models"
	"tazapay.com/elearning/common/repository/course"
	"tazapay.com/elearning/common/repository/coursesectionlesson"
	"tazapay.com/elearning/common/repository/coursesectionlessoncontent"
	"tazapay.com/elearning/common/request"
	"tazapay.com/elearning/common/responses"
	"tazapay.com/elearning/common/utils"

	"tazapay.com/elearning/svc/course/constants"
	"tazapay.com/elearning/svc/course/logic"
)

func init() {
	driver.ConnectMock()
}

func TestGetAllCourses(t *testing.T) {
	type args struct {
		payload *request.Payload
	}
	tests := []struct {
		name string
		args args
		want *responses.Response
	}{
		{"success", args{payload: &request.Payload{}}, &responses.Response{HTTPStatusCode: 200, Status: "success", MetaData: responses.MetaData{Message: "", Data: utils.MakeResponseData(getTestCourses()), Error: (*responses.Err)(nil), Cookies: []string(nil), Headers: map[string]string(nil)}}},
		{"error", args{payload: &request.Payload{}}, &responses.Response{HTTPStatusCode: 500, Status: "error", MetaData: responses.MetaData{Message: "internal server error", Data: (*responses.Data)(nil), Error: (*responses.Err)(nil), Cookies: []string(nil), Headers: map[string]string(nil)}}},
	}
	for _, tt := range tests {
		mockCourse := new(course.Mock)
		mockCourse.On("SetObj", mock.Anything).Return()

		if tt.name == "success" {
			mockCourse.On("GetAll", mock.Anything).Return(getTestCourses(), nil).Once()
		}

		if tt.name == "error" {
			mockCourse.On("GetAll", mock.Anything).Return(getTestCourses(), errors.New("test")).Once()
		}

		course.SetMock(mockCourse)

		t.Run(tt.name, func(t *testing.T) {
			got := logic.GetAllCourses(tt.args.payload)
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

func TestGetLessonContent(t *testing.T) {
	type args struct {
		payload *request.Payload
	}
	tests := []struct {
		name string
		args args
		want *responses.Response
	}{
		{"invalidLessonId", args{payload: &request.Payload{API: &request.API{PathParameters: make(map[string]string)}}}, &responses.Response{HTTPStatusCode: 400, Status: "error", MetaData: responses.MetaData{Message: "bad request", Data: (*responses.Data)(nil), Error: (*responses.Err)(nil), Cookies: []string(nil), Headers: map[string]string(nil)}}},
		{"success", args{payload: &request.Payload{API: &request.API{PathParameters: map[string]string{constants.PathParamLessonID: "1"}}}}, &responses.Response{HTTPStatusCode: 200, Status: "success", MetaData: responses.MetaData{Message: "", Data: utils.MakeResponseData(getLesson()), Error: (*responses.Err)(nil), Cookies: []string(nil), Headers: map[string]string(nil)}}},
		{"error", args{payload: &request.Payload{API: &request.API{PathParameters: map[string]string{constants.PathParamLessonID: "1"}}}}, &responses.Response{HTTPStatusCode: 500, Status: "error", MetaData: responses.MetaData{Message: "internal server error", Data: (*responses.Data)(nil), Error: (*responses.Err)(nil), Cookies: []string(nil), Headers: map[string]string(nil)}}},
	}
	for _, tt := range tests {
		mocklesson := new(coursesectionlesson.Mock)
		mocklesson.On("SetObj", mock.Anything).Return()

		if tt.name == "success" {
			mocklesson.On("GetByID", mock.Anything, mock.Anything).Return(getLesson(), nil).Once()
		}

		if tt.name == "error" {
			mocklesson.On("GetByID", mock.Anything, mock.Anything).Return(getLesson(), errors.New("test")).Once()
		}

		coursesectionlesson.SetMock(mocklesson)

		t.Run(tt.name, func(t *testing.T) {
			got := logic.GetLessonContent(tt.args.payload)
			assert.Equal(t, tt.want, got)
		})
	}
}

func getLesson() *models.CourseSectionLesson {
	return &models.CourseSectionLesson{
		CourseSectionID: 1,
		ContentID:       1,
		Content: &models.CourseSectionLessonContent{
			Link:              "https://cloud.storage.master.branch",
			Version:           "v1",
			State:             "saved",
			TimeRequiredInSec: 3000,
		},
		Order:         1,
		IsMandatory:   aws.Int(1),
		IsOpenForFree: aws.Int(1),
	}
}

func TestUpdateLessonContent(t *testing.T) {
	type args struct {
		payload *request.Payload
	}
	tests := []struct {
		name string
		args args
		want *responses.Response
	}{
		{"invalidLessonId", args{payload: &request.Payload{API: &request.API{PathParameters: make(map[string]string)}}}, &responses.Response{HTTPStatusCode: 400, Status: "error", MetaData: responses.MetaData{Message: "bad request", Data: (*responses.Data)(nil), Error: (*responses.Err)(nil), Cookies: []string(nil), Headers: map[string]string(nil)}}},
		{"apiReqUnmarshalError", args{payload: &request.Payload{API: &request.API{PathParameters: map[string]string{constants.PathParamLessonID: "1"}, Body: `{"action":"draft","version":what?}`}}}, &responses.Response{HTTPStatusCode: 500, Status: "error", MetaData: responses.MetaData{Message: "internal server error", Data: (*responses.Data)(nil), Error: (*responses.Err)(nil), Cookies: []string(nil), Headers: map[string]string(nil)}}},
		{"getLessonError", args{payload: &request.Payload{API: &request.API{PathParameters: map[string]string{constants.PathParamLessonID: "1"}, Body: `{"action":"draft","version":"v1.1","time_required":2500,"content": {}}`}}}, &responses.Response{HTTPStatusCode: 500, Status: "error", MetaData: responses.MetaData{Message: "internal server error", Data: (*responses.Data)(nil), Error: (*responses.Err)(nil), Cookies: []string(nil), Headers: map[string]string(nil)}}},
		{"invalidAction", args{payload: &request.Payload{API: &request.API{PathParameters: map[string]string{constants.PathParamLessonID: "1"}, Body: `{"action":"abc"}`}}}, &responses.Response{HTTPStatusCode: 400, Status: "error", MetaData: responses.MetaData{Message: "bad request", Data: (*responses.Data)(nil), Error: (*responses.Err)(nil), Cookies: []string(nil), Headers: map[string]string(nil)}}},
		{"newDraft", args{payload: &request.Payload{API: &request.API{PathParameters: map[string]string{constants.PathParamLessonID: "1"}, Body: `{"action":"draft","version":"v1.1","time_required":2500,"content": {}}`}}}, &responses.Response{HTTPStatusCode: 200, Status: "success", MetaData: responses.MetaData{Message: "", Data: (*responses.Data)(nil), Error: (*responses.Err)(nil), Cookies: []string(nil), Headers: map[string]string(nil)}}},
		{"actionDraftGetDraftedContentError", args{payload: &request.Payload{API: &request.API{PathParameters: map[string]string{constants.PathParamLessonID: "1"}, Body: `{"action":"draft","version":"v1.1","time_required":2500,"content": {}}`}}}, &responses.Response{HTTPStatusCode: 500, Status: "error", MetaData: responses.MetaData{Message: "internal server error", Data: (*responses.Data)(nil), Error: (*responses.Err)(nil), Cookies: []string(nil), Headers: map[string]string(nil)}}},
		{"newDraftCreateError", args{payload: &request.Payload{API: &request.API{PathParameters: map[string]string{constants.PathParamLessonID: "1"}, Body: `{"action":"draft","version":"v1.1","time_required":2500,"content": {}}`}}}, &responses.Response{HTTPStatusCode: 500, Status: "error", MetaData: responses.MetaData{Message: "internal server error", Data: (*responses.Data)(nil), Error: (*responses.Err)(nil), Cookies: []string(nil), Headers: map[string]string(nil)}}},
		{"merge", args{payload: &request.Payload{API: &request.API{PathParameters: map[string]string{constants.PathParamLessonID: "1"}, Body: `{"action":"merge","version":"v1.1"}`}}}, &responses.Response{HTTPStatusCode: 200, Status: "success", MetaData: responses.MetaData{Message: "", Data: (*responses.Data)(nil), Error: (*responses.Err)(nil), Cookies: []string(nil), Headers: map[string]string(nil)}}},
		{"mergeUpdateError", args{payload: &request.Payload{API: &request.API{PathParameters: map[string]string{constants.PathParamLessonID: "1"}, Body: `{"action":"merge","version":"v1.1"}`}}}, &responses.Response{HTTPStatusCode: 500, Status: "error", MetaData: responses.MetaData{Message: "internal server error", Data: (*responses.Data)(nil), Error: (*responses.Err)(nil), Cookies: []string(nil), Headers: map[string]string(nil)}}},
		{"actionMergeGetDraftedContentError", args{payload: &request.Payload{API: &request.API{PathParameters: map[string]string{constants.PathParamLessonID: "1"}, Body: `{"action":"merge","version":"v1.1"}`}}}, &responses.Response{HTTPStatusCode: 500, Status: "error", MetaData: responses.MetaData{Message: "internal server error", Data: (*responses.Data)(nil), Error: (*responses.Err)(nil), Cookies: []string(nil), Headers: map[string]string(nil)}}},
		{"mergeConflict", args{payload: &request.Payload{API: &request.API{PathParameters: map[string]string{constants.PathParamLessonID: "1"}, Body: `{"action":"merge","version":"v1.1"}`}}}, &responses.Response{HTTPStatusCode: 409, Status: "error", MetaData: responses.MetaData{Message: "conflict with the current state of the resource", Data: (*responses.Data)(nil), Error: (*responses.Err)(nil), Cookies: []string(nil), Headers: map[string]string(nil)}}},
		{"save", args{payload: &request.Payload{API: &request.API{PathParameters: map[string]string{constants.PathParamLessonID: "1"}, Body: `{"action":"save","version":"v1.1"}`}}}, &responses.Response{HTTPStatusCode: 200, Status: "success", MetaData: responses.MetaData{Message: "", Data: (*responses.Data)(nil), Error: (*responses.Err)(nil), Cookies: []string(nil), Headers: map[string]string(nil)}}},
		{"actionSaveGetDraftedContentError", args{payload: &request.Payload{API: &request.API{PathParameters: map[string]string{constants.PathParamLessonID: "1"}, Body: `{"action":"save","version":"v1.1"}`}}}, &responses.Response{HTTPStatusCode: 500, Status: "error", MetaData: responses.MetaData{Message: "internal server error", Data: (*responses.Data)(nil), Error: (*responses.Err)(nil), Cookies: []string(nil), Headers: map[string]string(nil)}}},
		{"saveConflict", args{payload: &request.Payload{API: &request.API{PathParameters: map[string]string{constants.PathParamLessonID: "1"}, Body: `{"action":"save","version":"v1.1"}`}}}, &responses.Response{HTTPStatusCode: 409, Status: "error", MetaData: responses.MetaData{Message: "conflict with the current state of the resource", Data: (*responses.Data)(nil), Error: (*responses.Err)(nil), Cookies: []string(nil), Headers: map[string]string(nil)}}},
		{"saveUpdateError", args{payload: &request.Payload{API: &request.API{PathParameters: map[string]string{constants.PathParamLessonID: "1"}, Body: `{"action":"save","version":"v1.1"}`}}}, &responses.Response{HTTPStatusCode: 500, Status: "error", MetaData: responses.MetaData{Message: "internal server error", Data: (*responses.Data)(nil), Error: (*responses.Err)(nil), Cookies: []string(nil), Headers: map[string]string(nil)}}},
		{"saveUpdateError2", args{payload: &request.Payload{API: &request.API{PathParameters: map[string]string{constants.PathParamLessonID: "1"}, Body: `{"action":"save","version":"v1.1"}`}}}, &responses.Response{HTTPStatusCode: 500, Status: "error", MetaData: responses.MetaData{Message: "internal server error", Data: (*responses.Data)(nil), Error: (*responses.Err)(nil), Cookies: []string(nil), Headers: map[string]string(nil)}}},
	}
	for _, tt := range tests {
		mocklesson := new(coursesectionlesson.Mock)
		mocklesson.On("SetObj", mock.Anything).Return()

		if tt.name == "getLessonError" {
			mocklesson.On("GetByID", mock.Anything, mock.Anything).Return(getLesson(), errors.New("test")).Once()
			coursesectionlesson.SetMock(mocklesson)
		} else {
			mocklesson.On("GetByID", mock.Anything, mock.Anything).Return(getLesson(), nil).Once()
			coursesectionlesson.SetMock(mocklesson)
		}

		mockContent := new(coursesectionlessoncontent.Mock)
		mockContent.On("SetObj", mock.Anything).Return()

		if tt.name == "newDraft" {
			mockContent.On("GetByParentIDAndVersion", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil).Once()
			mockContent.On("Create", mock.Anything, mock.Anything).Return(nil).Once()
		}

		if tt.name == "actionDraftGetDraftedContentError" || tt.name == "actionMergeGetDraftedContentError" || tt.name == "actionSaveGetDraftedContentError" {
			mockContent.On("GetByParentIDAndVersion", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("test")).Once()
		}

		if tt.name == "newDraftCreateError" {
			mockContent.On("GetByParentIDAndVersion", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil).Once()
			mockContent.On("Create", mock.Anything, mock.Anything).Return(errors.New("test")).Once()
		}

		if tt.name == "merge" {
			mockContent.On("GetByParentIDAndVersion", mock.Anything, mock.Anything, mock.Anything).Return(getDraftedContentToMerge(), nil).Once()
			mockContent.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
		}

		if tt.name == "mergeUpdateError" {
			mockContent.On("GetByParentIDAndVersion", mock.Anything, mock.Anything, mock.Anything).Return(getDraftedContentToMerge(), nil).Once()
			mockContent.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("test")).Once()
		}

		if tt.name == "mergeConflict" || tt.name == "saveConflict" {
			mockContent.On("GetByParentIDAndVersion", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil).Once()
		}

		if tt.name == "save" {
			mockContent.On("GetByParentIDAndVersion", mock.Anything, mock.Anything, mock.Anything).Return(getDraftedContentToSave(), nil).Once()
			mockContent.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(nil).Twice()
		}

		if tt.name == "saveUpdateError" {
			mockContent.On("GetByParentIDAndVersion", mock.Anything, mock.Anything, mock.Anything).Return(getDraftedContentToSave(), nil).Once()
			mockContent.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("test")).Once()
		}

		if tt.name == "saveUpdateError2" {
			mockContent.On("GetByParentIDAndVersion", mock.Anything, mock.Anything, mock.Anything).Return(getDraftedContentToSave(), nil).Once()
			mockContent.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
			mockContent.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("test")).Once()
		}

		coursesectionlessoncontent.SetMock(mockContent)

		t.Run(tt.name, func(t *testing.T) {
			got := logic.UpdateLessonContent(tt.args.payload)
			assert.Equal(t, tt.want, got)
		})
	}
}

func getDraftedContentToMerge() *models.CourseSectionLessonContent {
	return &models.CourseSectionLessonContent{
		Link:              "https://cloud.storage.update.branch.v1.1",
		Version:           "v1.1",
		State:             "draft",
		TimeRequiredInSec: 2500,
	}
}

func getDraftedContentToSave() *models.CourseSectionLessonContent {
	return &models.CourseSectionLessonContent{
		Link:              "https://cloud.storage.update.branch.v1.1",
		Version:           "v1.1",
		State:             "merged",
		TimeRequiredInSec: 2500,
	}
}
