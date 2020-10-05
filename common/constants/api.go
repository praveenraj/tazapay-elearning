package constants

// response status
const (
	StatusSuccess  = "success"
	StatusCreated  = "created"
	StatusAccepted = "accepted"
	StatusError    = "error"
)

// response messages
const (
	MsgMethodNotAllowed    = "method not allowed"
	MsgServiceUnavailable  = "service unavailable"
	MsgInternalServerError = "internal server error"
	MsgBadRequest          = "bad request"
	MsgNotFound            = "resource not found"
	MsgConflict            = "conflict with the current state of the resource"
	MsgUnauthorized        = "unauthorized"
	MsgForbidden           = "forbidden to access the resource" // not authorized to access
	MsgUnprocessableEntity = "unprocessable entity: please check the semantic erroneous"
)

// headers
const (
	HeaderKeyContentType = "content-type"
)
