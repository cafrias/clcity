package client_test

import (
	"fmt"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/soapclient"
	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/soapclient/swparadas"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/client"
)

type SOAPClient struct {
	*soapclient.SOAPClient
	// Spy on Call
	CallInvoked    bool
	CallSoapAction string
	CallReq        interface{}
	CallRes        interface{}
	CallError      error
}

func (s *SOAPClient) Call(soapAction string, request, response interface{}) error {
	s.CallInvoked = true
	s.CallSoapAction = soapAction
	s.CallReq = request

	// Mock results
	fmt.Println("Received: ", s.CallRes)
	response = &s.CallRes
	return s.CallError
}

func NewSWParadasSoap(url string, tls bool, auth *soapclient.BasicAuth) (*SOAPClient, *swparadas.SWParadasSoap) {
	if url == "" {
		url = "http://clsw.smartmovepro.net/ModuloParadas/SWParadas.asmx"
	}
	client := SOAPClient{
		SOAPClient: soapclient.NewSOAPClient(url, tls, auth),
	}

	return &client, swparadas.NewSWParadasSoapWithClient(&client)
}

// Client test wrapper for `client.Client`
type Client struct {
	*client.Client
}
