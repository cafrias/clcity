package files_test

import (
	"reflect"
	"testing"

	"github.com/friasdesign/clcity/pkg/gtfs"
	"github.com/friasdesign/clcity/pkg/gtfs/files"
	"github.com/friasdesign/clcity/pkg/gtfs/files/fixtures"
)

func TestFareRules_FileName(t *testing.T) {
	ag := files.FareRules{}
	out := ag.FileName()

	if out != files.FareRulesFileName {
		t.Fatal("Wrong filename returned.")
	}
}

func TestFareRules_FileHeaders(t *testing.T) {
	ag := files.FareRules{}
	out := ag.FileHeaders()

	if reflect.DeepEqual(out, files.FareRulesFileHeaders) == false {
		t.Fatalf("File headers wrong!\nExpected:\n%v\nReceived:\n%v\n", files.FareRulesFileHeaders, out)
	}
}

func TestFareRules_FileEntries(t *testing.T) {
	ag := &files.FareRule{Fare: &files.Fare{ID: "FA001"}}
	ags := files.FareRules{
		ag.Fare.ID: []*files.FareRule{
			ag,
		},
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

func TestFareRule_Flatten(t *testing.T) {
	fix, _ := fixtures.TestFareRuleFlatten()

	out := fix.Input.Flatten()
	if ok := reflect.DeepEqual(out, fix.Output); ok == false {
		t.Fatalf("Expected output '%+v' to be '%+v'", out, fix.Output)
	}
}

func TestFareRules_Flatten_Without(t *testing.T) {
	_, fix := fixtures.TestFareRuleFlatten()

	out := fix.Input.Flatten()
	if ok := reflect.DeepEqual(out, fix.Output); ok == false {
		t.Fatalf("Expected output '%+v' to be '%+v'", out, fix.Output)
	}
}
