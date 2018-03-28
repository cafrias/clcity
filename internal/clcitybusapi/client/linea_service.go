package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi"
	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/soapclient/swparadas"
)

var _ clcitybusapi.LineaService = &LineaService{}

// LineaService has actions to fetch 'Parada' data from Cuando Llega City Bus API.
type LineaService struct {
	client SOAPClient
	Path   string
}

// LineasPorEmpresa fetches all 'Parada' entities associated with a given 'Linea' identified by the code passed as `CodigoLineaParada`.
func (s *LineaService) LineasPorEmpresa(CodigoEmpresa int) ([]*clcitybusapi.Linea, error) {
	outFile := fmt.Sprintf("%s/%s", s.Path, "lineas.json")

	// If we can't find a dump file, fetch data from endpoint
	if _, err := os.Stat(outFile); os.IsNotExist(err) {
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

		// Write JSON file
		fcontent, err := json.MarshalIndent(result.Lineas, "", "    ")
		if err != nil {
			return nil, err
		}
		err = ioutil.WriteFile(outFile, fcontent, 0644)
		if err != nil {
			return nil, err
		}

		return result.Lineas, nil
	}

	return nil, nil
}
