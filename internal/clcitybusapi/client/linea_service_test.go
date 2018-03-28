package client_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/client"
	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/mock"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/soapclient/swparadas"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi"
)

func TestLineaService_LineasPorEmpresa(t *testing.T) {
	fixRequest := &swparadas.RecuperarLineasPorCodigoEmpresa{
		Usuario:       "WEB.SUR",
		Clave:         "PAR.SW.SUR",
		CodigoEmpresa: 355,
		IsSublinea:    false,
	}

	fixOut := []*clcitybusapi.Linea{
		&clcitybusapi.Linea{
			CodigoLineaParada: "1529",
			Descripcion:       "RAMAL A",
			CodigoEntidad:     "254",
			CodigoEmpresa:     356,
		},
	}

	fixResult := swparadas.RecuperarLineasPorCodigoEmpresaResult{
		CodigoEstado:  0,
		MensajeEstado: "ok",
		Lineas:        fixOut,
	}

	resultJSON, err := json.Marshal(fixResult)
	if err != nil {
		t.Fatal("Failed while parsing fixture to JSON.")
	}

	fixResponse := &swparadas.RecuperarLineasPorCodigoEmpresaResponse{
		RecuperarLineasPorCodigoEmpresaResult: string(resultJSON),
	}

	spy := &mock.Spy{
		Ret: []interface{}{
			fixResponse,
			nil,
		},
	}

	scli := NewSOAPClient("", false, nil)
	scli.RecuperarLineasPorCodigoEmpresaSpy = spy

	cli := client.NewClient()
	cli.Connect(scli)

	out, err := cli.LineaService().LineasPorEmpresa(355)
	if err != nil {
		t.Fatalf("Unexpected error: '%v'", err)
	}

	// Called call
	spy = scli.RecuperarLineasPorCodigoEmpresaSpy
	if spy.Invoked == false {
		t.Fatal("Didn't invoke Call")
	}

	// Called with correct input
	arg, _ := spy.Args[0].(*swparadas.RecuperarLineasPorCodigoEmpresa)
	if ok := reflect.DeepEqual(arg, fixRequest); ok == false {
		t.Fatalf("Didn't call with right request. Expected '%+v', got '%+v'.\n", fixRequest, arg)
	}

	// out valid
	if ok := reflect.DeepEqual(fixOut, out); ok == false {
		t.Fatalf("Didn't receive right output. Expected '%#v', got '%#v'\n", fixOut, out)
	}
}

func TestLineaService_LineasPorEmpresa_ErrNotConnected(t *testing.T) {
	cli := client.NewClient()
	_, err := cli.LineaService().LineasPorEmpresa(355)
	if err != client.ErrNotConnected {
		t.Fatalf("Received un expected error, '%v'", err)
	}
}
