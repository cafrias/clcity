package client

import (
	"encoding/json"
	"errors"
	"strconv"
	"sync"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi"
	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/soapclient/swparadas"
)

var _ clcitybusapi.ParadaService = &ParadaService{}

// ParadaService has actions to fetch 'Parada' data from Cuando Llega City Bus API.
type ParadaService struct {
	client       SOAPClient
	lineaService clcitybusapi.LineaService
}

// ParadasPorLinea fetches all 'Parada' entities associated with a given 'Linea' identified by the code passed as `CodigoLineaParada`.
func (s *ParadaService) ParadasPorLinea(CodigoLineaParada int) ([]*clcitybusapi.Parada, error) {
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
	lineas, err := s.lineaService.LineasPorEmpresa(CodigoEmpresa)
	if err != nil {
		return nil, err
	}

	type RequestResult struct {
		Value []*clcitybusapi.Parada
		Error error
	}

	var wg sync.WaitGroup

	lineasQty := len(lineas)
	wg.Add(lineasQty)

	pStream := make(chan *RequestResult, lineasQty)
	for _, linea := range lineas {
		go func(linea *clcitybusapi.Linea) {
			defer wg.Done()
			result := new(RequestResult)
			cod, err := strconv.Atoi(linea.CodigoLineaParada)
			if err != nil {
				result.Error = err
				pStream <- result
				return
			}
			res, err := s.ParadasPorLinea(cod)
			if err != nil {
				result.Error = err
				pStream <- result
				return
			}
			result.Value = res
			pStream <- result
			return
		}(linea)
	}

	wg.Wait()

	var paradas []*clcitybusapi.Parada
	for i := 0; i < lineasQty; i++ {
		result := <-pStream
		if result.Error != nil {
			return nil, result.Error
		}
		paradas = append(paradas, result.Value...)
	}

	return paradas, nil
}
