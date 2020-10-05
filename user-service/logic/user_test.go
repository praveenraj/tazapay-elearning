package logic_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"tazapay.com/elearning/common/driver"
	"tazapay.com/elearning/common/repository/user"
	"tazapay.com/elearning/common/request"
	"tazapay.com/elearning/common/responses"
	"tazapay.com/elearning/common/utils"

	"tazapay.com/elearning/svc/user/logic"
)

func init() {
	driver.ConnectMock()
}

func TestRegisterUser(t *testing.T) {
	type args struct {
		payload *request.Payload
	}
	tests := []struct {
		name string
		args args
		want *responses.Response
	}{
		{"apiReqUnmarshalError", args{payload: &request.Payload{API: &request.API{Body: `{"first_name":what?}`}}}, &responses.Response{HTTPStatusCode: 500, Status: "error", MetaData: responses.MetaData{Message: "internal server error", Data: (*responses.Data)(nil), Error: (*responses.Err)(nil), Cookies: []string(nil), Headers: map[string]string(nil)}}},
		{"invalidApiReq", args{payload: &request.Payload{API: &request.API{Body: `{"first_name":""}`}}}, &responses.Response{HTTPStatusCode: 400, Status: "error", MetaData: responses.MetaData{Message: "bad request", Data: (*responses.Data)(nil), Error: utils.MakeResponseError("general", 1100, "first_name"), Cookies: []string(nil), Headers: map[string]string(nil)}}},
		{"register", args{payload: &request.Payload{API: &request.API{Body: `{"first_name":"test","mobile":"2548758684","email":"abc@gmail.com"}`}}}, &responses.Response{HTTPStatusCode: 200, Status: "success", MetaData: responses.MetaData{Message: "", Data: (*responses.Data)(nil), Error: (*responses.Err)(nil), Cookies: []string(nil), Headers: map[string]string(nil)}}},
		{"registerError", args{payload: &request.Payload{API: &request.API{Body: `{"first_name":"test","mobile":"2548758684","email":"abc@gmail.com"}`}}}, &responses.Response{HTTPStatusCode: 500, Status: "error", MetaData: responses.MetaData{Message: "internal server error", Data: (*responses.Data)(nil), Error: (*responses.Err)(nil), Cookies: []string(nil), Headers: map[string]string(nil)}}},
		{"registerErrorUnique", args{payload: &request.Payload{API: &request.API{Body: `{"first_name":"test","mobile":"2548758684","email":"abc@gmail.com"}`}}}, &responses.Response{HTTPStatusCode: 409, Status: "error", MetaData: responses.MetaData{Message: "conflict with the current state of the resource", Data: (*responses.Data)(nil), Error: (*responses.Err)(nil), Cookies: []string(nil), Headers: map[string]string(nil)}}},
	}
	for _, tt := range tests {
		mockUser := new(user.Mock)
		mockUser.On("SetObj", mock.Anything).Return()

		if tt.name == "register" {
			mockUser.On("Create", mock.Anything, mock.Anything).Return(nil).Once()
		}

		if tt.name == "registerError" {
			mockUser.On("Create", mock.Anything, mock.Anything).Return(errors.New("test")).Once()
		}

		if tt.name == "registerErrorUnique" {
			mockUser.On("Create", mock.Anything, mock.Anything).Return(errors.New("Error 1062")).Once()
		}

		user.SetMock(mockUser)

		t.Run(tt.name, func(t *testing.T) {
			got := logic.RegisterUser(tt.args.payload)
			assert.Equal(t, tt.want, got)
		})
	}
}
