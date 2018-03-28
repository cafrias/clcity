package mock

import (
	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/client"
	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/soapclient/swparadas"
)

var _ client.SOAPClient = &SOAPClient{}

// SOAPClient mock implementation of client.SOAPClient
type SOAPClient struct {
	*swparadas.SWParadasSoap
	RecuperarLineasPorCodigoEmpresaSpy  *Spy
	RecuperarParadasCompletoPorLineaSpy *Spy
}

// RecuperarLineasPorCodigoEmpresa mock implementation of RecuperarLineasPorCodigoEmpresa for SOAPClient interface.
func (s *SOAPClient) RecuperarLineasPorCodigoEmpresa(request *swparadas.RecuperarLineasPorCodigoEmpresa) (*swparadas.RecuperarLineasPorCodigoEmpresaResponse, error) {
	s.RecuperarLineasPorCodigoEmpresaSpy.Invoked = true
	s.RecuperarLineasPorCodigoEmpresaSpy.Args = append(s.RecuperarLineasPorCodigoEmpresaSpy.Args, request)

	ret1, _ := s.RecuperarLineasPorCodigoEmpresaSpy.Ret[0].(*swparadas.RecuperarLineasPorCodigoEmpresaResponse)
	ret2, _ := s.RecuperarLineasPorCodigoEmpresaSpy.Ret[1].(error)

	return ret1, ret2
}

// RecuperarParadasCompletoPorLinea mock implementation of RecuperarParadasCompletoPorLinea for SOAPClient interface.
func (s *SOAPClient) RecuperarParadasCompletoPorLinea(request *swparadas.RecuperarParadasCompletoPorLinea) (*swparadas.RecuperarParadasCompletoPorLineaResponse, error) {
	s.RecuperarParadasCompletoPorLineaSpy.Invoked = true
	s.RecuperarParadasCompletoPorLineaSpy.Args = append(s.RecuperarParadasCompletoPorLineaSpy.Args, request)

	ret1, _ := s.RecuperarParadasCompletoPorLineaSpy.Ret[0].(*swparadas.RecuperarParadasCompletoPorLineaResponse)
	ret2, _ := s.RecuperarParadasCompletoPorLineaSpy.Ret[1].(error)

	return ret1, ret2
}
