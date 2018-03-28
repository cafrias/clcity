package client

// Client errors
const (
	ErrNotConnected = Error("Client isn't connected to a SOAPClient!")
)

// Error represents a WTF error.
type Error string

// Error returns the error message.
func (e Error) Error() string { return string(e) }
