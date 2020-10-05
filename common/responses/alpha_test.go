package responses_test

import (
	"testing"

	"tazapay.com/elearning/common/utils"

	"github.com/stretchr/testify/assert"
	"tazapay.com/elearning/common/responses"
)

func TestSuccess(t *testing.T) {
	data := utils.MakeResponseData(map[string]interface{}{"company": "tryllo"})
	type args struct {
		args []interface{}
	}
	tests := []struct {
		name string
		args args
		want *responses.Response
	}{
		{"default", args{[]interface{}{}}, &responses.Response{200, "success", responses.MetaData{"", (*responses.Data)(nil), (*responses.Err)(nil), []string(nil), map[string]string(nil)}}},
		{"msg", args{[]interface{}{"test"}}, &responses.Response{200, "success", responses.MetaData{"test", (*responses.Data)(nil), (*responses.Err)(nil), []string(nil), map[string]string(nil)}}},
		{"data", args{[]interface{}{*data}}, &responses.Response{200, "success", responses.MetaData{"", data, (*responses.Err)(nil), []string(nil), map[string]string(nil)}}},
		{"*data", args{[]interface{}{data}}, &responses.Response{200, "success", responses.MetaData{"", data, (*responses.Err)(nil), []string(nil), map[string]string(nil)}}},
		{"intfData", args{[]interface{}{map[string]interface{}{"company": "tryllo"}}}, &responses.Response{200, "success", responses.MetaData{"", data, (*responses.Err)(nil), []string(nil), map[string]string(nil)}}},
		{"msgAndData", args{[]interface{}{*data, "test"}}, &responses.Response{200, "success", responses.MetaData{"test", data, (*responses.Err)(nil), []string(nil), map[string]string(nil)}}},
		{"dataAndcookies", args{[]interface{}{*data, []string{"cookieName=cookieValue"}}}, &responses.Response{200, "success", responses.MetaData{"", data, (*responses.Err)(nil), []string{"cookieName=cookieValue"}, map[string]string(nil)}}},
		{"msgAndHeaders", args{[]interface{}{"test", map[string]string{"key": "value"}}}, &responses.Response{200, "success", responses.MetaData{"test", (*responses.Data)(nil), (*responses.Err)(nil), []string(nil), map[string]string{"key": "value"}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := responses.Success(tt.args.args...)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestError(t *testing.T) {
	err := utils.MakeResponseError("errors.auth", 1000)
	data := utils.MakeResponseData(map[string]interface{}{"company": "tryllo"})
	type args struct {
		statusCode int
		msg        string
		args       []interface{}
	}
	tests := []struct {
		name string
		args args
		want *responses.Response
	}{
		{"default", args{500, "internal server error", []interface{}{}}, &responses.Response{500, "error", responses.MetaData{"internal server error", (*responses.Data)(nil), (*responses.Err)(nil), []string(nil), map[string]string(nil)}}},
		{"Err", args{452, "custom error", []interface{}{*err}}, &responses.Response{452, "error", responses.MetaData{"custom error", (*responses.Data)(nil), err, []string(nil), map[string]string(nil)}}},
		{"*Err", args{452, "custom error", []interface{}{err}}, &responses.Response{452, "error", responses.MetaData{"custom error", (*responses.Data)(nil), err, []string(nil), map[string]string(nil)}}},
		{"dataErr", args{452, "custom error", []interface{}{*data, err}}, &responses.Response{452, "error", responses.MetaData{"custom error", data, err, []string(nil), map[string]string(nil)}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := responses.Error(tt.args.statusCode, tt.args.msg, tt.args.args...)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInternalError(t *testing.T) {
	type args struct {
		args []interface{}
	}
	tests := []struct {
		name string
		args args
		want *responses.Response
	}{
		{"default", args{[]interface{}{}}, &responses.Response{500, "error", responses.MetaData{"internal server error", (*responses.Data)(nil), (*responses.Err)(nil), []string(nil), map[string]string(nil)}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := responses.InternalError(tt.args.args...)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBadRequest(t *testing.T) {
	type args struct {
		args []interface{}
	}
	tests := []struct {
		name string
		args args
		want *responses.Response
	}{
		{"default", args{[]interface{}{}}, &responses.Response{400, "error", responses.MetaData{"bad request", (*responses.Data)(nil), (*responses.Err)(nil), []string(nil), map[string]string(nil)}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := responses.BadRequest(tt.args.args...)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestNotFound(t *testing.T) {
	type args struct {
		args []interface{}
	}
	tests := []struct {
		name string
		args args
		want *responses.Response
	}{
		{"default", args{[]interface{}{}}, &responses.Response{404, "error", responses.MetaData{"resource not found", (*responses.Data)(nil), (*responses.Err)(nil), []string(nil), map[string]string(nil)}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := responses.NotFound(tt.args.args...)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestConflict(t *testing.T) {
	type args struct {
		args []interface{}
	}
	tests := []struct {
		name string
		args args
		want *responses.Response
	}{
		{"default", args{[]interface{}{}}, &responses.Response{409, "error", responses.MetaData{"conflict with the current state of the resource", (*responses.Data)(nil), (*responses.Err)(nil), []string(nil), map[string]string(nil)}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := responses.Conflict(tt.args.args...)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUnAuthorized(t *testing.T) {
	type args struct {
		args []interface{}
	}
	tests := []struct {
		name string
		args args
		want *responses.Response
	}{
		{"default", args{[]interface{}{}}, &responses.Response{401, "error", responses.MetaData{"unauthorized", (*responses.Data)(nil), (*responses.Err)(nil), []string(nil), map[string]string(nil)}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := responses.UnAuthorized(tt.args.args...)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestForbidden(t *testing.T) {
	type args struct {
		args []interface{}
	}
	tests := []struct {
		name string
		args args
		want *responses.Response
	}{
		{"default", args{[]interface{}{}}, &responses.Response{403, "error", responses.MetaData{"forbidden to access the resource", (*responses.Data)(nil), (*responses.Err)(nil), []string(nil), map[string]string(nil)}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := responses.Forbidden(tt.args.args...)
			assert.Equal(t, tt.want, got)
		})
	}
}
