package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

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

	// If we can't find a dump file, fetch data from endpoint
	if _, err := os.Stat(outFile); os.IsNotExist(err) {
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
		var r []*clcitybusapi.Linea
		for _, linea := range result.Lineas {
			l, err := s.mapLineaFromSW(linea)
			if err != nil {
				return nil, err
			}
			r = append(r, l)
		}

		// Write JSON file
		fcontent, err := json.MarshalIndent(r, "", "    ")
		if err != nil {
			return nil, err
		}
		err = ioutil.WriteFile(outFile, fcontent, 0644)
		if err != nil {
			return nil, err
		}

		return r, nil
	}

	return nil, nil
}
