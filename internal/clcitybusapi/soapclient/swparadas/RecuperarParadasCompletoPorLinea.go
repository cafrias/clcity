package swparadas

import (
	"encoding/xml"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi"
)

type RecuperarParadasCompletoPorLinea struct {
	XMLName xml.Name `xml:"http://clsw.smartmovepro.net/ RecuperarParadasCompletoPorLinea"`

	Usuario           string `xml:"usuario,omitempty"`
	Clave             string `xml:"clave,omitempty"`
	CodigoLineaParada int32  `xml:"codigoLineaParada,omitempty"`
	IsSublinea        bool   `xml:"isSublinea,omitempty"`
	IsInteligente     bool   `xml:"isInteligente,omitempty"`
}

type RecuperarParadasCompletoPorLineaResponse struct {
	XMLName xml.Name `xml:"http://clsw.smartmovepro.net/ RecuperarParadasCompletoPorLineaResponse"`

	RecuperarParadasCompletoPorLineaResult string `xml:"RecuperarParadasCompletoPorLineaResult,omitempty"`
}

type RecuperarParadasCompletoPorLineaResult struct {
	CodigoEstado  int
	MensajeEstado string
	Paradas       []*clcitybusapi.Parada `json:"paradas"`
}

func (service *SWParadasSoap) RecuperarParadasCompletoPorLinea(request *RecuperarParadasCompletoPorLinea) (*RecuperarParadasCompletoPorLineaResponse, error) {
	response := new(RecuperarParadasCompletoPorLineaResponse)
	err := service.client.Call("http://clsw.smartmovepro.net/RecuperarParadasCompletoPorLinea", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
