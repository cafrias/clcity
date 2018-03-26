package client

import (
	"errors"
	"io"
	"net/http"
)

const paradasURL = "http://clsw.smartmovepro.net/ModuloParadas/SWParadas.asmx"

// SOAPClient represents the SOAP client used to communicate with `Cuando Llega CityBus API`.
type SOAPClient struct {
	cli *http.Client
	req *http.Request
}

// Send sends the request to the `Cuando Llega City Bus` API.
func (c *SOAPClient) Send(body io.Reader) (*http.Response, error) {
	if body == nil {
		return nil, errors.New("should pass a body to the request")
	}

	if c.req.Header.Get("SOAPAction") == "" {
		return nil, errors.New("should pass a 'SOAPAction' header to the request")
	}

	return c.cli.Do(c.req)
}

// NewSOAPClient creates a new SOAPClient with default configuration for communicating with `Cuando Llega City Bus API`.
func NewSOAPClient() (*SOAPClient, error) {
	req, err := http.NewRequest("POST", paradasURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "ksoap2-android/2.6.0+")
	req.Header.Add("Content-Type", "text/xml;charset=utf-8")
	req.Header.Add("Connection", "close")
	req.Header.Add("Host", "clswsur.smartmovepro.net")

	return &SOAPClient{
		cli: &http.Client{},
		req: req,
	}, nil
}
