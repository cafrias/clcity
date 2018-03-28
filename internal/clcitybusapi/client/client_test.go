package client_test

import (
	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/soapclient"
	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/soapclient/swparadas"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/client"
)

type Spy struct {
	Invoked bool
	Args    []interface{}
	Ret     []interface{}
}

type SOAPClient struct {
	*swparadas.SWParadasSoap
	RecuperarLineasPorCodigoEmpresaSpy *Spy
}

func (s *SOAPClient) RecuperarLineasPorCodigoEmpresa(request *swparadas.RecuperarLineasPorCodigoEmpresa) (*swparadas.RecuperarLineasPorCodigoEmpresaResponse, error) {
	s.RecuperarLineasPorCodigoEmpresaSpy.Invoked = true
	s.RecuperarLineasPorCodigoEmpresaSpy.Args = append(s.RecuperarLineasPorCodigoEmpresaSpy.Args, request)

	ret1, _ := s.RecuperarLineasPorCodigoEmpresaSpy.Ret[0].(*swparadas.RecuperarLineasPorCodigoEmpresaResponse)
	ret2, _ := s.RecuperarLineasPorCodigoEmpresaSpy.Ret[1].(error)

	return ret1, ret2
}

func NewSOAPClient(url string, tls bool, auth *soapclient.BasicAuth) *SOAPClient {
	if url == "" {
		url = "http://clsw.smartmovepro.net/ModuloParadas/SWParadas.asmx"
	}
	client := &SOAPClient{
		SWParadasSoap: swparadas.NewSWParadasSoap(url, tls, auth),
	}

	return client
}

// Client test wrapper for `client.Client`
type Client struct {
	*client.Client
}
