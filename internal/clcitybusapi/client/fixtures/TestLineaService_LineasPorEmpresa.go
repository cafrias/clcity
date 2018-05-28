package fixtures

import (
	"encoding/json"
	"testing"

	"github.com/friasdesign/clcity/internal/clcitybusapi/client"
	"github.com/friasdesign/clcity/pkg/geo"

	"github.com/friasdesign/clcity/internal/clcitybusapi"
	"github.com/friasdesign/clcity/internal/clcitybusapi/mock"
	"github.com/friasdesign/clcity/internal/clcitybusapi/soapclient/swparadas"
)

// TestLineaServiceLineasPorEmpresa fixture for test `TestLineaService_LineasPorEmpresa`.
func TestLineaServiceLineasPorEmpresa(t *testing.T) (
	fEmp *clcitybusapi.Empresa,
	// Recorrido
	fRec *clcitybusapi.Recorrido,
	sRec *mock.Spy,
	// Paradas
	fPar []*clcitybusapi.ParadaLinea,
	sPar *mock.Spy,
	// Lineas
	fLin []*clcitybusapi.Linea,
	fLinSW []*swparadas.Linea,
	fLinReq *swparadas.RecuperarLineasPorCodigoEmpresa,
	fLinResp *swparadas.RecuperarLineasPorCodigoEmpresaResponse,
	sLin *mock.Spy,
	fLinDump []*clcitybusapi.Linea,
) {
	fEmp = &clcitybusapi.Empresa{
		Codigo: 355,
	}

	user := "WEB.SUR"
	pass := "PAR.SW.SUR"

	// Recorrido
	fRec = &clcitybusapi.Recorrido{
		Puntos: []geo.Point{
			geo.Point{
				Lat: -20,
				Lon: 20,
			},
			geo.Point{
				Lat: -21,
				Lon: 21,
			},
		},
	}
	sRec = &mock.Spy{
		Ret: [][]interface{}{
			[]interface{}{
				fRec,
				nil,
			},
		},
	}

	// Paradas
	fPar = []*clcitybusapi.ParadaLinea{
		&clcitybusapi.ParadaLinea{
			Codigo:        123456,
			Descripcion:   "Alguna",
			Identificador: "RG001",
			Punto: geo.Point{
				Lat: 20,
				Lon: 20,
			},
		},
		&clcitybusapi.ParadaLinea{
			Codigo:        123458,
			Descripcion:   "Alguna",
			Identificador: "RG002",
			Punto: geo.Point{
				Lat: 21,
				Lon: 21,
			},
		},
	}
	sPar = &mock.Spy{
		Ret: [][]interface{}{
			[]interface{}{
				fPar,
				nil,
			},
		},
	}

	// Lineas
	fLin = []*clcitybusapi.Linea{
		&clcitybusapi.Linea{
			Codigo:        1529,
			Descripcion:   "RAMAL A",
			CodigoEntidad: 254,
			Color:         client.ColorMap[1529],
			Empresa:       fEmp,
			Paradas:       fPar,
			Recorrido:     fRec,
		},
	}
	fLinSW = []*swparadas.Linea{
		&swparadas.Linea{
			CodigoLineaParada: "1529",
			Descripcion:       "RAMAL A",
			CodigoEntidad:     "254",
			CodigoEmpresa:     356,
		},
	}

	fLinReq = &swparadas.RecuperarLineasPorCodigoEmpresa{
		Usuario:       user,
		Clave:         pass,
		CodigoEmpresa: 355,
		IsSublinea:    false,
	}

	fLinRes := swparadas.RecuperarLineasPorCodigoEmpresaResult{
		CodigoEstado:  0,
		MensajeEstado: "ok",
		Lineas:        fLinSW,
	}
	resultJSON, _ := json.Marshal(fLinRes)
	fLinResp = &swparadas.RecuperarLineasPorCodigoEmpresaResponse{
		RecuperarLineasPorCodigoEmpresaResult: string(resultJSON),
	}

	sLin = &mock.Spy{
		Ret: [][]interface{}{
			[]interface{}{
				fLinResp,
				nil,
			},
		},
	}

	fLinDump = []*clcitybusapi.Linea{
		&clcitybusapi.Linea{
			Codigo:        1529,
			Descripcion:   "RAMAL A",
			CodigoEntidad: 254,
			Color:         client.ColorMap[1529],
			Recorrido:     fRec,
		},
	}

	return
}
