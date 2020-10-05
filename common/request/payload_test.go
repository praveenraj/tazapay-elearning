package request_test

import (
	"testing"

	"tazapay.com/elearning/common/constants"
	"tazapay.com/elearning/common/request"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestPayload_SetFunc(t *testing.T) {
	type args struct {
		Source string
		API    *events.APIGatewayV2HTTPRequest
	}
	tests := []struct {
		name string
		args args
	}{
		{"default", args{Source: "test", API: &events.APIGatewayV2HTTPRequest{RouteKey: "GET /api"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &request.Payload{}
			p.SetSource(tt.args.Source)
			assert.Equal(t, "test", p.GetSource())

			p.SetAPI(request.CastAPIGatewayV2Request(tt.args.API))
			assert.NotNil(t, p.GetAPI())
		})
	}
}

func TestCreatePayload(t *testing.T) {
	apiV2Event := events.APIGatewayV2HTTPRequest{RouteKey: "GET /api"}
	apiV2EventWithoutDefaults := events.APIGatewayV2HTTPRequest{RouteKey: "GET /api", Headers: map[string]string{"key": "val"}, Cookies: []string{"cookieKey=cookieVal"},
		QueryStringParameters: map[string]string{"key": "val"}, PathParameters: map[string]string{"key": "val"}, StageVariables: map[string]string{"key": "val"}}

	reqCtx := &events.APIGatewayV2HTTPRequestContext{
		Authorizer: &events.APIGatewayV2HTTPRequestContextAuthorizerDescription{
			JWT: events.APIGatewayV2HTTPRequestContextAuthorizerJWTDescription{
				Claims: map[string]string{"sub": "1234"},
				Scopes: []string{"read"},
			},
		},
	}
	apiV2EventJWT := events.APIGatewayV2HTTPRequest{RouteKey: "GET /api", RequestContext: *reqCtx}

	type args struct {
		req interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    *request.Payload
		wantErr error
	}{
		{"apiV2", args{apiV2Event}, &request.Payload{Source: "API", API: request.CastAPIGatewayV2Request(&apiV2Event)}, nil},
		{"apiV2WithoutDefaults", args{apiV2EventWithoutDefaults}, &request.Payload{Source: "API", API: request.CastAPIGatewayV2Request(&apiV2EventWithoutDefaults)}, nil},
		{"apiV2JWT", args{apiV2EventJWT}, &request.Payload{Source: "API", API: request.CastAPIGatewayV2Request(&apiV2EventJWT)}, nil},
		{"invalid", args{}, nil, constants.ErrInvalidLambdaEvent},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := request.CreatePayload(tt.args.req)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestGetPayload(t *testing.T) {
	tests := []struct {
		name string
		want *request.Payload
	}{
		{"nilPayload", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request.SetPayload(nil)
			assert.Equal(t, tt.want, request.GetPayload())
		})
	}
}
