package mock

import (
	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/client"
	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/soapclient/swparadas"
)

var _ client.SOAPClient = &SOAPClient{}

// SOAPClient mock implementation of client.SOAPClient
type SOAPClient struct {
	*swparadas.SWParadasSoap
	RecuperarLineasPorCodigoEmpresaSpy            *Spy
	RecuperarParadasCompletoPorLineaSpy           *Spy
	RecuperarRecorridoParaMapaPorEntidadYLineaSpy *Spy
}

// RecuperarLineasPorCodigoEmpresa mock implementation of RecuperarLineasPorCodigoEmpresa for SOAPClient interface.
func (s *SOAPClient) RecuperarLineasPorCodigoEmpresa(request *swparadas.RecuperarLineasPorCodigoEmpresa) (*swparadas.RecuperarLineasPorCodigoEmpresaResponse, error) {
	s.RecuperarLineasPorCodigoEmpresaSpy.Invoked = true
	call := s.RecuperarLineasPorCodigoEmpresaSpy.Calls
	s.RecuperarLineasPorCodigoEmpresaSpy.Args = append(s.RecuperarLineasPorCodigoEmpresaSpy.Args, []interface{}{
		request,
	})

	ret1, _ := s.RecuperarLineasPorCodigoEmpresaSpy.Ret[call][0].(*swparadas.RecuperarLineasPorCodigoEmpresaResponse)
	ret2, _ := s.RecuperarLineasPorCodigoEmpresaSpy.Ret[call][1].(error)

	s.RecuperarLineasPorCodigoEmpresaSpy.Calls++

	return ret1, ret2
}

// RecuperarParadasCompletoPorLinea mock implementation of RecuperarParadasCompletoPorLinea for SOAPClient interface.
func (s *SOAPClient) RecuperarParadasCompletoPorLinea(request *swparadas.RecuperarParadasCompletoPorLinea) (*swparadas.RecuperarParadasCompletoPorLineaResponse, error) {
	s.RecuperarParadasCompletoPorLineaSpy.Invoked = true
	call := s.RecuperarParadasCompletoPorLineaSpy.Calls
	s.RecuperarParadasCompletoPorLineaSpy.Args = append(s.RecuperarParadasCompletoPorLineaSpy.Args, []interface{}{
		request,
	})

	ret1, _ := s.RecuperarParadasCompletoPorLineaSpy.Ret[call][0].(*swparadas.RecuperarParadasCompletoPorLineaResponse)
	ret2, _ := s.RecuperarParadasCompletoPorLineaSpy.Ret[call][1].(error)

	s.RecuperarParadasCompletoPorLineaSpy.Calls++
	return ret1, ret2
}

// RecuperarRecorridoParaMapaPorEntidadYLinea mock implementation of RecuperarRecorridoParaMapaPorEntidadYLinea for SOAPClient interface.
func (s *SOAPClient) RecuperarRecorridoParaMapaPorEntidadYLinea(request *swparadas.RecuperarRecorridoParaMapaPorEntidadYLinea) (*swparadas.RecuperarRecorridoParaMapaPorEntidadYLineaResponse, error) {
	s.RecuperarRecorridoParaMapaPorEntidadYLineaSpy.Invoked = true
	call := s.RecuperarRecorridoParaMapaPorEntidadYLineaSpy.Calls
	s.RecuperarRecorridoParaMapaPorEntidadYLineaSpy.Args = append(s.RecuperarRecorridoParaMapaPorEntidadYLineaSpy.Args, []interface{}{
		request,
	})

	ret1, _ := s.RecuperarRecorridoParaMapaPorEntidadYLineaSpy.Ret[call][0].(*swparadas.RecuperarRecorridoParaMapaPorEntidadYLineaResponse)
	ret2, _ := s.RecuperarRecorridoParaMapaPorEntidadYLineaSpy.Ret[call][1].(error)

	s.RecuperarRecorridoParaMapaPorEntidadYLineaSpy.Calls++
	return ret1, ret2
}
