package client_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/dump"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/client/fixtures"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/client"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/soapclient/swparadas"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi"
)

func TestLineaService_LineasPorEmpresa(t *testing.T) {
	CreateDump()
	defer ClearDump()

	spy, fixReq, fixOut, _ := fixtures.TestLineaServiceLineasPorEmpresa(t)

	scli := NewSOAPClient("", false, nil)
	scli.RecuperarLineasPorCodigoEmpresaSpy = spy

	cli := client.NewClient(scli, DumpPath)

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
	arg, _ := spy.Args[0][0].(*swparadas.RecuperarLineasPorCodigoEmpresa)
	if ok := reflect.DeepEqual(arg, fixReq); ok == false {
		t.Fatalf("Didn't call with right request. Expected '%+v', got '%+v'.\n", fixReq, arg)
	}

	// out valid
	if ok := reflect.DeepEqual(fixOut, out); ok == false {
		t.Fatalf("Didn't receive right output. Expected '%#v', got '%#v'\n", fixOut, out)
	}

	// Check dump file
	dumpFile := fmt.Sprintf("%s/%s", DumpPath, "lineas.json")
	if _, err := os.Stat(dumpFile); os.IsNotExist(err) {
		t.Fatal("Didn't create a dump file")
	}

	var fout []*clcitybusapi.Linea
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

func TestLineaService_LineasPorEmpresa_ReadsFromDump(t *testing.T) {
	CreateDump()
	defer ClearDump()

	spy, _, fixOut, _ := fixtures.TestLineaServiceLineasPorEmpresa(t)

	err := dump.Write(fixOut, fmt.Sprintf("%s/lineas.json", DumpPath))
	if err != nil {
		t.Fatal("Failed to write fixture dump")
	}

	scli := NewSOAPClient("", false, nil)
	scli.RecuperarLineasPorCodigoEmpresaSpy = spy

	cli := client.NewClient(scli, DumpPath)

	out, err := cli.LineaService().LineasPorEmpresa(355)
	if err != nil {
		t.Fatalf("Unexpected error: '%v'", err)
	}

	// Called call
	spy = scli.RecuperarLineasPorCodigoEmpresaSpy
	if spy.Invoked == true {
		t.Fatal("Invoked Call")
	}

	// out valid
	if ok := reflect.DeepEqual(fixOut, out); ok == false {
		t.Fatalf("Didn't receive right output. Expected '%#v', got '%#v'\n", fixOut, out)
	}
}
