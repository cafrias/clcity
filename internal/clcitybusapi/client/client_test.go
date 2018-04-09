package client_test

import (
	"os"

	"bitbucket.org/friasdesign/clcity/internal/clcitybusapi/mock"
	"bitbucket.org/friasdesign/clcity/internal/clcitybusapi/soapclient"
	"bitbucket.org/friasdesign/clcity/internal/clcitybusapi/soapclient/swparadas"

	"bitbucket.org/friasdesign/clcity/internal/clcitybusapi/client"
)

const DumpPath = "testdata/dump"

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

func CreateDump() {
	os.Mkdir(DumpPath, 0764)
}

func ClearDump() {
	os.RemoveAll(DumpPath)
}
