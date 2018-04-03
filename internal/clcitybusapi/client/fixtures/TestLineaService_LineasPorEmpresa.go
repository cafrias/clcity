package fixtures

import (
	"encoding/json"
	"testing"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi"
	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/mock"
	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/soapclient/swparadas"
)

// TestLineaServiceLineasPorEmpresa fixture for test `TestLineaService_LineasPorEmpresa`.
func TestLineaServiceLineasPorEmpresa(t *testing.T) (
	spy *mock.Spy,
	fixReq *swparadas.RecuperarLineasPorCodigoEmpresa,
	fixOut []*clcitybusapi.Linea,
	fixRes *swparadas.RecuperarLineasPorCodigoEmpresaResponse,
) {
	fixReq = &swparadas.RecuperarLineasPorCodigoEmpresa{
		Usuario:       "WEB.SUR",
		Clave:         "PAR.SW.SUR",
		CodigoEmpresa: 355,
		IsSublinea:    false,
	}

	fixOut = []*clcitybusapi.Linea{
		&clcitybusapi.Linea{
			Codigo:        1529,
			Descripcion:   "RAMAL A",
			CodigoEntidad: 254,
			CodigoEmpresa: 356,
		},
	}

	fixResult := swparadas.RecuperarLineasPorCodigoEmpresaResult{
		CodigoEstado:  0,
		MensajeEstado: "ok",
		Lineas: []*swparadas.Linea{
			&swparadas.Linea{
				CodigoLineaParada: "1529",
				Descripcion:       "RAMAL A",
				CodigoEntidad:     "254",
				CodigoEmpresa:     356,
			},
		},
	}

	resultJSON, err := json.Marshal(fixResult)
	if err != nil {
		t.Fatal("Failed while parsing fixture to JSON.")
	}

	fixRes = &swparadas.RecuperarLineasPorCodigoEmpresaResponse{
		RecuperarLineasPorCodigoEmpresaResult: string(resultJSON),
	}

	spy = &mock.Spy{
		Ret: [][]interface{}{
			[]interface{}{
				fixRes,
				nil,
			},
		},
	}

	return
}
