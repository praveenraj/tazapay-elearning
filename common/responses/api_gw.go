package responses

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"

	"tazapay.com/elearning/common/constants"
)

// apiResponse response framing for AWS-API-Gateway.
// defaults IsBase64Encoded sets to false, AccessControlAllowOrigin header sets to *.
// @param code: http response's status code
// @param body: API response's body
// @param headers: header values to be sent
// @param cookies: cookie values to be sent
func apiResponse(code int, body *APIResponseBody, headers map[string]string, cookies []string) *APIResponse {
	// marshal the response body
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		log.Error().Msgf("error marshaling api response body: %v", err.Error())
		return APIInternalError()
	}

	// frame aws api gateway response
	response := &APIResponse{
		IsBase64Encoded: constants.BoolFalse,
		Body:            string(bodyBytes),
		StatusCode:      code,
		Cookies:         cookies,
	}
	// add headers
	if len(headers) == constants.IntZero {
		headers = make(map[string]string)
	}
	headers[constants.HeaderKeyContentType] = constants.MediaTypeApplicationJSON // Required for client to receive body in JSON format
	response.Headers = headers

	// log the final response
	log.Debug().Msgf("response - statusCode: %v body: %v", code, string(bodyBytes))
	return response
}

// API convert common response object APIResponse
func API(res *Response) *APIResponse {
	body := APIResponseBody{
		Status:  res.Status,
		Message: res.Message,
	}

	// set data
	if res.Data != nil && res.Data.Value != nil {
		body.Data = res.Data.Value
	}

	// set error
	if res.Error != nil {
		body.Error = res.Error
	}

	// return api_gw response
	return apiResponse(res.HTTPStatusCode, &body, res.Headers, res.Cookies)
}

// APIInternalError frame API-GW response for internal server error.
// https://httpstatuses.com/500
func APIInternalError() *APIResponse {
	body := APIResponseBody{
		Status:  constants.StatusError,
		Message: constants.MsgInternalServerError,
	}
	return apiResponse(http.StatusInternalServerError, &body, nil, nil)
}

// APIServiceUnavailable frame API-GW response for service unavailable.
// https://httpstatuses.com/503
func APIServiceUnavailable() *APIResponse {
	body := APIResponseBody{
		Status:  constants.StatusError,
		Message: constants.MsgServiceUnavailable,
	}
	return apiResponse(http.StatusServiceUnavailable, &body, nil, nil)
}

// APIBadRequest frame API-GW response for bad request
// https://httpstatuses.com/400
func APIBadRequest(err *Err) *APIResponse {
	body := APIResponseBody{
		Status:  constants.StatusError,
		Message: constants.MsgBadRequest,
		Error:   err,
	}
	return apiResponse(http.StatusBadRequest, &body, nil, nil)
}

// APIMethodNotAllowed frame API-GW response for method not allowed.
// https://httpstatuses.com/405
func APIMethodNotAllowed() *APIResponse {
	body := APIResponseBody{
		Status:  constants.StatusError,
		Message: constants.MsgMethodNotAllowed,
	}
	return apiResponse(http.StatusMethodNotAllowed, &body, nil, nil)
}
