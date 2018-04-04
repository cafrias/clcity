package fixtures

import (
	"encoding/json"
	"strconv"
	"testing"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/geo"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi"
	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/soapclient/swparadas"
)

// TestParadaServiceParadasPorEmpresa fixture for test `TestParadaService_ParadasPorEmpresa`.
func TestParadaServiceParadasPorEmpresa(t *testing.T) (
	emp *clcitybusapi.Empresa,
	l1str string,
	l2str string,
	fixl []*clcitybusapi.Linea,
	fixpl map[string][]*clcitybusapi.ParadaLinea,
	flinreq *swparadas.RecuperarLineasPorCodigoEmpresa,
	fparreq [2]*swparadas.RecuperarParadasCompletoPorLinea,
	flinresp *swparadas.RecuperarLineasPorCodigoEmpresaResponse,
	fparresp [2]*swparadas.RecuperarParadasCompletoPorLineaResponse,
	fOut []*clcitybusapi.Parada,
) {
	emp = &clcitybusapi.Empresa{
		Codigo: 355,
	}
	fixl = []*clcitybusapi.Linea{
		&clcitybusapi.Linea{
			Codigo:        1529,
			Descripcion:   "RAMAL A",
			CodigoEntidad: 254,
			Empresa:       emp,
		},
		&clcitybusapi.Linea{
			Codigo:        1530,
			Descripcion:   "RAMAL B",
			CodigoEntidad: 254,
			Empresa:       emp,
		},
	}
	l1str = strconv.Itoa(fixl[0].Codigo)
	l2str = strconv.Itoa(fixl[1].Codigo)
	fixpl = map[string][]*clcitybusapi.ParadaLinea{
		l1str: []*clcitybusapi.ParadaLinea{
			&clcitybusapi.ParadaLinea{
				Codigo:                     57720,
				Identificador:              "RG001",
				Descripcion:                "HACIA CHACRA 11",
				AbreviaturaBandera:         "RAMAL A",
				AbreviaturaAmpliadaBandera: "HACIA CHACRA 11",
				Punto: geo.Point{
					Lat:  -53.803239,
					Long: -67.661785,
				},
				AbreviaturaBanderaGIT: "IDA A",
				Linea: fixl[0],
			},
			&clcitybusapi.ParadaLinea{
				Codigo:                     57721,
				Identificador:              "RG002",
				Descripcion:                "HACIA CHACRA 11",
				AbreviaturaBandera:         "RAMAL A",
				AbreviaturaAmpliadaBandera: "HACIA CHACRA 11",
				Punto: geo.Point{
					Lat:  -53.803239,
					Long: -67.661785,
				},
				AbreviaturaBanderaGIT: "IDA A",
				Linea: fixl[0],
			},
		},
		l2str: []*clcitybusapi.ParadaLinea{
			&clcitybusapi.ParadaLinea{
				Codigo:                     57725,
				Identificador:              "RG001",
				Descripcion:                "HACIA CHACRA Mi casa",
				AbreviaturaBandera:         "RAMAL B",
				AbreviaturaAmpliadaBandera: "HACIA CHACRA 11",
				Punto: geo.Point{
					Lat:  -53.803239,
					Long: -67.661785,
				},
				AbreviaturaBanderaGIT: "IDA B",
				Linea: fixl[1],
			},
			&clcitybusapi.ParadaLinea{
				Codigo:                     57731,
				Identificador:              "RG003",
				Descripcion:                "HACIA asd 11",
				AbreviaturaBandera:         "RAMAL B",
				AbreviaturaAmpliadaBandera: "HACIA CHaaACRA 11",
				Punto: geo.Point{
					Lat:  -53.803239,
					Long: -67.661785,
				},
				AbreviaturaBanderaGIT: "IDA B",
				Linea: fixl[1],
			},
		},
	}

	fixpSW := map[string][]*swparadas.ParadaLinea{
		l1str: []*swparadas.ParadaLinea{
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
			&swparadas.ParadaLinea{
				Codigo:                     "57721",
				Identificador:              "RG002",
				Descripcion:                "HACIA CHACRA 11",
				AbreviaturaBandera:         "RAMAL A",
				AbreviaturaAmpliadaBandera: "HACIA CHACRA 11",
				LatitudParada:              "-53,803239",
				LongitudParada:             "-67.661785",
				AbreviaturaBanderaGIT:      "IDA A",
			},
		},
		l2str: []*swparadas.ParadaLinea{
			&swparadas.ParadaLinea{
				Codigo:                     "57725",
				Identificador:              "RG001",
				Descripcion:                "HACIA CHACRA Mi casa",
				AbreviaturaBandera:         "RAMAL B",
				AbreviaturaAmpliadaBandera: "HACIA CHACRA 11",
				LatitudParada:              "-53,803239",
				LongitudParada:             "-67,661785",
				AbreviaturaBanderaGIT:      "IDA B",
			},
			&swparadas.ParadaLinea{
				Codigo:                     "57731",
				Identificador:              "RG003",
				Descripcion:                "HACIA asd 11",
				AbreviaturaBandera:         "RAMAL B",
				AbreviaturaAmpliadaBandera: "HACIA CHaaACRA 11",
				LatitudParada:              "-53,803239",
				LongitudParada:             "-67,661785",
				AbreviaturaBanderaGIT:      "IDA B",
			},
		},
	}

	// Fixture requests
	flinreq = &swparadas.RecuperarLineasPorCodigoEmpresa{
		Usuario:       "WEB.SUR",
		Clave:         "PAR.SW.SUR",
		CodigoEmpresa: 355,
		IsSublinea:    false,
	}
	fparreq = [2]*swparadas.RecuperarParadasCompletoPorLinea{
		&swparadas.RecuperarParadasCompletoPorLinea{
			Usuario:           "WEB.SUR",
			Clave:             "PAR.SW.SUR",
			CodigoLineaParada: int32(fixl[0].Codigo),
			IsSublinea:        false,
			IsInteligente:     false,
		},
		&swparadas.RecuperarParadasCompletoPorLinea{
			Usuario:           "WEB.SUR",
			Clave:             "PAR.SW.SUR",
			CodigoLineaParada: int32(fixl[1].Codigo),
			IsSublinea:        false,
			IsInteligente:     false,
		},
	}

	// Fixture results
	flinresu := &swparadas.RecuperarLineasPorCodigoEmpresaResult{
		CodigoEstado:  0,
		MensajeEstado: "ok",
		Lineas: []*swparadas.Linea{
			&swparadas.Linea{
				CodigoLineaParada: l1str,
				Descripcion:       "RAMAL A",
				CodigoEntidad:     "254",
				CodigoEmpresa:     355,
			},
			&swparadas.Linea{
				CodigoLineaParada: l2str,
				Descripcion:       "RAMAL B",
				CodigoEntidad:     "254",
				CodigoEmpresa:     355,
			},
		},
	}
	flinresuJSON, err := json.Marshal(flinresu)
	if err != nil {
		t.Fatal("Error parsing JSON", err)
	}

	fparresu := [2]*swparadas.RecuperarParadasCompletoPorLineaResult{
		&swparadas.RecuperarParadasCompletoPorLineaResult{
			CodigoEstado:  0,
			MensajeEstado: "ok",
			Paradas:       fixpSW[l1str],
		},
		&swparadas.RecuperarParadasCompletoPorLineaResult{
			CodigoEstado:  0,
			MensajeEstado: "ok",
			Paradas:       fixpSW[l2str],
		},
	}

	var fparresuJSON [2][]byte
	for idx, result := range fparresu {
		resultJSON, err := json.Marshal(result)
		if err != nil {
			t.Fatal("Error parsing JSON", err)
		}
		fparresuJSON[idx] = resultJSON
	}

	// Fixture responses
	flinresp = &swparadas.RecuperarLineasPorCodigoEmpresaResponse{
		RecuperarLineasPorCodigoEmpresaResult: string(flinresuJSON),
	}

	fparresp = [2]*swparadas.RecuperarParadasCompletoPorLineaResponse{
		&swparadas.RecuperarParadasCompletoPorLineaResponse{
			RecuperarParadasCompletoPorLineaResult: string(fparresuJSON[0]),
		},
		&swparadas.RecuperarParadasCompletoPorLineaResponse{
			RecuperarParadasCompletoPorLineaResult: string(fparresuJSON[1]),
		},
	}

	// Fixture output
	fOut = []*clcitybusapi.Parada{
		&clcitybusapi.Parada{
			Codigo: "RG001",
			Punto: geo.Point{
				Lat:  -53.803239,
				Long: -67.661785,
			},
		},
		&clcitybusapi.Parada{
			Codigo: "RG002",
			Punto: geo.Point{
				Lat:  -53.803239,
				Long: -67.661785,
			},
		},
		&clcitybusapi.Parada{
			Codigo: "RG003",
			Punto: geo.Point{
				Lat:  -53.803239,
				Long: -67.661785,
			},
		},
	}

	return
}
