package swparadas

import (
	"encoding/xml"
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
	Paradas       []*Parada `json:"paradas"`
}

func (service *SWParadasSoap) RecuperarParadasCompletoPorLinea(request *RecuperarParadasCompletoPorLinea) (*RecuperarParadasCompletoPorLineaResponse, error) {
	response := new(RecuperarParadasCompletoPorLineaResponse)
	err := service.client.Call("http://clsw.smartmovepro.net/RecuperarParadasCompletoPorLinea", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Parada represents a stop as fetched from 'Cuando Llega City Bus' API.
type Parada struct {
	Codigo                     string
	Identificador              string
	Descripcion                string
	AbreviaturaBandera         string
	AbreviaturaAmpliadaBandera string
	LatitudParada              string
	LongitudParada             string
	AbreviaturaBanderaGIT      string
}
