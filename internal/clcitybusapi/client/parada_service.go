package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/dump"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/geo"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi"
	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/soapclient/swparadas"
)

var _ clcitybusapi.ParadaService = &ParadaService{}

// ParadaService has actions to fetch 'Parada' data from Cuando Llega City Bus API.
type ParadaService struct {
	client       SOAPClient
	lineaService clcitybusapi.LineaService
	Path         string
}

func (s *ParadaService) mapParadaFromSW(swp *swparadas.Parada) (*clcitybusapi.Parada, error) {
	cod, err := strconv.Atoi(swp.Codigo)
	if err != nil {
		return nil, err
	}

	latStr := strings.Replace(swp.LatitudParada, ",", ".", -1)
	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		return nil, err
	}

	longStr := strings.Replace(swp.LongitudParada, ",", ".", -1)
	long, err := strconv.ParseFloat(longStr, 64)
	if err != nil {
		return nil, err
	}

	return &clcitybusapi.Parada{
		Codigo:                     cod,
		Identificador:              swp.Identificador,
		Descripcion:                swp.Descripcion,
		AbreviaturaBanderaGIT:      swp.AbreviaturaBanderaGIT,
		AbreviaturaBandera:         swp.AbreviaturaBandera,
		AbreviaturaAmpliadaBandera: swp.AbreviaturaAmpliadaBandera,
		Punto: geo.Point{
			Lat:  lat,
			Long: long,
		},
	}, nil
}

// ParadasPorLinea fetches all 'Parada' entities associated with a given 'Linea' identified by the code passed as `CodigoLineaParada`.
func (s *ParadaService) ParadasPorLinea(CodigoLineaParada int) ([]*clcitybusapi.Parada, error) {
	outFile := fmt.Sprintf("%s/paradas_linea_%v.json", s.Path, CodigoLineaParada)

	// If we can't find a dump file, fetch data from endpoint
	if _, err := os.Stat(outFile); os.IsNotExist(err) {
		in := &swparadas.RecuperarParadasCompletoPorLinea{
			Usuario:           Usuario,
			Clave:             Clave,
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

		// Map to local struct
		var r []*clcitybusapi.Parada
		for _, parada := range result.Paradas {
			p, err := s.mapParadaFromSW(parada)
			if err != nil {
				return nil, err
			}
			r = append(r, p)
		}

		// Write dump file
		err = dump.Write(r, outFile)
		if err != nil {
			return nil, err
		}

		return r, nil
	}

	c, err := ioutil.ReadFile(outFile)
	if err != nil {
		return nil, err
	}

	var r []*clcitybusapi.Parada
	err = json.Unmarshal(c, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// ParadasPorEmpresa fetches all 'Parada' entities associated with given 'Empresa' identified by `CodigoEmpresa`.
func (s *ParadaService) ParadasPorEmpresa(CodigoEmpresa int) ([]*clcitybusapi.Parada, error) {
	outFile := fmt.Sprintf("%s/paradas_empresa.json", s.Path)

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
			res, err := s.ParadasPorLinea(linea.Codigo)
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

	// Write dump file
	err = dump.Write(paradas, outFile)
	if err != nil {
		return nil, err
	}

	return paradas, nil
}
