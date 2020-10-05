package controllers_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"tazapay.com/elearning/common/request"
	"tazapay.com/elearning/common/responses"

	"tazapay.com/elearning/svc/user/controllers"
)

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
		{"userInvalidMethod", args{payload: &request.Payload{API: &request.API{Resource: "/user", HTTPMethod: "HEAD"}}}, &responses.APIResponse{StatusCode: 405, Headers: map[string]string{"content-type": "application/json"}, Body: "{\"status\":\"error\",\"message\":\"method not allowed\"}", Cookies: []string(nil), IsBase64Encoded: false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := controllers.API(tt.args.payload)
			assert.Equal(t, tt.want, got)
		})
	}
}
