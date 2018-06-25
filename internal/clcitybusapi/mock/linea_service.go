package mock

import (
	"github.com/cafrias/clcity/internal/clcitybusapi"
	"github.com/cafrias/clcity/internal/clcitybusapi/client"
)

var _ clcitybusapi.LineaService = &LineaService{}

// LineaService mock implementation of client.LineaService
type LineaService struct {
	*client.LineaService
	LineasPorEmpresaSpy *Spy
}

// LineasPorEmpresa mock implementation of LineasPorEmpresa for LineaService.
func (s *LineaService) LineasPorEmpresa(e *clcitybusapi.Empresa) ([]*clcitybusapi.Linea, error) {
	s.LineasPorEmpresaSpy.Invoked = true
	call := s.LineasPorEmpresaSpy.Calls
	s.LineasPorEmpresaSpy.Args = append(s.LineasPorEmpresaSpy.Args, []interface{}{
		e,
	})

	ret1, _ := s.LineasPorEmpresaSpy.Ret[call][0].([]*clcitybusapi.Linea)
	ret2, _ := s.LineasPorEmpresaSpy.Ret[call][1].(error)

	s.LineasPorEmpresaSpy.Calls++

	return ret1, ret2
}
