package mock

import (
	"io"
	"net/http"
)

// SOAPClient mock for `clcitybusapi.SOAPClient`.
type SOAPClient struct {
	SendFn      func(body io.Reader) (*http.Response, error)
	SendInvoked bool
}

// Send mock for `clcitybusapi.SOAPClient.Send`.
func (s *SOAPClient) Send(body io.Reader) (*http.Response, error) {
	s.SendInvoked = true
	return s.SendFn(body)

}
