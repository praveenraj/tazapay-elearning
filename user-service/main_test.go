package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"tazapay.com/elearning/common/responses"
)

func Test_handler(t *testing.T) {
	type args struct {
		ctx   context.Context
		event interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr error
	}{
		{"invalidPath", args{ctx: context.Background(), event: map[string]interface{}{"routeKey": "GET /test"}}, &responses.APIResponse{StatusCode: 400, Headers: map[string]string{"content-type": "application/json"}, Body: "{\"status\":\"error\",\"message\":\"bad request\",\"error\":null}", Cookies: []string(nil), IsBase64Encoded: false}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := handler(tt.args.ctx, tt.args.event)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
