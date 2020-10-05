package responses_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"tazapay.com/elearning/common/constants"
	"tazapay.com/elearning/common/responses"
	"tazapay.com/elearning/common/utils"
)

func TestAPI(t *testing.T) {
	type args struct {
		res *responses.Response
	}
	tests := []struct {
		name string
		args args
		want *responses.APIResponse
	}{
		{"plain", args{responses.Success()}, &responses.APIResponse{StatusCode: 200, Headers: map[string]string{constants.HeaderKeyContentType: constants.MediaTypeApplicationJSON}, Body: "{\"status\":\"success\"}", IsBase64Encoded: false}},
		{"data", args{responses.Success(map[string]interface{}{"company": "tryllo"})}, &responses.APIResponse{StatusCode: 200, Headers: map[string]string{constants.HeaderKeyContentType: constants.MediaTypeApplicationJSON}, Body: "{\"status\":\"success\",\"data\":{\"company\":\"tryllo\"}}", IsBase64Encoded: false}},
		{"error", args{responses.InternalError(utils.MakeResponseError("auth", 1100))}, &responses.APIResponse{StatusCode: 500, Headers: map[string]string{constants.HeaderKeyContentType: constants.MediaTypeApplicationJSON}, Body: "{\"status\":\"error\",\"message\":\"internal server error\",\"error\":{\"code\":1100,\"type\":\"auth\"}}", IsBase64Encoded: false}},
		{"marshalError", args{responses.Success(utils.MakeResponseData(make(chan int)))}, &responses.APIResponse{StatusCode: 500, Headers: map[string]string{constants.HeaderKeyContentType: constants.MediaTypeApplicationJSON}, Body: "{\"status\":\"error\",\"message\":\"internal server error\"}", IsBase64Encoded: false}},
		{"dataCookies", args{responses.Success(map[string]interface{}{"company": "tryllo"}, []string{"cookieName=CookieValue"})}, &responses.APIResponse{StatusCode: 200, Headers: map[string]string{constants.HeaderKeyContentType: constants.MediaTypeApplicationJSON}, Body: "{\"status\":\"success\",\"data\":{\"company\":\"tryllo\"}}", Cookies: []string{"cookieName=CookieValue"}, IsBase64Encoded: false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := responses.API(tt.args.res)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestAPIInternalError(t *testing.T) {
	tests := []struct {
		name string
		want *responses.APIResponse
	}{
		{"default", &responses.APIResponse{StatusCode: 500, Headers: map[string]string{constants.HeaderKeyContentType: constants.MediaTypeApplicationJSON}, Body: "{\"status\":\"error\",\"message\":\"internal server error\"}", IsBase64Encoded: false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := responses.APIInternalError()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestAPIServiceUnavailable(t *testing.T) {
	tests := []struct {
		name string
		want *responses.APIResponse
	}{
		{"default", &responses.APIResponse{StatusCode: 503, Headers: map[string]string{constants.HeaderKeyContentType: constants.MediaTypeApplicationJSON}, Body: "{\"status\":\"error\",\"message\":\"service unavailable\"}", IsBase64Encoded: false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := responses.APIServiceUnavailable()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestAPIMethodNotAllowed(t *testing.T) {
	tests := []struct {
		name string
		want *responses.APIResponse
	}{
		{"default", &responses.APIResponse{StatusCode: 405, Headers: map[string]string{constants.HeaderKeyContentType: constants.MediaTypeApplicationJSON}, Body: "{\"status\":\"error\",\"message\":\"method not allowed\"}", IsBase64Encoded: false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := responses.APIMethodNotAllowed()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestAPIBadRequest(t *testing.T) {
	tests := []struct {
		name string
		want *responses.APIResponse
	}{
		{"default", &responses.APIResponse{StatusCode: 400, Headers: map[string]string{constants.HeaderKeyContentType: constants.MediaTypeApplicationJSON}, Body: "{\"status\":\"error\",\"message\":\"bad request\",\"error\":null}", IsBase64Encoded: false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := responses.APIBadRequest(nil)
			assert.Equal(t, tt.want, got)
		})
	}
}
