package client_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/cafrias/clcity/internal/clcitybusapi"
	"github.com/cafrias/clcity/internal/clcitybusapi/client"
	"github.com/cafrias/clcity/internal/clcitybusapi/mock"
	"github.com/cafrias/clcity/internal/clcitybusapi/soapclient/swparadas"
	"github.com/cafrias/clcity/pkg/geo"
)

func TestRecorridoService_Recorrido(t *testing.T) {
	CreateDump()
	defer ClearDump()

	fixLinea := &clcitybusapi.Linea{
		Codigo:        1529,
		CodigoEntidad: 355,
		Empresa: &clcitybusapi.Empresa{
			Codigo: 355,
		},
		Descripcion: "RAMAL A",
	}
	fixRequest := &swparadas.RecuperarRecorridoParaMapaPorEntidadYLinea{
		Usuario:           "WEB.SUR",
		Clave:             "PAR.SW.SUR",
		CodigoLineaParada: int32(fixLinea.Codigo),
		IsSublinea:        false,
	}

	fixOut := &clcitybusapi.Recorrido{
		Puntos: []geo.Point{
			geo.Point{
				Lat: -53,
				Lon: -67,
			},
			geo.Point{
				Lat: -54,
				Lon: -68,
			},
		},
	}

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
