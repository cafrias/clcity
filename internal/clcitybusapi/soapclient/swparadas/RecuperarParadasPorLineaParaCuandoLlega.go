package swparadas

import (
	"encoding/xml"
)

type RecuperarParadasPorLineaParaCuandoLlega struct {
	XMLName xml.Name `xml:"http://clsw.smartmovepro.net/ RecuperarParadasPorLineaParaCuandoLlega"`

	Usuario           string `xml:"usuario,omitempty"`
	Clave             string `xml:"clave,omitempty"`
	CodigoLineaParada int32  `xml:"codigoLineaParada,omitempty"`
	IsSubLinea        bool   `xml:"isSubLinea,omitempty"`
	IsInteligente     bool   `xml:"isInteligente,omitempty"`
}

type RecuperarParadasPorLineaParaCuandoLlegaResponse struct {
	XMLName xml.Name `xml:"http://clsw.smartmovepro.net/ RecuperarParadasPorLineaParaCuandoLlegaResponse"`

	RecuperarParadasPorLineaParaCuandoLlegaResult string `xml:"RecuperarParadasPorLineaParaCuandoLlegaResult,omitempty"`
}

type RecuperarParadasPorLineaParaCuandoLlegaResult struct {
	CodigoEstado  int
	MensajeEstado string
	Paradas       []*ParadaLinea `json:"lineas"`
}

func (service *SWParadasSoap) RecuperarParadasPorLineaParaCuandoLlega(request *RecuperarParadasPorLineaParaCuandoLlega) (*RecuperarParadasPorLineaParaCuandoLlegaResponse, error) {
	response := new(RecuperarParadasPorLineaParaCuandoLlegaResponse)
	err := service.client.Call("http://clsw.smartmovepro.net/RecuperarParadasPorLineaParaCuandoLlega", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// ParadaLinea represents a stop as fetched from 'Cuando Llega City Bus' API.
type ParadaLinea struct {
	Codigo            string
	Identificador     string
	Descripcion       string
	CallePrincipal    string
	CalleInterseccion string
	Latitud           string
	Longitud          string
	Lineas            string
	Inteligente       bool
}
