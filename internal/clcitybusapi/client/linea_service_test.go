package client_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/client"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/soapclient/swparadas"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi"
)

func TestLineaService_LineasPorEmpresa(t *testing.T) {
	expectedRequest := swparadas.RecuperarLineasPorCodigoEmpresa{
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

	fixResponse := swparadas.RecuperarLineaPorCuandoLlegaResponse{
		RecuperarLineaPorCuandoLlegaResult: string(resultJSON),
	}

	scli, swcli := NewSWParadasSoap("", false, nil)

	cli := client.NewClient(swcli)
	scli.CallRes = fixResponse

	out, err := cli.LineaService().LineasPorEmpresa(355)
	if err != nil {
		t.Fatalf("Unexpected error: '%v'", err)
	}

	// Called call
	if scli.CallInvoked == false {
		t.Fatal("Didn't invoke Call")
	}

	// Called with correct input
	if ok := reflect.DeepEqual(scli.CallReq, &expectedRequest); ok == false {
		t.Fatalf("Didn't call with right request. Expected '%v', got '%v'.\n", expectedRequest, scli.CallReq)
	}

	// out valid
	if ok := reflect.DeepEqual(fixOut, out); ok == false {
		t.Fatalf("Didn't receive right output. Expected '%v', got '%v'", fixOut, out)
	}

	// for _, codigo := range codigos {
	// 	resName := fmt.Sprintf("testdata/RecuperarLineasPorCodigoEmpresa/%s/OKResponse.xml", codigo)
	// 	file, err := os.Open(resName)
	// 	if err != nil {
	// 		t.Fatalf("'%v': Failed reading fixture file.\n", codigo)
	// 	}

	// 	httpc := &HTTPClient{
	// 		Res: &http.Response{
	// 			Header:     ResponseHeaders,
	// 			StatusCode: 200,
	// 			Body:       ioutil.NopCloser(file),
	// 		},
	// 	}
	// 	cli := NewClient(httpc)

	// 	result, err := cli.Client.LineaService().LineasPorEmpresa(codigo)
	// 	if err != nil {
	// 		t.Fatalf("'%v': Unexpected error, %v\n", codigo, err)
	// 	}
	// 	// Do gets invoked
	// 	if httpc.DoInvoked == false {
	// 		t.Fatalf("'%v': Didn't call Do function\n", codigo)
	// 	}

	// 	// Check request.
	// 	buf := new(bytes.Buffer)
	// 	buf.ReadFrom(httpc.Req.Body)
	// 	reqBy := buf.Bytes()
	// 	reqName := fmt.Sprintf("testdata/RecuperarLineasPorCodigoEmpresa/%s/OKRequest.xml", codigo)
	// 	fileBy, err := ioutil.ReadFile(reqName)
	// 	if err != nil {
	// 		t.Fatalf("'%v': Error while reading fixture file.", codigo)
	// 	}

	// 	if ok := bytes.Equal(reqBy, fileBy); ok == false {
	// 		t.Fatalf("'%v': Request body isn't correct. Expected: '%s', got: '%s'", codigo, fileBy, reqBy)
	// 	}

	// 	// Check result.
	// 	if len(result) == 0 {
	// 		t.Fatalf("'%v': Received no results.", codigo)
	// 	}
	// 	for key, linea := range fixture[codigo] {
	// 		if ok := reflect.DeepEqual(result[key], &linea); ok == false {
	// 			t.Fatalf("'%v': Unexpected result. Expected '%v', got '%v'", codigo, linea, result[key])
	// 		}
	// 	}
	// }

}
