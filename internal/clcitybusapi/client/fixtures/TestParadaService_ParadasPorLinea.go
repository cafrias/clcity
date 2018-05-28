package fixtures

import (
	"encoding/json"
	"testing"

	"github.com/friasdesign/clcity/internal/clcitybusapi/mock"
	"github.com/friasdesign/clcity/pkg/geo"

	"github.com/friasdesign/clcity/internal/clcitybusapi"
	"github.com/friasdesign/clcity/internal/clcitybusapi/soapclient/swparadas"
)

// TestParadaServiceParadasPorLinea fixture for test `TestParadaService_ParadasPorLinea`.
func TestParadaServiceParadasPorLinea(t *testing.T) (
	linea *clcitybusapi.Linea,
	fixReq *swparadas.RecuperarParadasCompletoPorLinea,
	fixOut []*clcitybusapi.ParadaLinea,
	fixRes *swparadas.RecuperarParadasCompletoPorLineaResponse,
	spy *mock.Spy,
	fixDump []*clcitybusapi.ParadaLinea,
) {
	linea = &clcitybusapi.Linea{
		Codigo:        1529,
		Empresa:       &clcitybusapi.Empresa{Codigo: 355},
		CodigoEntidad: 1234,
	}
	fixReq = &swparadas.RecuperarParadasCompletoPorLinea{
		Usuario:           "WEB.SUR",
		Clave:             "PAR.SW.SUR",
		CodigoLineaParada: int32(linea.Codigo),
		IsSublinea:        false,
		IsInteligente:     false,
	}

	fixOut = []*clcitybusapi.ParadaLinea{
		&clcitybusapi.ParadaLinea{
			Codigo:                     57720,
			Identificador:              "RG001",
			Descripcion:                "HACIA CHACRA 11",
			AbreviaturaBandera:         "RAMAL A",
			AbreviaturaAmpliadaBandera: "HACIA CHACRA 11",
			Punto: geo.Point{
				Lat: -53.803239,
				Lon: -67.661785,
			},
			AbreviaturaBanderaGIT: "IDA A",
			Linea: linea,
		},
		&clcitybusapi.ParadaLinea{
			Codigo:                     57721,
			Identificador:              "RG002",
			Descripcion:                "HACIA CHACRA 11",
			AbreviaturaBandera:         "RAMAL A",
			AbreviaturaAmpliadaBandera: "HACIA CHACRA 11",
			Punto: geo.Point{
				Lat: -53.803109,
				Lon: -67.662526,
			},
			AbreviaturaBanderaGIT: "IDA A",
			Linea: linea,
		},
	}

	fixResult := swparadas.RecuperarParadasCompletoPorLineaResult{
		CodigoEstado:  0,
		MensajeEstado: "ok",
		Paradas: map[string][]*swparadas.ParadaLinea{
			"IDA": {
				&swparadas.ParadaLinea{
					Codigo:                     "57720",
					Identificador:              "RG001",
					Descripcion:                "HACIA CHACRA 11",
					AbreviaturaBandera:         "RAMAL A",
					AbreviaturaAmpliadaBandera: "HACIA CHACRA 11",
					LatitudParada:              "-53,803239",
					LongitudParada:             "-67,661785",
					AbreviaturaBanderaGIT:      "IDA A",
				},
			},
			"VTA": {
				&swparadas.ParadaLinea{
					Codigo:                     "57721",
					Identificador:              "RG002",
					Descripcion:                "HACIA CHACRA 11",
					AbreviaturaBandera:         "RAMAL A",
					AbreviaturaAmpliadaBandera: "HACIA CHACRA 11",
					LatitudParada:              "-53,803109",
					LongitudParada:             "-67,662526",
					AbreviaturaBanderaGIT:      "IDA A",
				},
			},
		},
	}

	resultJSON, err := json.Marshal(fixResult)
	if err != nil {
		t.Fatal("Failed while parsing fixture to JSON.")
	}

	fixRes = &swparadas.RecuperarParadasCompletoPorLineaResponse{
		RecuperarParadasCompletoPorLineaResult: string(resultJSON),
	}

	spy = &mock.Spy{
		Ret: [][]interface{}{
			[]interface{}{
				fixRes,
				nil,
			},
		},
	}

	fixDump = []*clcitybusapi.ParadaLinea{
		&clcitybusapi.ParadaLinea{
			Codigo:                     57720,
			Identificador:              "RG001",
			Descripcion:                "HACIA CHACRA 11",
			AbreviaturaBandera:         "RAMAL A",
			AbreviaturaAmpliadaBandera: "HACIA CHACRA 11",
			Punto: geo.Point{
				Lat: -53.803239,
				Lon: -67.661785,
			},
			AbreviaturaBanderaGIT: "IDA A",
		},
		&clcitybusapi.ParadaLinea{
			Codigo:                     57721,
			Identificador:              "RG002",
			Descripcion:                "HACIA CHACRA 11",
			AbreviaturaBandera:         "RAMAL A",
			AbreviaturaAmpliadaBandera: "HACIA CHACRA 11",
			Punto: geo.Point{
				Lat: -53.803109,
				Lon: -67.662526,
			},
			AbreviaturaBanderaGIT: "IDA A",
		},
	}

	return
}
