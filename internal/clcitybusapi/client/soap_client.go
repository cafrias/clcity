package client

import (
	"bytes"
	"encoding/xml"
	"errors"
	"net/http"
)

const paradasURL = "http://clsw.smartmovepro.net/ModuloParadas/SWParadas.asmx"

// HTTPClient represents the underlying client used for HTTP calls.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// SOAPClient represents the SOAP client used to communicate with `Cuando Llega CityBus API`.
type SOAPClient struct {
	cli     HTTPClient
	headers map[string]string
}

// SOAPRequest represents a SOAP request to `Cuando Llega City Bus` API.
type SOAPRequest struct {
	XMLName           xml.Name `xml:"v:Envelope"`
	XMLInstanceSchema string   `xml:"xmlns:i,attr"`
	XMLSchema         string   `xml:"xmlns:d,attr"`
	XMLSOAP           string   `xml:"xmlns:c,attr"`
	XMLEnvelope       string   `xml:"xmlns:v,attr"`
	Header            SOAPHeader
	Body              SOAPRequestBody
}

// SOAPResponse represents a SOAP response from `Cuando Llega City Bus` API.
type SOAPResponse struct {
	XMLName           xml.Name `xml:"soap:Envelope"`
	XMLEnvelope       string   `xml:"xmlns:soap,attr"`
	XMLInstanceSchema string   `xml:"xmlns:xsi,attr"`
	XMLSchema         string   `xml:"xmlns:xsd,attr"`
	Body              SOAPResponseBody
}

// SOAPHeader represents a SOAP header.
type SOAPHeader struct {
	XMLName xml.Name `xml:"v:Header"`
}

// SOAPRequestBody represents a SOAP body for a request.
type SOAPRequestBody struct {
	XMLName xml.Name `xml:"v:Body"`
	Call    interface{}
}

// SOAPResponseBody represents a SOAP body for a response.
type SOAPResponseBody struct {
	XMLName xml.Name `xml:"soap:Body"`
	Call    interface{}
}

// NewSOAPRequest creates a new SOAPRequest
func NewSOAPRequest(body interface{}) SOAPRequest {
	return SOAPRequest{
		XMLInstanceSchema: "http://www.w3.org/2001/XMLSchema-instance",
		XMLSchema:         "http://www.w3.org/2001/XMLSchema",
		XMLSOAP:           "http://schemas.xmlsoap.org/soap/encoding/",
		XMLEnvelope:       "http://schemas.xmlsoap.org/soap/envelope/",
		Header:            SOAPHeader{},
		Body:              SOAPRequestBody{Call: body},
	}
}

// Send sends the request to the `Cuando Llega City Bus` API.
func (c *SOAPClient) Send(body []byte) (*http.Response, error) {
	if len(body) == 0 {
		return nil, errors.New("should pass a body to the request")
	}

	req, err := http.NewRequest("POST", paradasURL, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	for key, val := range c.headers {
		req.Header.Add(key, val)
	}

	return c.cli.Do(req)
}

// NewSOAPClient generates a new SOAPClient with defaults to make requests to `Cuando Llega City Bus` API.
func NewSOAPClient(cli HTTPClient) *SOAPClient {
	return &SOAPClient{
		cli: cli,
		headers: map[string]string{
			"User-Agent":   "ksoap2-android/2.6.0+",
			"Content-Type": "text/xml;charset=utf-8",
			"Connection":   "close",
			"Host":         "clswsur.smartmovepro.net",
		},
	}
}
