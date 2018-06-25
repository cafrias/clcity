package kml_test

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/cafrias/clcity/internal/clcitybusapi/kml/fixtures"

	"github.com/cafrias/clcity/internal/clcitybusapi/kml"

	"github.com/cafrias/clcity/internal/clcitybusapi"
)

const dumpPath = "testdata/kml"

func setUp() {
	_ = os.MkdirAll(dumpPath, os.ModePerm)
}

func tearDown() {
	os.RemoveAll(dumpPath)
}

func TestKML_Generate(t *testing.T) {
	tearDown()
	setUp()

	kmlPath := filepath.Join(dumpPath, "test.kml")

	fEmp := fixtures.TestKMLGenerate(t)

	err := kml.Generate(fEmp, kmlPath)
	if err != nil {
		t.Fatalf("Unexpected error: '%v'\n", err)
	}

	// Compare output files
	fixFile, err := ioutil.ReadFile("testdata/fixture.kml")
	if err != nil {
		t.Fatalf("Error while reading fixture file: %v\n", err)
	}
	outFile, err := ioutil.ReadFile(kmlPath)
	if err != nil {
		t.Fatalf("Error while reading output file: %v\n", err)
	}

	if bytes.Equal(fixFile, outFile) == false {
		t.Fatal("Files are different!")
	}
}

func TestKML_Generate_ErrNoLineas(t *testing.T) {
	emp := &clcitybusapi.Empresa{
		Paradas: []*clcitybusapi.Parada{
			&clcitybusapi.Parada{},
		},
	}

	err := kml.Generate(emp, "")
	if err != kml.ErrNoLineas {
		t.Fatalf("Unexpected error: %v\n", err)
	}
}

func TestKML_Generate_ErrNoParadas(t *testing.T) {
	emp := &clcitybusapi.Empresa{}

	err := kml.Generate(emp, "")
	if err != kml.ErrNoParadas {
		t.Fatalf("Unexpected error: %v\n", err)
	}
}
