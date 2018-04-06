package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"image/color"
	"log"
	"strconv"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/dump"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi"
	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/soapclient/swparadas"
)

var _ clcitybusapi.LineaService = &LineaService{}

// LineaService has actions to fetch 'ParadaLinea' data from Cuando Llega City Bus API.
type LineaService struct {
	scli SOAPClient
	cli  clcitybusapi.Client
	Path string
}

var colorMap = map[int]color.RGBA{
	1529: color.RGBA{R: 66, G: 134, B: 244, A: 1},
	1530: color.RGBA{R: 104, G: 244, B: 66, A: 1},
	1531: color.RGBA{R: 244, G: 66, B: 66, A: 1},
	1532: color.RGBA{R: 244, G: 244, B: 66, A: 1},
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
		CodigoEntidad: codEnt,
		Color:         colorMap[codPar],
		Descripcion:   swl.Descripcion,
	}, nil
}

// LineasPorEmpresa fetches all 'ParadaLinea' entities associated with a given 'Linea' identified by the code passed as `CodigoLineaParada`.
func (s *LineaService) LineasPorEmpresa(empresa *clcitybusapi.Empresa) ([]*clcitybusapi.Linea, error) {
	outFile := fmt.Sprintf("%s/%s", s.Path, "lineas.json")

	var ret []*clcitybusapi.Linea
	if ok := dump.Read(&ret, outFile); ok == true {
		log.Printf("Reading lineas from dump file: `%s`\n", outFile)
		return ret, nil
	}

	in := &swparadas.RecuperarLineasPorCodigoEmpresa{
		Usuario:       Usuario,
		Clave:         Clave,
		CodigoEmpresa: int32(empresa.Codigo),
		IsSublinea:    false,
	}
	res, err := s.scli.RecuperarLineasPorCodigoEmpresa(in)
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
		l.Empresa = empresa

		// Fetch Recorrido
		rec, err := s.cli.RecorridoService().RecorridoDeLinea(l)
		if err != nil {
			return nil, err
		}
		l.Recorrido = rec

		// Fetch Paradas
		par, err := s.cli.ParadaService().ParadasPorLinea(l)
		if err != nil {
			return nil, err
		}
		l.Paradas = par

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

	log.Printf("Wrote lineas to dump file: `%s`\n", outFile)

	return ret, nil
}
