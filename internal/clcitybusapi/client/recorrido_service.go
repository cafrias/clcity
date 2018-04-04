package client

import (
	"encoding/json"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/geo"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi"
	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/soapclient/swparadas"
)

var _ clcitybusapi.RecorridoService = &RecorridoService{}

// RecorridoService has actions to fetch 'Parada' data from Cuando Llega City Bus API.
type RecorridoService struct {
	client       SOAPClient
	lineaService clcitybusapi.LineaService
	Path         string
}

// RecorridoDeLinea fetches a 'Recorrido' entity associated with a given 'Linea'.
func (s *RecorridoService) RecorridoDeLinea(l *clcitybusapi.Linea) (*clcitybusapi.Recorrido, error) {
	in := &swparadas.RecuperarRecorridoParaMapaPorEntidadYLinea{
		Usuario:           Usuario,
		Clave:             Clave,
		CodigoLineaParada: int32(l.Codigo),
		IsSublinea:        false,
	}
	res, err := s.client.RecuperarRecorridoParaMapaPorEntidadYLinea(in)
	if err != nil {
		return nil, err
	}

	result := new(swparadas.RecuperarRecorridoParaMapaPorEntidadYLineaResult)
	err = json.Unmarshal([]byte(res.RecuperarRecorridoParaMapaPorEntidadYLineaResult), result)
	if err != nil {
		return nil, err
	}

	// Map result to usable local struct.
	var points []geo.Point
	for _, punto := range result.Puntos {
		point := geo.Point{
			Lat:  punto.Latitud,
			Long: punto.Longitud,
		}
		points = append(points, point)
	}

	// Create result 'Recorrido' object.
	r := clcitybusapi.NewRecorrido(l, points)

	return r, nil
}

// RecorridosPorEmpresa fetches all 'Recorrido' entities associated with given 'Empresa' identified by `CodigoEmpresa`.
func (s *RecorridoService) RecorridosPorEmpresa(CodigoEmpresa int) ([]*clcitybusapi.Recorrido, error) {
	return nil, nil
}
