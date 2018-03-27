package client_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"testing"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi"
)

func TestLineaService_LineasPorEmpresa(t *testing.T) {
	codigos := []string{"1", "2"}

	fixture := map[string][]clcitybusapi.Linea{
		"1": []clcitybusapi.Linea{
			clcitybusapi.Linea{
				CodigoLineaParada: "1529",
				Descripcion:       "RAMAL A",
				CodigoEntidad:     "254",
				CodigoEmpresa:     355,
			},
		},
		"2": []clcitybusapi.Linea{
			clcitybusapi.Linea{
				CodigoLineaParada: "1529",
				Descripcion:       "RAMAL A",
				CodigoEntidad:     "254",
				CodigoEmpresa:     356,
			},
			clcitybusapi.Linea{
				CodigoLineaParada: "1530",
				Descripcion:       "RAMAL B",
				CodigoEntidad:     "254",
				CodigoEmpresa:     356,
			},
		},
	}

	for _, codigo := range codigos {
		resName := fmt.Sprintf("testdata/RecuperarLineasPorCodigoEmpresa/%s/OKResponse.xml", codigo)
		file, err := os.Open(resName)
		if err != nil {
			t.Fatalf("'%v': Failed reading fixture file.\n", codigo)
		}

		httpc := &HTTPClient{
			Res: &http.Response{
				Header:     ResponseHeaders,
				StatusCode: 200,
				Body:       ioutil.NopCloser(file),
			},
		}
		cli := NewClient(httpc)

		result, err := cli.Client.LineaService().LineasPorEmpresa(codigo)
		if err != nil {
			t.Fatalf("'%v': Unexpected error, %v\n", codigo, err)
		}
		// Do gets invoked
		if httpc.DoInvoked == false {
			t.Fatalf("'%v': Didn't call Do function\n", codigo)
		}

		// Check request.
		buf := new(bytes.Buffer)
		buf.ReadFrom(httpc.Req.Body)
		reqBy := buf.Bytes()
		reqName := fmt.Sprintf("testdata/RecuperarLineasPorCodigoEmpresa/%s/OKRequest.xml", codigo)
		fileBy, err := ioutil.ReadFile(reqName)
		if err != nil {
			t.Fatalf("'%v': Error while reading fixture file.", codigo)
		}

		if ok := bytes.Equal(reqBy, fileBy); ok == false {
			t.Fatalf("'%v': Request body isn't correct. Expected: '%s', got: '%s'", codigo, fileBy, reqBy)
		}

		// Check result.
		if len(result) == 0 {
			t.Fatalf("'%v': Received no results.", codigo)
		}
		for key, linea := range fixture[codigo] {
			if ok := reflect.DeepEqual(result[key], &linea); ok == false {
				t.Fatalf("'%v': Unexpected result. Expected '%v', got '%v'", codigo, linea, result[key])
			}
		}
	}

}
