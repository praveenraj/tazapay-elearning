package handlers_test

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"

	"github.com/stretchr/testify/assert"

	"tazapay.com/elearning/common/handlers"
	"tazapay.com/elearning/common/request"
	"tazapay.com/elearning/common/responses"
)

func TestAPI(t *testing.T) {
	type args struct {
		event      interface{}
		controller func(*request.Payload) *responses.APIResponse
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr error
	}{
		{"internalError", args{`{"key":what?}`, testController}, &responses.APIResponse{StatusCode: 500, Headers: map[string]string{"content-type": "application/json"}, Body: "{\"status\":\"error\",\"message\":\"internal server error\"}", Cookies: []string(nil), IsBase64Encoded: false}, nil},
		{"success", args{events.APIGatewayV2HTTPRequest{RouteKey: "GET /api", Headers: map[string]string{"key": "val"}, Cookies: []string{"cookieKey=cookieVal"},
			QueryStringParameters: map[string]string{"key": "val"}, PathParameters: map[string]string{"key": "val"}, StageVariables: map[string]string{"key": "val"}, Body: "test"},
			testController}, &responses.APIResponse{StatusCode: 200, Headers: map[string]string{"content-type": "application/json"}, Body: "{\"status\":\"success\"}", Cookies: []string(nil), IsBase64Encoded: false}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := handlers.API(tt.args.event, tt.args.controller)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func testController(payload *request.Payload) *responses.APIResponse {
	return responses.API(responses.Success())
}
