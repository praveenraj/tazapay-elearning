package request

import (
	"strings"

	"github.com/aws/aws-lambda-go/events"

	"tazapay.com/elearning/common/constants"
)

var _payload *Payload

// GetPayload return global payload object
func GetPayload() *Payload {
	return _payload
}

// SetPayload assign and make payload to be available at global level
func SetPayload(payload *Payload) {
	_payload = payload
}

// Payload common request structure for the application.
type Payload struct {
	Source string // source of the request
	API    *API   // api-gateway
}

// API custom api request
type API struct {
	Version               string
	RouteKey              string
	HTTPMethod            string
	Resource              string
	Cookies               map[string]string
	Headers               map[string]string
	QueryStringParameters map[string]string
	PathParameters        map[string]string
	StageVariables        map[string]string
	RequestContext        *apiRequestContext
	Body                  string
	IsBase64Encoded       bool
}

// apiRequestContext custom api request context
type apiRequestContext struct {
	APIID        string
	AccountID    string
	Stage        string
	RequestID    string
	Protocol     string
	SourceIP     string
	UserAgent    string
	DomainName   string
	DomainPrefix string
	Authorizer   *apiRequestAuthorizer
	TimeEpoch    int64
}

// apiRequestAuthorizer custom api request authorizer
type apiRequestAuthorizer struct {
	JWT *jwt
}

// jwt JSON web-token values for the api authorizer
type jwt struct {
	Claims map[string]string
	Scopes []string
}

// GetSource get request source
func (p *Payload) GetSource() string {
	return p.Source
}

// SetSource set request source name
func (p *Payload) SetSource(source string) {
	p.Source = source
}

// GetAPI get api request
func (p *Payload) GetAPI() *API {
	return p.API
}

// SetAPI set api request
func (p *Payload) SetAPI(api *API) {
	p.API = api
}

// CreatePayload convert every request into common payload for the application
func CreatePayload(req interface{}) (*Payload, error) {
	var err error
	payload := &Payload{}

	// switch to the respective event type
	switch v := req.(type) {
	case events.APIGatewayV2HTTPRequest: // api-gateway v2 request
		payload.SetSource(constants.HandlerAPIGW)
		payload.SetAPI(castAPIGatewayV2Request(&v))

	default:
		return nil, constants.ErrInvalidLambdaEvent
	}

	// set payload to global object
	SetPayload(payload)
	return payload, err
}

// castAPIGatewayV2Request parse api-gateway v2 request to payload's api
func castAPIGatewayV2Request(v2 *events.APIGatewayV2HTTPRequest) *API {
	/* check and assign default values */
	cookies := make(map[string]string)
	// parse cookies and save as a key-value pair
	for _, pair := range v2.Cookies {
		values := strings.Split(pair, constants.Equal)
		if len(values) == constants.IntTwo {
			cookies[values[constants.IntZero]] = values[constants.IntOne]
		}
	}
	headers := make(map[string]string)
	// convert header names in lowercase
	for k, v := range v2.Headers {
		headers[strings.ToLower(k)] = v
	}
	queryParameters := make(map[string]string)
	if len(v2.QueryStringParameters) > constants.IntZero {
		queryParameters = v2.QueryStringParameters
	}
	pathParameters := make(map[string]string)
	if len(v2.PathParameters) > constants.IntZero {
		pathParameters = v2.PathParameters
	}
	stageVariables := make(map[string]string)
	if len(v2.StageVariables) > constants.IntZero {
		stageVariables = v2.StageVariables
	}

	// construct and return custom API request
	return &API{
		Version:               v2.Version,
		RouteKey:              v2.RouteKey,
		Resource:              strings.Split(v2.RouteKey, constants.Space)[constants.IntOne],
		HTTPMethod:            v2.RequestContext.HTTP.Method,
		Cookies:               cookies,
		Headers:               headers,
		QueryStringParameters: queryParameters,
		PathParameters:        pathParameters,
		StageVariables:        stageVariables,
		Body:                  v2.Body,
		IsBase64Encoded:       v2.IsBase64Encoded,
		RequestContext: &apiRequestContext{
			APIID:        v2.RequestContext.APIID,
			AccountID:    v2.RequestContext.AccountID,
			Stage:        v2.RequestContext.Stage,
			RequestID:    v2.RequestContext.RequestID,
			Protocol:     v2.RequestContext.HTTP.Protocol,
			SourceIP:     v2.RequestContext.HTTP.SourceIP,
			UserAgent:    v2.RequestContext.HTTP.UserAgent,
			DomainName:   v2.RequestContext.DomainName,
			DomainPrefix: v2.RequestContext.DomainPrefix,
			Authorizer: &apiRequestAuthorizer{
				JWT: getJWTFromAPIv2Authorizer(&v2.RequestContext),
			},
			TimeEpoch: v2.RequestContext.TimeEpoch,
		},
	}
}

// getJWTFromAPIv2Authorizer parse jwt values from api-gateway v2 request
func getJWTFromAPIv2Authorizer(requestContext *events.APIGatewayV2HTTPRequestContext) *jwt {
	// init with default values
	JWT := jwt{
		Claims: make(map[string]string),
		Scopes: make([]string, constants.IntZero),
	}

	// check whether the authorizer object have values
	if requestContext.Authorizer != nil {
		// jwt key claims
		if len(requestContext.Authorizer.JWT.Claims) > constants.IntZero {
			JWT.Claims = requestContext.Authorizer.JWT.Claims
		}

		// jwt key scopes
		if len(requestContext.Authorizer.JWT.Scopes) > constants.IntZero {
			JWT.Scopes = requestContext.Authorizer.JWT.Scopes
		}
	}
	return &JWT
}
