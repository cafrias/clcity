package client

// General errors.
const (
	ErrRequestFailed = Error("request failed with status other than OK.")
)

// Error represents a client error.
type Error string

// Error returns the error message.
func (e Error) Error() string { return string(e) }
