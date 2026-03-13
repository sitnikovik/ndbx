package resp

// ErrorResponse represents an error response from the server.
type ErrorResponse struct {
	// msg is the error message extracted from the response body.
	msg string
}

// NewErrorResponse creates a new ErrorResponse with the provided message.
func NewErrorResponse(msg string) ErrorResponse {
	return ErrorResponse{
		msg: msg,
	}
}

// Error returns the error message of the ErrorResponse.
func (e ErrorResponse) Error() string {
	return e.msg
}
