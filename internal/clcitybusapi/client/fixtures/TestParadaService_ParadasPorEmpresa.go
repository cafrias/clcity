package fixtures

import (
	"encoding/json"
	"testing"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/geo"
	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/soapclient/swparadas"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/mock"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi"
)

// TestParadaServiceParadasPorEmpresa fixture for test `TestParadaService_ParadasPorEmpresa`.
func TestParadaServiceParadasPorEmpresa(t *testing.T) (
	fEmp *clcitybusapi.Empresa,
	l1str string,
	l2str string,
	// Lineas
	fLin map[string]*clcitybusapi.Linea,
	sLin *mock.Spy,
	// ParadaLineas
	// fParLin map[string][]*clcitybusapi.ParadaLinea,
	fParLinSW map[string][]*swparadas.ParadaLinea,
	fParLinReq map[string]*swparadas.RecuperarParadasCompletoPorLinea,
	fParLinResp map[string]*swparadas.RecuperarParadasCompletoPorLineaResponse,
	sParLin *mock.Spy,
	// Paradas
	fPar []*clcitybusapi.Parada,
) {
	fEmp = &clcitybusapi.Empresa{
		Codigo: 355,
	}
	l1str = "1529"
	l2str = "1530"

	// Lineas
	fLin = map[string]*clcitybusapi.Linea{
		l1str: &clcitybusapi.Linea{
			Codigo:        1529,
			Descripcion:   "RAMAL A",
			CodigoEntidad: 254,
			Empresa:       fEmp,
		},
		l2str: &clcitybusapi.Linea{
			Codigo:        1530,
			Descripcion:   "RAMAL B",
			CodigoEntidad: 254,
			Empresa:       fEmp,
		},
	}
	var LineasPorEmpresaRet []*clcitybusapi.Linea
	for _, lin := range fLin {
		LineasPorEmpresaRet = append(LineasPorEmpresaRet, lin)
	}
	sLin = &mock.Spy{
		Ret: [][]interface{}{
			[]interface{}{
				LineasPorEmpresaRet,
				nil,
			},
		},
	}

	// ParadaLineas
	fParLinSW = map[string][]*swparadas.ParadaLinea{
		l1str: []*swparadas.ParadaLinea{
			&swparadas.ParadaLinea{
				Codigo:         "123456",
				Identificador:  "RG001",
				LatitudParada:  "20,1",
				LongitudParada: "20,1",
			},
			&swparadas.ParadaLinea{
				Codigo:         "123457",
				Identificador:  "RG002",
				LatitudParada:  "21,1",
				LongitudParada: "21,1",
			},
		},
		l2str: []*swparadas.ParadaLinea{
			&swparadas.ParadaLinea{
				Codigo:         "123458",
				Identificador:  "RG001",
				LatitudParada:  "20,1",
				LongitudParada: "20,1",
			},
			&swparadas.ParadaLinea{
				Codigo:         "123412",
				Identificador:  "RG003",
				LatitudParada:  "22,1",
				LongitudParada: "22,1",
			},
		},
	}
	user := "WEB.SUR"
	pass := "PAR.SW.SUR"
	fParLinReq = map[string]*swparadas.RecuperarParadasCompletoPorLinea{
		l1str: &swparadas.RecuperarParadasCompletoPorLinea{
			Usuario:           user,
			Clave:             pass,
			CodigoLineaParada: int32(fLin[l1str].Codigo),
			IsSublinea:        false,
			IsInteligente:     false,
		},
		l2str: &swparadas.RecuperarParadasCompletoPorLinea{
			Usuario:           user,
			Clave:             pass,
			CodigoLineaParada: int32(fLin[l2str].Codigo),
			IsSublinea:        false,
			IsInteligente:     false,
		},
	}
	fParLinRes := map[string]*swparadas.RecuperarParadasCompletoPorLineaResult{
		l1str: &swparadas.RecuperarParadasCompletoPorLineaResult{
			CodigoEstado:  0,
			MensajeEstado: "ok",
			Paradas:       fParLinSW[l1str],
		},
		l2str: &swparadas.RecuperarParadasCompletoPorLineaResult{
			CodigoEstado:  0,
			MensajeEstado: "ok",
			Paradas:       fParLinSW[l2str],
		},
	}
	fParLinResStr := make(map[string]string)
	for lin, res := range fParLinRes {
		by, _ := json.Marshal(res)
		fParLinResStr[lin] = string(by)
	}
	fParLinResp = map[string]*swparadas.RecuperarParadasCompletoPorLineaResponse{
		l1str: &swparadas.RecuperarParadasCompletoPorLineaResponse{
			RecuperarParadasCompletoPorLineaResult: fParLinResStr[l1str],
		},
		l2str: &swparadas.RecuperarParadasCompletoPorLineaResponse{
			RecuperarParadasCompletoPorLineaResult: fParLinResStr[l2str],
		},
	}
	sParLin = &mock.Spy{
		Ret: [][]interface{}{
			[]interface{}{
				fParLinResp[l1str],
				nil,
			},
			[]interface{}{
				fParLinResp[l2str],
				nil,
			},
		},
	}

	// Parada
	fPar = []*clcitybusapi.Parada{
		&clcitybusapi.Parada{
			Codigo: "RG001",
			Punto: geo.Point{
				Lat:  20.1,
				Long: 20.1,
			},
		},
		&clcitybusapi.Parada{
			Codigo: "RG002",
			Punto: geo.Point{
				Lat:  21.1,
				Long: 21.1,
			},
		},
		&clcitybusapi.Parada{
			Codigo: "RG003",
			Punto: geo.Point{
				Lat:  22.1,
				Long: 22.1,
			},
		},
	}

	return
}
