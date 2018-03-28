package client

import (
	"encoding/json"
	"errors"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi"
	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/soapclient/swparadas"
)

var _ clcitybusapi.LineaService = &LineaService{}

// LineaService has actions to fetch 'Parada' data from Cuando Llega City Bus API.
type LineaService struct {
	client SOAPClient
}

// LineasPorEmpresa fetches all 'Parada' entities associated with a given 'Linea' identified by the code passed as `CodigoLineaParada`.
func (s *LineaService) LineasPorEmpresa(CodigoEmpresa int) ([]*clcitybusapi.Linea, error) {
	in := &swparadas.RecuperarLineasPorCodigoEmpresa{
		Usuario:       "WEB.SUR",
		Clave:         "PAR.SW.SUR",
		CodigoEmpresa: 355,
		IsSublinea:    false,
	}
	res, err := s.client.RecuperarLineasPorCodigoEmpresa(in)
	if err != nil {
		return nil, err
	}

	result := new(swparadas.RecuperarLineasPorCodigoEmpresaResult)
	err = json.Unmarshal([]byte(res.RecuperarLineasPorCodigoEmpresaResult), result)
	if err != nil {
		return nil, err
	}

	if result.CodigoEstado != 0 {
		return nil, errors.New(result.MensajeEstado)
	}

	return result.Lineas, nil
}
