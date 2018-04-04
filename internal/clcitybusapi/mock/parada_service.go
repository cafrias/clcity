package mock

import (
	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi"
	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/client"
)

var _ clcitybusapi.ParadaService = &ParadaService{}

// ParadaService mock implementation of client.ParadaService
type ParadaService struct {
	*client.ParadaService
	ParadasPorLineaSpy   *Spy
	ParadasPorEmpresaSpy *Spy
}

// ParadasPorLinea mock implementation of ParadasPorLinea for ParadaService.
func (s *ParadaService) ParadasPorLinea(l *clcitybusapi.Linea) ([]*clcitybusapi.ParadaLinea, error) {
	s.ParadasPorLineaSpy.Invoked = true
	call := s.ParadasPorLineaSpy.Calls
	s.ParadasPorLineaSpy.Args = append(s.ParadasPorLineaSpy.Args, []interface{}{
		l,
	})

	ret1, _ := s.ParadasPorLineaSpy.Ret[call][0].([]*clcitybusapi.ParadaLinea)
	ret2, _ := s.ParadasPorLineaSpy.Ret[call][1].(error)

	s.ParadasPorLineaSpy.Calls++

	return ret1, ret2
}

// ParadasPorEmpresa mock implementation of ParadasPorEmpresa for ParadaService.
func (s *ParadaService) ParadasPorEmpresa(e *clcitybusapi.Empresa) ([]*clcitybusapi.Parada, error) {
	s.ParadasPorEmpresaSpy.Invoked = true
	call := s.ParadasPorEmpresaSpy.Calls
	s.ParadasPorEmpresaSpy.Args = append(s.ParadasPorEmpresaSpy.Args, []interface{}{
		e,
	})

	ret1, _ := s.ParadasPorEmpresaSpy.Ret[call][0].([]*clcitybusapi.Parada)
	ret2, _ := s.ParadasPorEmpresaSpy.Ret[call][1].(error)

	s.ParadasPorEmpresaSpy.Calls++

	return ret1, ret2
}
