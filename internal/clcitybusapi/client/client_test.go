package client_test

import (
	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/mock"
	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/soapclient"
	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/soapclient/swparadas"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/client"
)

// Client test wrapper for `client.Client`
type Client struct {
	*client.Client
}

func NewSOAPClient(url string, tls bool, auth *soapclient.BasicAuth) *mock.SOAPClient {
	if url == "" {
		url = "http://clsw.smartmovepro.net/ModuloParadas/SWParadas.asmx"
	}
	client := &mock.SOAPClient{
		SWParadasSoap: swparadas.NewSWParadasSoap(url, tls, auth),
	}

	return client
}
