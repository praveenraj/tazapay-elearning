package responses

/*
	Common response class for the application.
*/

import (
	"net/http"

	"tazapay.com/elearning/common/constants"
)

// response builds *Response struct.
func response(statuscode int, status string, meta *MetaData) *Response {
	return &Response{
		HTTPStatusCode: statuscode,
		Status:         status,
		MetaData:       *meta,
	}
}

// parseMetaData extract message, Data, Err, headers, cookies from the args
// args...
// @type string: response message
// @type Data/interface{}: data to be consumed
// @type Err: custom err object to tell what exactly happened
// @type map[string]string: headers to be added to the API response
// @type []string: cookies to be added to the API response
func parseMetaData(args ...interface{}) *MetaData {
	// init meta obj to assign data
	meta := MetaData{}
	for _, arg := range args {
		switch v := arg.(type) {
		// headers
		case map[string]string:
			meta.Headers = v

		// cookies
		case []string:
			meta.Cookies = v

		// message
		case string:
			meta.Message = v

		// Err
		case *Err:
			meta.Error = v
		case Err:
			meta.Error = &v

		// Data
		case *Data:
			meta.Data = v
		case Data:
			meta.Data = &v
		case interface{}: // interface{} satisfies all data types so it should be the last check
			meta.Data = &Data{
				Value: v,
			}
		}
	}
	return &meta
}

// Success https://httpstatuses.com/200.
// The request has succeeded.
// args...
// @type string: response message
// @type Data/interface{}: data to be consumed
// @type map[string]string: headers to be added to the API response
// @type []string: cookies to be added to the API response
func Success(args ...interface{}) *Response {
	return response(http.StatusOK, constants.StatusSuccess, parseMetaData(args...))
}

// Error build common error response.
// @param statusCode: http status code for the error.
// @param msg: error response message.
// args...
// @type Err: custom err object to tell what exactly happened.
// @type Data/interface{}: data to be used for other unusal controllers like (alexa, lex, connect, etc.)
func Error(statusCode int, msg string, args ...interface{}) *Response {
	args = append(args, msg)
	return response(statusCode, constants.StatusError, parseMetaData(args...))
}

// InternalError https://httpstatuses.com/500
// args...
// @type Err: custom err object to tell what exactly happened.
// @type Data/interface{}: data to be used for other unusal controllers like (alexa, lex, connect, etc.)
func InternalError(args ...interface{}) *Response {
	return Error(http.StatusInternalServerError, constants.MsgInternalServerError, args...)
}

// BadRequest https://httpstatuses.com/400
// The server cannot or will not process the request due to something that is perceived to be a client error
// (e.g., malformed request syntax, invalid request message framing, or deceptive request routing).
// args...
// @type Err: custom err object to tell what exactly happened.
// @type Data/interface{}: data to be used for other unusal controllers like (alexa, lex, connect, etc.)
func BadRequest(args ...interface{}) *Response {
	return Error(http.StatusBadRequest, constants.MsgBadRequest, args...)
}

// UnAuthorized https://httpstatuses.com/401
// The request has not been applied because it lacks valid authentication credentials.
// args...
// @type Err: custom err object to tell what exactly happened.
// @type Data/interface{}: data to be used for other unusal controllers like (alexa, lex, connect, etc.)
func UnAuthorized(args ...interface{}) *Response {
	return Error(http.StatusUnauthorized, constants.MsgUnauthorized, args...)
}

// Forbidden https://httpstatuses.com/403
// The server understood the request but refuses to authorize it.
// Primarily due to a lack of permission to access the requested resource.
// args...
// @type Err: custom err object to tell what exactly happened.
// @type Data/interface{}: data to be used for other unusal controllers like (alexa, lex, connect, etc.)
func Forbidden(args ...interface{}) *Response {
	return Error(http.StatusForbidden, constants.MsgForbidden, args...)
}

// NotFound https://httpstatuses.com/404
// The origin server did not find a current representation for the target resource.
// args...
// @type Err: custom err object to tell what exactly happened.
// @type Data/interface{}: data to be used for other unusal controllers like (alexa, lex, connect, etc.)
func NotFound(args ...interface{}) *Response {
	return Error(http.StatusNotFound, constants.MsgNotFound, args...)
}

// Conflict https://httpstatuses.com/409
// The request could not be completed due to a conflict with the current state of the resource.
// Ex: resource already exists or should be in this state to process the resource.
// args... in any order.
// @type Err: (mandate) custom err object to tell what exactly happened.
// @type Data/interface{}: data to be used for other unusal controllers like (alexa, lex, connect, etc.)
func Conflict(args ...interface{}) *Response {
	return Error(http.StatusConflict, constants.MsgConflict, args...)
}
