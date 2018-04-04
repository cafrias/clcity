package client

import (
	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi"
)

var _ clcitybusapi.EmpresaService = &EmpresaService{}

// EmpresaService has actions to fetch 'ParadaLinea' data from Cuando Llega City Bus API.
type EmpresaService struct {
	client        SOAPClient
	lineaService  clcitybusapi.LineaService
	paradaService clcitybusapi.ParadaService
	Path          string
}

// ObtenerLineas returns all 'Linea' entities for given 'Empresa'.
func (s *EmpresaService) ObtenerLineas(e *clcitybusapi.Empresa) error {
	return nil
}
