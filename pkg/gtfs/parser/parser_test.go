package parser_test

import (
	"bytes"
	"io/ioutil"
	"net/mail"
	"os"
	"path"
	"testing"

	"bitbucket.org/friasdesign/clcity/pkg/gtfs/parser"

	"bitbucket.org/friasdesign/clcity/pkg/gtfs"
)

func setUp(p string) {
	os.MkdirAll(p, os.ModePerm)
}

func tearDown(p string) {
	os.RemoveAll(p)
}

func TestParser_Write(t *testing.T) {
	fPath := path.Join("testdata", "write")
	tearDown(fPath)
	setUp(fPath)

	feed := &gtfs.Feed{
		Agencies: map[gtfs.AgencyID]gtfs.Agency{
			"001": gtfs.Agency{
				ID: "001",
				Email: mail.Address{
					Address: "pepe@pepe.com",
				},
				Name:     "City Bus",
				Timezone: "America/Argentina/Ushuaia",
				Lang:     "es",
			},
		},
	}

	p := parser.NewParser(fPath)

	err := p.Write(feed)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	res, err := ioutil.ReadFile(path.Join(fPath, "agency.txt"))
	if err != nil {
		t.Fatalf("Unexpected error while reading output file: %v \n", err)
	}
	fix, err := ioutil.ReadFile(path.Join("testdata", "fixture", "agency.txt"))
	if err != nil {
		t.Fatalf("Unexpected error while reading fixture file: %v\n", err)
	}
	if bytes.Equal(res, fix) == false {
		t.Fatalf("Output file is different from expected fixture file.")
	}
}
