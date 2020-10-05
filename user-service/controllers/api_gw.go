package controllers

import (
	"net/http"

	"tazapay.com/elearning/svc/user/constants"
	"tazapay.com/elearning/svc/user/logic"

	"tazapay.com/elearning/common/request"
	"tazapay.com/elearning/common/responses"
)

// API API gateway controller for the service
func API(payload *request.Payload) *responses.APIResponse {
	// resource path routing
	switch payload.GetAPI().Resource {
	case constants.PathUser: // /user
		return handleUser(payload)

	default: // invalid api path
		return responses.APIBadRequest(nil)
	}
}

// handleUser handles /user request
func handleUser(payload *request.Payload) *responses.APIResponse {
	// method routing
	switch payload.GetAPI().HTTPMethod {
	case http.MethodPut: // PUT
		return responses.API(logic.RegisterUser(payload))

	default: // invalid http method
		return responses.APIMethodNotAllowed()
	}
}
