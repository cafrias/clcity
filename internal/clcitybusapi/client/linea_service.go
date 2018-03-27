package client

import (
	"encoding/xml"
	"net/http"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi"
)

var _ clcitybusapi.LineaService = &LineaService{}

// LineaService has actions to fetch 'Parada' data from Cuando Llega City Bus API.
type LineaService struct {
	client *SOAPClient
}

type RecuperarLineasPorCodigoEmpresaRequest struct {
	XMLName xml.Name `xml:"RecuperarLineasPorCodigoEmpresa"`
	XMLns   string   `xml:"xmlns,attr"`
	ID      string   `xml:"id,attr"`
	Root    string   `xml:"c:root,attr"`
	Credenciales
	CodigoEmpresa string `xml:"codigoEmpresa"`
	Sublinea      bool   `xml:"isSublinea"`
}

type RecuperarLineasPorCodigoEmpresaResponse struct {
	XMLName xml.Name `xml:"RecuperarLineasPorCodigoEmpresa"`
	XMLns   string   `xml:"xmlns,attr"`
	Result  RecuperarLineasPorCodigoEmpresaResult
}

type RecuperarLineasPorCodigoEmpresaResult struct {
	XMLName xml.Name `xml:"RecuperarLineasPorCodigoEmpresaResult"`
}

// Credenciales represents a API credentials required to make HTTP calls.
type Credenciales struct {
	Usuario Usuario
	Clave   Clave
}

// Usuario represents a user for API credentials.
type Usuario struct {
	XMLName xml.Name `xml:"usuario"`
	Type    string   `xml:"i:type,attr"`
	Value   string   `xml:",chardata"`
}

// Clave represents the password for a user in API credentials.
type Clave struct {
	XMLName xml.Name `xml:"clave"`
	Type    string   `xml:"i:type,attr"`
	Value   string   `xml:",chardata"`
}

func newRecuperarLineasPorCodigoEmpresaRequest(usuario string, clave string, codigoEmpresa string, sublinea bool) RecuperarLineasPorCodigoEmpresaRequest {
	return RecuperarLineasPorCodigoEmpresaRequest{
		XMLns: "http://clsw.smartmovepro.net/",
		ID:    "o0",
		Root:  "1",
		Credenciales: Credenciales{
			Usuario: Usuario{
				Type:  "d:string",
				Value: usuario,
			},
			Clave: Clave{
				Type:  "d:string",
				Value: clave,
			},
		},
		CodigoEmpresa: codigoEmpresa,
		Sublinea:      sublinea,
	}
}

// LineasPorEmpresa fetches all 'Parada' entities associated with a given 'Linea' identified by the code passed as `CodigoLineaParada`.
func (s *LineaService) LineasPorEmpresa(CodigoEmpresa string) ([]*clcitybusapi.Linea, error) {
	sreq := NewSOAPRequest(newRecuperarLineasPorCodigoEmpresaBody("WEB.SUR", "PAR.SW.SUR", CodigoEmpresa, false))

	xmlReq, err := xml.Marshal(sreq)
	if err != nil {
		return nil, err
	}

	res, err := s.client.Send(xmlReq)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, ErrRequestFailed
	}

	return []*clcitybusapi.Linea{}, nil
}
