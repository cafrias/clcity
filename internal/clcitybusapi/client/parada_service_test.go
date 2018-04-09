package client_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"bitbucket.org/friasdesign/clcity/internal/clcitybusapi"
	"bitbucket.org/friasdesign/clcity/internal/clcitybusapi/client"
	"bitbucket.org/friasdesign/clcity/internal/clcitybusapi/client/fixtures"
	"bitbucket.org/friasdesign/clcity/internal/clcitybusapi/dump"
	"bitbucket.org/friasdesign/clcity/internal/clcitybusapi/mock"
	"bitbucket.org/friasdesign/clcity/internal/clcitybusapi/soapclient/swparadas"
	"github.com/davecgh/go-spew/spew"
)

func TestParadaService_ParadasPorLinea(t *testing.T) {
	CreateDump()
	defer ClearDump()

	linea, fixReq, fixOut, _, spy, fixDump := fixtures.TestParadaServiceParadasPorLinea(t)

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

	if ok := reflect.DeepEqual(fixDump, fout); ok == false {
		t.Fatalf("Outputs are different. EXPECTED '%s' RECEIVED '%s'", spew.Sdump(fout), spew.Sdump(fixDump))
	}
}

func TestParadaService_ParadasPorLinea_ReadsFromDump(t *testing.T) {
	CreateDump()
	defer ClearDump()

	linea, _, fixOut, _, spy, _ := fixtures.TestParadaServiceParadasPorLinea(t)

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
		t.Fatalf("Outputs are different. EXPECTED '%s' RECEIVED '%s'", spew.Sdump(fixOut), spew.Sdump(out))
	}
}

func TestParadaService_ParadasPorEmpresa(t *testing.T) {
	CreateDump()
	defer ClearDump()

	fEmp, _, _, _, sLin, _, _, _, sParLin, fPar := fixtures.TestParadaServiceParadasPorEmpresa(t)

	scli := NewSOAPClient("", false, nil)
	scli.RecuperarParadasCompletoPorLineaSpy = sParLin

	cli := client.NewClient(scli, DumpPath)

	cli.SetLineaService(&mock.LineaService{
		LineasPorEmpresaSpy: sLin,
	})

	out, err := cli.ParadaService().ParadasPorEmpresa(fEmp)
	if err != nil {
		t.Fatalf("Unexpected error: '%v'\n", err)
	}

	// out valid
	for _, exp := range fPar {
		for _, rec := range out {
			if exp.Codigo == rec.Codigo {
				if ok := reflect.DeepEqual(*exp, *rec); ok == false {
					t.Fatalf("Outputs are different. EXPECTED '%s' RECEIVED '%s'", spew.Sdump(exp), spew.Sdump(rec))
				}
			}
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

	for _, exp := range fPar {
		for _, rec := range fout {
			if exp.Codigo == rec.Codigo {
				if ok := reflect.DeepEqual(*exp, *rec); ok == false {
					t.Fatalf("Outputs are different. EXPECTED '%s' RECEIVED '%s'", spew.Sdump(exp), spew.Sdump(rec))
				}
			}
		}
	}
}

func TestParadaService_ParadasPorEmpresa_ReadFromDump(t *testing.T) {
	CreateDump()
	defer ClearDump()

	fEmp, _, _, _, sLin, _, _, _, sParLin, fPar := fixtures.TestParadaServiceParadasPorEmpresa(t)

	err := dump.Write(fPar, fmt.Sprintf("%s/paradas_empresa.json", DumpPath))
	if err != nil {
		t.Fatal("Failed to write fixture dump")
	}

	scli := NewSOAPClient("", false, nil)
	scli.RecuperarParadasCompletoPorLineaSpy = sParLin

	cli := client.NewClient(scli, DumpPath)

	cli.SetLineaService(&mock.LineaService{
		LineasPorEmpresaSpy: sLin,
	})

	out, err := cli.ParadaService().ParadasPorEmpresa(fEmp)
	if err != nil {
		t.Fatalf("Unexpected error: '%v'\n", err)
	}

	// out valid
	for _, exp := range fPar {
		for _, rec := range out {
			if exp.Codigo == rec.Codigo {
				if ok := reflect.DeepEqual(*exp, *rec); ok == false {
					t.Fatalf("Outputs are different. EXPECTED '%s' RECEIVED '%s'", spew.Sdump(exp), spew.Sdump(rec))
				}
			}
		}
	}
}
