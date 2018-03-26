package mock

import "bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi"

// ParadaService mock for `clcitybusapi.ParadaService`.
type ParadaService struct {
	ParadasPorLineaFn      func(CodigoLineaParada string) ([]*clcitybusapi.Parada, error)
	ParadasPorLineaInvoked bool
}

// ParadasPorLinea mock for `clcitybusapi.ParadaService.ParadasPorLinea`.
func (s *ParadaService) ParadasPorLinea(CodigoLineaParada string) ([]*clcitybusapi.Parada, error) {
	s.ParadasPorLineaInvoked = true
	return s.ParadasPorLineaFn(CodigoLineaParada)

}
