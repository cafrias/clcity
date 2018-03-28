package swparadas

import (
	"encoding/xml"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi"
)

type RecuperarLineasPorCodigoEmpresa struct {
	XMLName xml.Name `xml:"http://clsw.smartmovepro.net/ RecuperarLineasPorCodigoEmpresa"`

	Usuario       string `xml:"usuario,omitempty"`
	Clave         string `xml:"clave,omitempty"`
	CodigoEmpresa int32  `xml:"codigoEmpresa,omitempty"`
	IsSublinea    bool   `xml:"isSublinea,omitempty"`
}

type RecuperarLineasPorCodigoEmpresaResponse struct {
	XMLName xml.Name `xml:"http://clsw.smartmovepro.net/ RecuperarLineasPorCodigoEmpresaResponse"`

	RecuperarLineasPorCodigoEmpresaResult string `xml:"RecuperarLineasPorCodigoEmpresaResult,omitempty"`
}

type RecuperarLineasPorCodigoEmpresaResult struct {
	CodigoEstado  int
	MensajeEstado string
	Lineas        []*clcitybusapi.Linea `json:"lineas"`
}

func (service *SWParadasSoap) RecuperarLineasPorCodigoEmpresa(request *RecuperarLineasPorCodigoEmpresa) (*RecuperarLineasPorCodigoEmpresaResponse, error) {
	response := new(RecuperarLineasPorCodigoEmpresaResponse)
	err := service.client.Call("http://clsw.smartmovepro.net/RecuperarLineasPorCodigoEmpresa", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
