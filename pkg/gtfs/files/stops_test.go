package files_test

import (
	"reflect"
	"testing"

	"bitbucket.org/friasdesign/clcity/pkg/gtfs"
	"bitbucket.org/friasdesign/clcity/pkg/gtfs/files"
	"bitbucket.org/friasdesign/clcity/pkg/gtfs/files/fixtures"
)

func TestStops_FileName(t *testing.T) {
	ag := files.Stops{}
	out := ag.FileName()

	if out != files.StopsFileName {
		t.Fatal("Wrong filename returned.")
	}
}

func TestStops_FileHeaders(t *testing.T) {
	ag := files.Stops{}
	out := ag.FileHeaders()

	if reflect.DeepEqual(out, files.StopsFileHeaders) == false {
		t.Fatalf("File headers wrong!\nExpected:\n%v\nReceived:\n%v\n", files.StopsFileHeaders, out)
	}
}

func TestStops_FileEntries(t *testing.T) {
	ag := files.Stop{ID: "001"}
	ags := files.Stops{
		ag.ID: ag,
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
func TestStop_Flatten(t *testing.T) {
	fix, _ := fixtures.TestStopFlatten()

	out := fix.Input.Flatten()
	if ok := reflect.DeepEqual(out, fix.Output); ok == false {
		t.Fatalf("Expected output '%+v' to be '%+v'", out, fix.Output)
	}
}

func TestStop_Flatten_WithoutParent(t *testing.T) {
	_, fix := fixtures.TestStopFlatten()

	out := fix.Input.Flatten()
	if ok := reflect.DeepEqual(out, fix.Output); ok == false {
		t.Fatalf("Expected output '%+v' to be '%+v'", out, fix.Output)
	}
}
