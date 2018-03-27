package client

import (
	"encoding/json"
	"errors"
	"fmt"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi"
	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/soapclient/swparadas"
)

var _ clcitybusapi.LineaService = &LineaService{}

// LineaService has actions to fetch 'Parada' data from Cuando Llega City Bus API.
type LineaService struct {
	client *swparadas.SWParadasSoap
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
	fmt.Println("Result: ", res)
	if err != nil {
		return nil, err
	}

	var result swparadas.RecuperarLineasPorCodigoEmpresaResult
	json.Unmarshal([]byte(res.RecuperarLineasPorCodigoEmpresaResult), result)

	if result.CodigoEstado != 0 {
		return nil, errors.New(result.MensajeEstado)
	}

	return result.Lineas, nil
	// sreq := NewSOAPRequest(newRecuperarLineasPorCodigoEmpresaBody("WEB.SUR", "PAR.SW.SUR", CodigoEmpresa, false))

	// xmlReq, err := xml.Marshal(sreq)
	// if err != nil {
	// 	return nil, err
	// }

	// res, err := s.client.Send(xmlReq)
	// if err != nil {
	// 	return nil, err
	// }

	// if res.StatusCode != http.StatusOK {
	// 	return nil, ErrRequestFailed
	// }

	return []*clcitybusapi.Linea{}, nil
}
