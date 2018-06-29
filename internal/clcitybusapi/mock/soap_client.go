package mock

import (
	"github.com/cafrias/clcity/internal/clcitybusapi/client"
	"github.com/cafrias/clcity/internal/clcitybusapi/soapclient/swparadas"
)

var _ client.SOAPClient = &SOAPClient{}

// SOAPClient mock implementation of swparadas.SOAPClient
type SOAPClient struct {
	*swparadas.SWParadasSoap
	RecuperarLineasPorCodigoEmpresaSpy            *Spy
	RecuperarParadasPorLineaParaCuandoLlegaSpy    *Spy
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

// RecuperarParadasPorLineaParaCuandoLlega mock implementation of RecuperarParadasPorLineaParaCuandoLlega for SOAPClient interface.
func (s *SOAPClient) RecuperarParadasPorLineaParaCuandoLlega(request *swparadas.RecuperarParadasPorLineaParaCuandoLlega) (*swparadas.RecuperarParadasPorLineaParaCuandoLlegaResponse, error) {
	s.RecuperarParadasPorLineaParaCuandoLlegaSpy.Invoked = true
	call := s.RecuperarParadasPorLineaParaCuandoLlegaSpy.Calls
	s.RecuperarParadasPorLineaParaCuandoLlegaSpy.Args = append(s.RecuperarParadasPorLineaParaCuandoLlegaSpy.Args, []interface{}{
		request,
	})

	ret1, _ := s.RecuperarParadasPorLineaParaCuandoLlegaSpy.Ret[call][0].(*swparadas.RecuperarParadasPorLineaParaCuandoLlegaResponse)
	ret2, _ := s.RecuperarParadasPorLineaParaCuandoLlegaSpy.Ret[call][1].(error)

	s.RecuperarParadasPorLineaParaCuandoLlegaSpy.Calls++
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
