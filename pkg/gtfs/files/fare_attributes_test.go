package files_test

import (
	"reflect"
	"testing"

	"github.com/friasdesign/clcity/pkg/gtfs"
	"github.com/friasdesign/clcity/pkg/gtfs/files"
	"github.com/friasdesign/clcity/pkg/gtfs/files/fixtures"
)

func TestFares_FileName(t *testing.T) {
	ag := files.FareAttributes{}
	out := ag.FileName()

	if out != files.FareAttributesFileName {
		t.Fatal("Wrong filename returned.")
	}
}

func TestFares_FileHeaders(t *testing.T) {
	ag := files.FareAttributes{}
	out := ag.FileHeaders()

	if reflect.DeepEqual(out, files.FareAttributesFileHeaders) == false {
		t.Fatalf("File headers wrong!\nExpected:\n%v\nReceived:\n%v\n", files.FareAttributesFileHeaders, out)
	}
}

func TestFares_FileEntries(t *testing.T) {
	ag := &files.Fare{ID: "001"}
	ags := files.FareAttributes{
		ag.ID: ag,
	}
	fOut := []gtfs.FeedFileEntry{
		ag,
	}
	out := ags.FileEntries()

	for _, expected := range fOut {
		var found bool
		for _, received := range out {
			if reflect.DeepEqual(expected, received) == true {
				found = true
			}
		}
		if found == false {
			t.Fatalf("Expected:\n%v\nto be in:\n%v\n", expected, out)
		}
	}
}

func TestFare_Flatten(t *testing.T) {
	fix, _ := fixtures.TestFareFlatten()

	out := fix.Input.Flatten()
	if ok := reflect.DeepEqual(out, fix.Output); ok == false {
		t.Fatalf("Expected output '%+v' to be '%+v'", out, fix.Output)
	}
}

func TestFares_Flatten_Without(t *testing.T) {
	_, fix := fixtures.TestFareFlatten()

	out := fix.Input.Flatten()
	if ok := reflect.DeepEqual(out, fix.Output); ok == false {
		t.Fatalf("Expected output '%+v' to be '%+v'", out, fix.Output)
	}
}
