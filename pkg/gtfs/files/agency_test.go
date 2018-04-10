package files_test

import (
	"net/url"
	"reflect"
	"testing"

	"bitbucket.org/friasdesign/clcity/pkg/gtfs"
	"bitbucket.org/friasdesign/clcity/pkg/gtfs/files"
)

func TestAgencies_FileName(t *testing.T) {
	ag := files.Agencies{}
	out := ag.FileName()

	if out != files.AgencyFileName {
		t.Fatal("Wrong filename returned.")
	}
}

func TestAgencies_Flatten(t *testing.T) {
	ag := files.Agencies{
		"001": files.Agency{
			ID:    "001",
			Name:  "city",
			Phone: "123456",
		},
	}
	fOut := [][]string{
		{
			"agency_id", "agency_name", "agency_url", "agency_timezone", "agency_lang", "agency_phone", "agency_fare_url", "agency_email",
		},
		{
			"001", "city", "", "", "", "123456", "", "",
		},
	}
	out := ag.Flatten()

	if reflect.DeepEqual(out, fOut) == false {
		t.Fatalf("Invalid output. Expected:\n%+v\nReceived:\n%+v\n", fOut, out)
	}
}

func TestAgency_Validate(t *testing.T) {
	type testCase struct {
		input  files.Agency
		output bool
		err    *gtfs.ErrValidation
	}
	fixURL, _ := url.Parse("https://github.com")
	fix := testCase{
		input: files.Agency{
			ID:       "ASD356",
			Name:     "City",
			URL:      *fixURL,
			Timezone: gtfs.Timezone("America/Argentina/Ushuaia"),
		},
		output: true,
		err:    nil,
	}

	OK, err := fix.input.Validate()
	if err != fix.err {
		t.Fatalf("Unexpected error: %v \n", err)
	}

	if OK != fix.output {
		t.Fatalf("Expected output '%v' to be '%v'", OK, fix.output)
	}
}

func TestAgency_Validate_ErrValidation(t *testing.T) {
	type testCase struct {
		input  files.Agency
		output bool
		err    *gtfs.ErrValidation
	}
	fix := testCase{
		input: files.Agency{
			ID:       "ASD356",
			Name:     "",
			URL:      url.URL{},
			Timezone: gtfs.Timezone("America/Argentina/Ushuaiaas"),
			Lang:     "asd",
		},
		output: false,
		err: &gtfs.ErrValidation{
			File: "agency.txt",
			Fields: map[string]string{
				"agency_name":     "",
				"agency_timezone": "America/Argentina/Ushuaiaas",
				"agency_lang":     "asd",
				"agency_url":      "",
			},
		},
	}

	OK, err := fix.input.Validate()
	if ok := reflect.DeepEqual(err, fix.err); ok == false {
		t.Fatalf("Unexpected error: %v.\n Expected: %v \n", err, fix.err)
	}

	if OK != fix.output {
		t.Fatalf("Expected output '%v' to be '%v'", OK, fix.output)
	}
}

func TestAgency_Flatten(t *testing.T) {
	type testCase struct {
		input  files.Agency
		output []string
	}
	fixURL, _ := url.Parse("https://github.com")
	agency := files.Agency{
		ID:       "ASD356",
		Name:     "City",
		URL:      *fixURL,
		Timezone: gtfs.Timezone("America/Argentina/Ushuaia"),
		Lang:     gtfs.LanguageISO6391("es"),
	}
	fix := testCase{
		input: agency,
		output: []string{
			// agency_id
			"ASD356",
			// agency_name
			agency.Name,
			// agency_url
			"https://github.com",
			// agency_timezone
			"America/Argentina/Ushuaia",
			// agency_lang
			"es",
			// agency_phone
			agency.Phone,
			// agency_fare_url
			"",
			// agency_email
			"",
		},
	}

	out := fix.input.Flatten()
	if ok := reflect.DeepEqual(out, fix.output); ok == false {
		t.Fatalf("Expected output '%+v' to be '%+v'", out, fix.output)
	}
}
