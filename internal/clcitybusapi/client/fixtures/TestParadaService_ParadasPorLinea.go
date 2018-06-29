package fixtures

import (
	"encoding/json"
	"testing"

	"github.com/cafrias/clcity/internal/clcitybusapi/mock"
	"github.com/cafrias/clcity/pkg/geo"

	"github.com/cafrias/clcity/internal/clcitybusapi"
	"github.com/cafrias/clcity/internal/clcitybusapi/soapclient/swparadas"
)

// TestParadaServiceParadasPorLinea fixture for test `TestParadaService_ParadasPorLinea`.
func TestParadaServiceParadasPorLinea(t *testing.T) (
	linea *clcitybusapi.Linea,
	fixReq *swparadas.RecuperarParadasPorLineaParaCuandoLlega,
	fixOut []*clcitybusapi.ParadaLinea,
	fixRes *swparadas.RecuperarParadasPorLineaParaCuandoLlegaResponse,
	spy *mock.Spy,
	fixDump []*clcitybusapi.ParadaLinea,
) {
	linea = &clcitybusapi.Linea{
		Codigo:        1529,
		Empresa:       &clcitybusapi.Empresa{Codigo: 355},
		CodigoEntidad: 1234,
	}
	fixReq = &swparadas.RecuperarParadasPorLineaParaCuandoLlega{
		Usuario:           "WEB.SUR",
		Clave:             "PAR.SW.SUR",
		CodigoLineaParada: int32(linea.Codigo),
		IsSubLinea:        false,
		IsInteligente:     false,
	}

	fixOut = []*clcitybusapi.ParadaLinea{
		&clcitybusapi.ParadaLinea{
			Codigo:        57720,
			Identificador: "RG001",
			Nombre:        "Uno y Dos",
			Punto: geo.Point{
				Lat: -53.803239,
				Lon: -67.661785,
			},
			Linea: linea,
		},
		&clcitybusapi.ParadaLinea{
			Codigo:        57721,
			Identificador: "RG002",
			Nombre:        "Tres y Cuatro",
			Punto: geo.Point{
				Lat: -53.803109,
				Lon: -67.662526,
			},
			Linea: linea,
		},
	}

	fixResult := swparadas.RecuperarParadasPorLineaParaCuandoLlegaResult{
		CodigoEstado:  0,
		MensajeEstado: "ok",
		Paradas: []*swparadas.ParadaLinea{
			&swparadas.ParadaLinea{
				Codigo:            "57720",
				Identificador:     "RG001",
				CallePrincipal:    "UNO",
				CalleInterseccion: "DOS",
				Latitud:           "-53,803239",
				Longitud:          "-67,661785",
			},
			&swparadas.ParadaLinea{
				Codigo:            "57721",
				Identificador:     "RG002",
				CallePrincipal:    "TRES",
				CalleInterseccion: "CUATRO",
				Latitud:           "-53,803109",
				Longitud:          "-67,662526",
			},
		},
	}

	resultJSON, err := json.Marshal(fixResult)
	if err != nil {
		t.Fatal("Failed while parsing fixture to JSON.")
	}

	fixRes = &swparadas.RecuperarParadasPorLineaParaCuandoLlegaResponse{
		RecuperarParadasPorLineaParaCuandoLlegaResult: string(resultJSON),
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
			Codigo:        57720,
			Nombre:        "Uno y Dos",
			Identificador: "RG001",
			Punto: geo.Point{
				Lat: -53.803239,
				Lon: -67.661785,
			},
		},
		&clcitybusapi.ParadaLinea{
			Codigo:        57721,
			Nombre:        "Tres y Cuatro",
			Identificador: "RG002",
			Punto: geo.Point{
				Lat: -53.803109,
				Lon: -67.662526,
			},
		},
	}

	return
}
