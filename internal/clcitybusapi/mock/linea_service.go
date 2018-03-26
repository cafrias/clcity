package mock

import "bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi"

// LineaService mock for `clcitybusapi.LineaService`.
type LineaService struct {
	LineasPorEmpresaFn      func(CodigoEmpresa string) ([]*clcitybusapi.Linea, error)
	LineasPorEmpresaInvoked bool
}

// LineasPorEmpresa mock for `clcitybusapi.LineaService.LineasPorEmpresa`.
func (s *LineaService) LineasPorEmpresa(CodigoEmpresa string) ([]*clcitybusapi.Linea, error) {
	s.LineasPorEmpresaInvoked = true
	return s.LineasPorEmpresaFn(CodigoEmpresa)

}
