package swparadas

import (
	"encoding/xml"
)

type RecuperarRecorridoParaMapaPorEntidadYLinea struct {
	XMLName xml.Name `xml:"http://clsw.smartmovepro.net/ RecuperarRecorridoParaMapaPorEntidadYLinea"`

	Usuario           string `xml:"usuario,omitempty"`
	Clave             string `xml:"clave,omitempty"`
	CodigoLineaParada int32  `xml:"codigoLineaParada,omitempty"`
	IsSublinea        bool   `xml:"isSublinea,omitempty"`
}

type RecuperarRecorridoParaMapaPorEntidadYLineaResponse struct {
	XMLName xml.Name `xml:"http://clsw.smartmovepro.net/ RecuperarRecorridoParaMapaPorEntidadYLineaResponse"`

	RecuperarRecorridoParaMapaPorEntidadYLineaResult string `xml:"RecuperarRecorridoParaMapaPorEntidadYLineaResult,omitempty"`
}

type RecuperarRecorridoParaMapaPorEntidadYLineaResult struct {
	CodigoEstado  int
	MensajeEstado string
	Puntos        []*Recorrido `json:"puntos"`
}

func (service *SWParadasSoap) RecuperarRecorridoParaMapaPorEntidadYLinea(request *RecuperarRecorridoParaMapaPorEntidadYLinea) (*RecuperarRecorridoParaMapaPorEntidadYLineaResponse, error) {
	response := new(RecuperarRecorridoParaMapaPorEntidadYLineaResponse)
	err := service.client.Call("http://clsw.smartmovepro.net/RecuperarRecorridoParaMapaPorEntidadYLinea", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Recorrido represents a 'Recorrido' as fetched from 'Cuando Llega City Bus' API.
type Recorrido struct {
	Descripcion string
	Latitud     float64
	Longitud    float64
}
