package client

import (
	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi"
	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/soapclient/swparadas"
)

var _ clcitybusapi.ParadaService = &ParadaService{}

// ParadaService has actions to fetch 'Parada' data from Cuando Llega City Bus API.
type ParadaService struct {
	client *swparadas.SWParadasSoap
}

// ParadasPorLinea fetches all 'Parada' entities associated with a given 'Linea' identified by the code passed as `CodigoLineaParada`.
func (s *ParadaService) ParadasPorLinea(CodigoLineaParada string) ([]*clcitybusapi.Parada, error) {
	return nil, nil
}
