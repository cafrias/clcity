package client_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/client/fixtures"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/client"
	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/mock"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/soapclient/swparadas"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi"
)

func TestParadaService_ParadasPorLinea(t *testing.T) {
	cod := 1529
	fixRequest := &swparadas.RecuperarParadasCompletoPorLinea{
		Usuario:           "WEB.SUR",
		Clave:             "PAR.SW.SUR",
		CodigoLineaParada: int32(cod),
		IsSublinea:        false,
		IsInteligente:     false,
	}

	fixOut := []*clcitybusapi.Parada{
		&clcitybusapi.Parada{
			Codigo:                     "57720",
			Identificador:              "RG001",
			Descripcion:                "HACIA CHACRA 11",
			AbreviaturaBandera:         "RAMAL A",
			AbreviaturaAmpliadaBandera: "HACIA CHACRA 11",
			LatitudParada:              "-53,803239",
			LongitudParada:             "-67,661785",
			AbreviaturaBanderaGIT:      "IDA A",
		},
		&clcitybusapi.Parada{
			Codigo:                     "57721",
			Identificador:              "RG002",
			Descripcion:                "HACIA CHACRA 11",
			AbreviaturaBandera:         "RAMAL A",
			AbreviaturaAmpliadaBandera: "HACIA CHACRA 11",
			LatitudParada:              "-53,803109",
			LongitudParada:             "-67,662526",
			AbreviaturaBanderaGIT:      "IDA A",
		},
	}

	fixResult := swparadas.RecuperarParadasCompletoPorLineaResult{
		CodigoEstado:  0,
		MensajeEstado: "ok",
		Paradas:       fixOut,
	}

	resultJSON, err := json.Marshal(fixResult)
	if err != nil {
		t.Fatal("Failed while parsing fixture to JSON.")
	}

	fixResponse := &swparadas.RecuperarParadasCompletoPorLineaResponse{
		RecuperarParadasCompletoPorLineaResult: string(resultJSON),
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
	scli.RecuperarParadasCompletoPorLineaSpy = spy

	cli := client.NewClient(scli)

	out, err := cli.ParadaService().ParadasPorLinea(cod)
	if err != nil {
		t.Fatalf("Unexpected error: '%v'", err)
	}

	// Called call
	spy = scli.RecuperarParadasCompletoPorLineaSpy
	if spy.Invoked == false {
		t.Fatal("Didn't invoke Call")
	}

	// Called with correct input
	arg, _ := spy.Args[0][0].(*swparadas.RecuperarParadasCompletoPorLinea)
	if ok := reflect.DeepEqual(arg, fixRequest); ok == false {
		t.Fatalf("Didn't call with right request. Expected '%+v', got '%+v'.\n", fixRequest, arg)
	}

	// out valid
	if ok := reflect.DeepEqual(fixOut, out); ok == false {
		t.Fatalf("Didn't receive right output. Expected '%#v', got '%#v'\n", fixOut, out)
	}
}

func TestParadaService_ParadasPorEmpresa(t *testing.T) {
	_, _, _, _, _, _, _, _, flinresp, fparresp, fOut := fixtures.TestParadaServiceParadasPorEmpresa(t)

	scli := NewSOAPClient("", false, nil)

	// Setup spies
	scli.RecuperarLineasPorCodigoEmpresaSpy = &mock.Spy{
		Ret: [][]interface{}{
			[]interface{}{
				flinresp,
				nil,
			},
		},
	}
	scli.RecuperarParadasCompletoPorLineaSpy = &mock.Spy{
		Ret: [][]interface{}{
			[]interface{}{
				fparresp[0],
				nil,
			},
			[]interface{}{
				fparresp[1],
				nil,
			},
		},
	}

	cli := client.NewClient(scli)

	out, err := cli.ParadaService().ParadasPorEmpresa(355)
	if err != nil {
		t.Fatalf("Unexpected error: '%v'\n", err)
	}

	// out valid
	if len(fOut) != len(out) {
		t.Fatal("Didn't return the expected number of elements")
	}

	for _, fvalue := range fOut {
		var found bool
		for _, value := range out {
			if value.Identificador == fvalue.Identificador {
				if ok := reflect.DeepEqual(fvalue, value); ok == true {
					found = true
				}
			}
		}

		if found == false {
			t.Fatalf("Couldn't find '%#v' among the results\n", fvalue)
		}
	}
}
