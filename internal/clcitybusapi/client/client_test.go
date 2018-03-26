package client_test

import (
	"net/http"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/client"
)

// HTTPClient test wrapper over `client.HTTPClient`
type HTTPClient struct {
	DoInvoked bool
	DoErr     error
	Req       *http.Request
	Res       *http.Response
}

// Do mocks a call to Do
func (c *HTTPClient) Do(req *http.Request) (*http.Response, error) {
	c.DoInvoked = true
	c.Req = req
	return c.Res, c.DoErr
}

// Client test wrapper for `client.Client`
type Client struct {
	*client.Client
}

// NewClient returns a new instance of Client with a mocked SOAPClient.
func NewClient(cli *HTTPClient) (*Client, error) {
	scli, err := client.NewSOAPClient(cli)
	if err != nil {
		return nil, err
	}

	return &Client{
		Client: client.NewClient(scli),
	}, nil
}
