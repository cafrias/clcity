package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/dump"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi"
	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/soapclient/swparadas"
)

var _ clcitybusapi.LineaService = &LineaService{}

// LineaService has actions to fetch 'Parada' data from Cuando Llega City Bus API.
type LineaService struct {
	client SOAPClient
	Path   string
}

func (s *LineaService) mapLineaFromSW(swl *swparadas.Linea) (*clcitybusapi.Linea, error) {
	codEnt, err := strconv.Atoi(swl.CodigoEntidad)
	if err != nil {
		return nil, err
	}

	codPar, err := strconv.Atoi(swl.CodigoLineaParada)
	if err != nil {
		return nil, err
	}

	return &clcitybusapi.Linea{
		Codigo:        codPar,
		CodigoEmpresa: swl.CodigoEmpresa,
		CodigoEntidad: codEnt,
		Descripcion:   swl.Descripcion,
	}, nil
}

// LineasPorEmpresa fetches all 'Parada' entities associated with a given 'Linea' identified by the code passed as `CodigoLineaParada`.
func (s *LineaService) LineasPorEmpresa(CodigoEmpresa int) ([]*clcitybusapi.Linea, error) {
	outFile := fmt.Sprintf("%s/%s", s.Path, "lineas.json")

	var ret []*clcitybusapi.Linea
	if ok := dump.Read(&ret, outFile); ok == true {
		return ret, nil
	}

	in := &swparadas.RecuperarLineasPorCodigoEmpresa{
		Usuario:       Usuario,
		Clave:         Clave,
		CodigoEmpresa: int32(CodigoEmpresa),
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

	// Map result to local struct
	for _, linea := range result.Lineas {
		l, err := s.mapLineaFromSW(linea)
		if err != nil {
			return nil, err
		}
		ret = append(ret, l)
	}

	// Write JSON file
	err = dump.Write(ret, outFile)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
