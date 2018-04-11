package files_test

import (
	"reflect"
	"testing"

	"bitbucket.org/friasdesign/clcity/pkg/gtfs"
	"bitbucket.org/friasdesign/clcity/pkg/gtfs/files"
	"bitbucket.org/friasdesign/clcity/pkg/gtfs/files/fixtures"
)

func TestFrequencies_FileName(t *testing.T) {
	ag := files.Frequencies{}
	out := ag.FileName()

	if out != files.FrequenciesFileName {
		t.Fatal("Wrong filename returned.")
	}
}

func TestFrequencies_FileHeaders(t *testing.T) {
	ag := files.Frequencies{}
	out := ag.FileHeaders()

	if reflect.DeepEqual(out, files.FrequenciesFileHeaders) == false {
		t.Fatalf("File headers wrong!\nExpected:\n%v\nReceived:\n%v\n", files.FrequenciesFileHeaders, out)
	}
}

func TestFrequencies_FileEntries(t *testing.T) {
	ag := files.Frequency{Trip: &files.Trip{ID: "001"}}
	ags := files.Frequencies{
		ag.Trip.ID: []files.Frequency{
			ag,
		},
	}
	fOut := []gtfs.FeedFileEntry{
		&ag,
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

func TestFrequency_Flatten(t *testing.T) {
	fix, _ := fixtures.TestFrequencyFlatten()

	out := fix.Input.Flatten()
	if ok := reflect.DeepEqual(out, fix.Output); ok == false {
		t.Fatalf("Expected output '%+v' to be '%+v'", out, fix.Output)
	}
}

func TestFrequency_Flatten_Without(t *testing.T) {
	_, fix := fixtures.TestFrequencyFlatten()

	out := fix.Input.Flatten()
	if ok := reflect.DeepEqual(out, fix.Output); ok == false {
		t.Fatalf("Expected output '%+v' to be '%+v'", out, fix.Output)
	}
}
