package client

import (
	"encoding/json"
	"errors"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi"
	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/soapclient/swparadas"
)

var _ clcitybusapi.ParadaService = &ParadaService{}

// ParadaService has actions to fetch 'Parada' data from Cuando Llega City Bus API.
type ParadaService struct {
	client       SOAPClient
	lineaService LineaService
}

// ParadasPorLinea fetches all 'Parada' entities associated with a given 'Linea' identified by the code passed as `CodigoLineaParada`.
func (s *ParadaService) ParadasPorLinea(CodigoLineaParada int) ([]*clcitybusapi.Parada, error) {
	if s.client == nil {
		return nil, ErrNotConnected
	}

	in := &swparadas.RecuperarParadasCompletoPorLinea{
		Usuario:           "WEB.SUR",
		Clave:             "PAR.SW.SUR",
		CodigoLineaParada: int32(CodigoLineaParada),
		IsSublinea:        false,
		IsInteligente:     false,
	}
	res, err := s.client.RecuperarParadasCompletoPorLinea(in)
	if err != nil {
		return nil, err
	}

	result := new(swparadas.RecuperarParadasCompletoPorLineaResult)
	err = json.Unmarshal([]byte(res.RecuperarParadasCompletoPorLineaResult), result)
	if err != nil {
		return nil, err
	}

	if result.CodigoEstado != 0 {
		return nil, errors.New(result.MensajeEstado)
	}

	return result.Paradas, nil
}

// ParadasPorEmpresa fetches all 'Parada' entities associated with given 'Empresa' identified by `CodigoEmpresa`.
func (s *ParadaService) ParadasPorEmpresa(CodigoEmpresa int) ([]*clcitybusapi.Parada, error) {
	if s.client == nil {
		return nil, ErrNotConnected
	}

	lineas, err := s.lineaService.LineasPorEmpresa(CodigoEmpresa)
	if err != nil {
		return nil, err
	}

	// paradas := make(chan []*clcitybusapi.Parada, )

	// for _, linea := range l {

	// }

	// in := &swparadas.RecuperarParadasCompletoPorLinea{
	// 	Usuario:           "WEB.SUR",
	// 	Clave:             "PAR.SW.SUR",
	// 	CodigoLineaParada: int32(CodigoLineaParada),
	// 	IsSublinea:        false,
	// 	IsInteligente:     false,
	// }
	// res, err := s.client.RecuperarParadasCompletoPorLinea(in)
	// if err != nil {
	// 	return nil, err
	// }

	// result := new(swparadas.RecuperarParadasCompletoPorLineaResult)
	// err = json.Unmarshal([]byte(res.RecuperarParadasCompletoPorLineaResult), result)
	// if err != nil {
	// 	return nil, err
	// }

	// if result.CodigoEstado != 0 {
	// 	return nil, errors.New(result.MensajeEstado)
	// }

	return nil, nil
}
