package client_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/geo"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/client"
	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/mock"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/soapclient/swparadas"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi"
)

func TestRecorridoService_Recorrido(t *testing.T) {
	CreateDump()
	defer ClearDump()

	fixLinea := &clcitybusapi.Linea{
		Codigo:        1529,
		CodigoEntidad: 355,
		CodigoEmpresa: 355,
		Descripcion:   "RAMAL A",
	}
	fixRequest := &swparadas.RecuperarRecorridoParaMapaPorEntidadYLinea{
		Usuario:           "WEB.SUR",
		Clave:             "PAR.SW.SUR",
		CodigoLineaParada: int32(fixLinea.Codigo),
		IsSublinea:        false,
	}

	fixOut := clcitybusapi.NewRecorrido(fixLinea, []*geo.Point{
		&geo.Point{
			Lat:  -53,
			Long: -67,
		},
		&geo.Point{
			Lat:  -54,
			Long: -68,
		},
	})

	fixResult := swparadas.RecuperarRecorridoParaMapaPorEntidadYLineaResult{
		CodigoEstado:  0,
		MensajeEstado: "ok",
		Puntos: []*swparadas.Recorrido{
			&swparadas.Recorrido{
				Descripcion: "1; ..HACIA HACIA CHACRA 11",
				Latitud:     -53,
				Longitud:    -67,
			},
			&swparadas.Recorrido{
				Descripcion: "1; ..HACIA HACIA CHACRA 11",
				Latitud:     -54,
				Longitud:    -68,
			},
		},
	}

	resultJSON, err := json.Marshal(fixResult)
	if err != nil {
		t.Fatal("Failed while parsing fixture to JSON.")
	}

	fixResponse := &swparadas.RecuperarRecorridoParaMapaPorEntidadYLineaResponse{
		RecuperarRecorridoParaMapaPorEntidadYLineaResult: string(resultJSON),
	}

	spy := &mock.Spy{
		Ret: [][]interface{}{
			[]interface{}{
				fixResponse,
				nil,
			},
		},
	}

	scli := NewSOAPClient("", false, nil)
	scli.RecuperarRecorridoParaMapaPorEntidadYLineaSpy = spy

	cli := client.NewClient(scli, "testdata")

	out, err := cli.RecorridoService().RecorridoDeLinea(fixLinea)
	if err != nil {
		t.Fatalf("Unexpected error: '%v'", err)
	}

	// Called call
	spy = scli.RecuperarRecorridoParaMapaPorEntidadYLineaSpy
	if spy.Invoked == false {
		t.Fatal("Didn't invoke Call")
	}

	// Called with correct input
	arg, _ := spy.Args[0][0].(*swparadas.RecuperarRecorridoParaMapaPorEntidadYLinea)
	if ok := reflect.DeepEqual(arg, fixRequest); ok == false {
		t.Fatalf("Didn't call with right request. Expected '%+v', got '%+v'.\n", fixRequest, arg)
	}

	// out valid
	if ok := reflect.DeepEqual(fixOut, out); ok == false {
		t.Fatalf("Didn't receive right output. Expected '%#v', got '%#v'\n", fixOut, out)
	}
}

// func TestRecorridoService_ParadasPorEmpresa(t *testing.T) {
// 	CreateDump()
// 	defer ClearDump()

// 	_, _, _, _, _, _, _, _, flinresp, fparresp, fOut := fixtures.TestRecorridoServiceParadasPorEmpresa(t)

// 	scli := NewSOAPClient("", false, nil)

// 	// Setup spies
// 	scli.RecuperarLineasPorCodigoEmpresaSpy = &mock.Spy{
// 		Ret: [][]interface{}{
// 			[]interface{}{
// 				flinresp,
// 				nil,
// 			},
// 		},
// 	}
// 	scli.RecuperarParadasCompletoPorLineaSpy = &mock.Spy{
// 		Ret: [][]interface{}{
// 			[]interface{}{
// 				fparresp[0],
// 				nil,
// 			},
// 			[]interface{}{
// 				fparresp[1],
// 				nil,
// 			},
// 		},
// 	}

// 	cli := client.NewClient(scli, DumpPath)

// 	out, err := cli.ParadaService().ParadasPorEmpresa(355)
// 	if err != nil {
// 		t.Fatalf("Unexpected error: '%v'\n", err)
// 	}

// 	// out valid
// 	if len(fOut) != len(out) {
// 		t.Fatal("Didn't return the expected number of elements")
// 	}

// 	for _, fvalue := range fOut {
// 		var found bool
// 		for _, value := range out {
// 			if value.Identificador == fvalue.Identificador {
// 				if ok := reflect.DeepEqual(fvalue, value); ok == true {
// 					found = true
// 				}
// 			}
// 		}

// 		if found == false {
// 			t.Fatalf("Couldn't find '%#v' among the results\n", fvalue)
// 		}
// 	}
// }
