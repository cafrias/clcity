package client_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"

	"github.com/cafrias/clcity/internal/clcitybusapi"
	"github.com/cafrias/clcity/internal/clcitybusapi/dump"
	"github.com/cafrias/clcity/internal/clcitybusapi/mock"

	"github.com/cafrias/clcity/internal/clcitybusapi/client/fixtures"

	"github.com/cafrias/clcity/internal/clcitybusapi/client"
)

func TestLineaService_LineasPorEmpresa(t *testing.T) {
	CreateDump()
	defer ClearDump()

	fEmp, _, sRec, _, sPar, fLin, _, _, _, sLin, fLinDump := fixtures.TestLineaServiceLineasPorEmpresa(t)

	scli := NewSOAPClient("", false, nil)
	scli.RecuperarLineasPorCodigoEmpresaSpy = sLin

	cli := client.NewClient(scli, DumpPath)

	// Mock recorrido service.
	cli.SetRecorridoService(&mock.RecorridoService{
		RecorridoDeLineaSpy: sRec,
	})

	// Mock parada service.
	cli.SetParadaService(&mock.ParadaService{
		ParadasPorLineaSpy: sPar,
	})

	out, err := cli.LineaService().LineasPorEmpresa(fEmp)
	if err != nil {
		t.Fatalf("Unexpected error: '%v'", err)
	}

	// Called deps
	if sLin.Invoked == false {
		t.Fatal("Didn't invoke `RecuperarLineasPorCodigoEmpresa`")
	}
	if sRec.Invoked == false {
		t.Fatal("Didn't invoke `RecorridoDeLinea`")
	}
	if sPar.Invoked == false {
		t.Fatal("Didn't invoke `ParadasPorLinea`")
	}

	// out valid
	for _, exp := range fLin {
		for _, rec := range out {
			if exp.Codigo == rec.Codigo {
				if ok := reflect.DeepEqual(exp, rec); ok == false {
					t.Fatalf("Outputs are different. EXPECTED '%s' RECEIVED '%s'", spew.Sdump(exp), spew.Sdump(rec))
				}
			}
		}
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

	if ok := reflect.DeepEqual(fLinDump, fout); ok == false {
		t.Fatalf("Outputs are different. EXPECTED '%s'\nRECEIVED '%s'", spew.Sdump(fLin), spew.Sdump(fout))
	}
}

func TestLineaService_LineasPorEmpresa_ReadsFromDump(t *testing.T) {
	CreateDump()
	defer ClearDump()

	fEmp, _, _, _, sPar, fLin, _, _, _, sLin, _ := fixtures.TestLineaServiceLineasPorEmpresa(t)

	err := dump.Write(fLin, fmt.Sprintf("%s/lineas.json", DumpPath))
	if err != nil {
		t.Fatal("Failed to write fixture dump")
	}

	scli := NewSOAPClient("", false, nil)
	scli.RecuperarLineasPorCodigoEmpresaSpy = sLin

	cli := client.NewClient(scli, DumpPath)

	// Mock parada service.
	cli.SetParadaService(&mock.ParadaService{
		ParadasPorLineaSpy: sPar,
	})

	out, err := cli.LineaService().LineasPorEmpresa(fEmp)
	if err != nil {
		t.Fatalf("Unexpected error: '%v'", err)
	}

	// Called call
	if sLin.Invoked == true {
		t.Fatal("Invoked Call")
	}

	// out valid
	if ok := reflect.DeepEqual(fLin, out); ok == false {
		t.Fatalf("Outputs are different. EXPECTED '%s'\nRECEIVED '%s'", spew.Sdump(fLin), spew.Sdump(out))
	}
}
