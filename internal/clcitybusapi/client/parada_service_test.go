package client_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/dump"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi"
	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/client/fixtures"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/client"
	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/mock"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/soapclient/swparadas"
)

func TestParadaService_ParadasPorLinea(t *testing.T) {
	CreateDump()
	defer ClearDump()

	linea, fixReq, fixOut, _, spy := fixtures.TestParadaServiceParadasPorLinea(t)

	scli := NewSOAPClient("", false, nil)
	scli.RecuperarParadasCompletoPorLineaSpy = spy

	cli := client.NewClient(scli, DumpPath)

	out, err := cli.ParadaService().ParadasPorLinea(linea)
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
	if ok := reflect.DeepEqual(arg, fixReq); ok == false {
		t.Fatalf("Didn't call with right request. Expected '%+v', got '%+v'.\n", fixReq, arg)
	}

	// out valid
	if ok := reflect.DeepEqual(fixOut, out); ok == false {
		t.Fatalf("Didn't receive right output. Expected '%#v', got '%#v'\n", fixOut, out)
	}

	// Check dump file
	dumpFile := fmt.Sprintf("%s/paradas_linea_%v.json", DumpPath, linea.Codigo)
	if _, err := os.Stat(dumpFile); os.IsNotExist(err) {
		t.Fatal("Didn't create a dump file")
	}

	var fout []*clcitybusapi.ParadaLinea
	fcon, err := ioutil.ReadFile(dumpFile)
	if err != nil {
		t.Fatalf("Unexpected error, %v", err)
	}

	err = json.Unmarshal(fcon, &fout)
	if err != nil {
		t.Fatalf("Unexpected error, %v", err)
	}

	if ok := reflect.DeepEqual(out, fout); ok == false {
		t.Fatalf("Didn't receive right output. Expected '%#v', got '%#v'\n", fixOut, out)
	}
}

func TestParadaService_ParadasPorLinea_ReadsFromDump(t *testing.T) {
	CreateDump()
	defer ClearDump()

	linea, _, fixOut, _, spy := fixtures.TestParadaServiceParadasPorLinea(t)

	// Write dump file
	err := dump.Write(fixOut, fmt.Sprintf("%s/paradas_linea_%v.json", DumpPath, linea.Codigo))
	if err != nil {
		t.Fatal("Error while writing dump file", err)
	}

	scli := NewSOAPClient("", false, nil)
	scli.RecuperarParadasCompletoPorLineaSpy = spy

	cli := client.NewClient(scli, DumpPath)

	out, err := cli.ParadaService().ParadasPorLinea(linea)
	if err != nil {
		t.Fatalf("Unexpected error: '%v'", err)
	}

	// Called call
	spy = scli.RecuperarParadasCompletoPorLineaSpy
	if spy.Invoked == true {
		t.Fatal("Invoked Call")
	}

	// out valid
	if ok := reflect.DeepEqual(fixOut, out); ok == false {
		t.Fatalf("Didn't receive right output. Expected '%#v', got '%#v'\n", fixOut, out)
	}
}

func TestParadaService_ParadasPorEmpresa(t *testing.T) {
	CreateDump()
	defer ClearDump()

	emp, _, _, _, _, _, _, flinresp, fparresp, fOut := fixtures.TestParadaServiceParadasPorEmpresa(t)

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

	cli := client.NewClient(scli, DumpPath)

	out, err := cli.ParadaService().ParadasPorEmpresa(emp)
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
			if value.Codigo == fvalue.Codigo {
				found = true
				// if ok := reflect.DeepEqual(fvalue, value); ok == true {
				// 	found = true
				// }
			}
		}

		if found == false {
			buff := bytes.NewBufferString(fmt.Sprintf("Couldn't find\n'%#v'\namong the results:\n", fvalue))
			for _, v := range out {
				buff.WriteString(fmt.Sprintf("'%#v'\n", v))
			}
			t.Fatalf(buff.String())
		}
	}

	// Check dump file
	dumpFile := fmt.Sprintf("%s/%s", DumpPath, "paradas_empresa.json")
	if _, err := os.Stat(dumpFile); os.IsNotExist(err) {
		t.Fatal("Didn't create a dump file")
	}

	var fout []*clcitybusapi.Parada
	fcon, err := ioutil.ReadFile(dumpFile)
	if err != nil {
		t.Fatalf("Unexpected error, %v", err)
	}

	err = json.Unmarshal(fcon, &fout)
	if err != nil {
		t.Fatalf("Unexpected error, %v", err)
	}

	if ok := reflect.DeepEqual(out, fout); ok == false {
		t.Fatalf("Didn't receive right output. Expected '%#v', got '%#v'\n", fOut, out)
	}
}

func TestParadaService_ParadasPorEmpresa_ReadFromDump(t *testing.T) {
	CreateDump()
	defer ClearDump()

	emp, _, _, _, _, _, _, flinresp, fparresp, fOut := fixtures.TestParadaServiceParadasPorEmpresa(t)

	err := dump.Write(fOut, fmt.Sprintf("%s/paradas_empresa.json", DumpPath))
	if err != nil {
		t.Fatal("Failed to write fixture dump")
	}

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

	cli := client.NewClient(scli, DumpPath)

	out, err := cli.ParadaService().ParadasPorEmpresa(emp)
	if err != nil {
		t.Fatalf("Unexpected error: '%v'\n", err)
	}

	// out valid
	if len(fOut) != len(out) {
		t.Fatal("Didn't return the expected number of elements")
	}

	// Called call
	spy := scli.RecuperarLineasPorCodigoEmpresaSpy
	if spy.Invoked == true {
		t.Fatal("Invoked Call")
	}
	spy = scli.RecuperarParadasCompletoPorLineaSpy
	if spy.Invoked == true {
		t.Fatal("Invoked Call")
	}

	// out valid
	if ok := reflect.DeepEqual(fOut, out); ok == false {
		t.Fatalf("Didn't receive right output. Expected '%#v', got '%#v'\n", fOut, out)
	}
}
