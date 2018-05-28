package parser_test

import (
	"bytes"
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/friasdesign/clcity/pkg/gtfs/parser/fixtures"

	"github.com/friasdesign/clcity/pkg/gtfs/parser"
)

func setUp(p string) {
	os.MkdirAll(p, os.ModePerm)
}

func tearDown(p string) {
	os.RemoveAll(p)
}

func TestParser_Write(t *testing.T) {
	fixPath := path.Join("testdata", "fixture")
	outPath := path.Join("testdata", "write")
	tearDown(outPath)
	setUp(outPath)

	feed := fixtures.Feed()

	p := parser.NewParser(outPath)

	err := p.Write(feed)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	outFiles, err := ioutil.ReadDir(outPath)
	if err != nil {
		t.Fatalf("Error while reading write directory:\n%v\n", err)
	}

	for _, outF := range outFiles {
		outCon, err := ioutil.ReadFile(path.Join(outPath, outF.Name()))
		if err != nil {
			t.Fatalf("Error while reading '%s' file from '%s' folder:\n%v\n", outF.Name(), outPath, err)
		}
		fixCon, err := ioutil.ReadFile(path.Join(fixPath, outF.Name()))
		if err != nil {
			t.Fatalf("Error while reading '%s' file from '%s' folder:\n%v\n", outF.Name(), outPath, err)
		}

		if bytes.Equal(outCon, fixCon) == false {
			t.Fatalf("Output file is different from expected fixture file for file '%s'.\n", outF.Name())
		}
	}
}
