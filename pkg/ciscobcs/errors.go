package ciscobcs

// Err implements the error interface so we can have constant errors.
type Err string

func (e Err) Error() string {
	return string(e)
}

// Error Constants
const (
	ErrMissingAPIKey = Err("ciscobcs: missing required apikey")

	ErrBadRequest    = Err("ciscobcs: bad request")
	ErrUnauthorized  = Err("ciscobcs: unauthorized request")
	ErrForbidden     = Err("ciscobcs: forbidden")
	ErrInternalError = Err("ciscobcs: internal error")
	ErrUnknown       = Err("ciscobcs: unexpected error occurred")
)
