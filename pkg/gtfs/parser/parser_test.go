package parser_test

import (
	"net/mail"
	"os"
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
	fPath := "testdata/write"
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
}
