package files_test

import (
	"reflect"
	"testing"

	"bitbucket.org/friasdesign/clcity/pkg/gtfs/files/fixtures"

	"bitbucket.org/friasdesign/clcity/pkg/gtfs"
	"bitbucket.org/friasdesign/clcity/pkg/gtfs/files"
)

func TestTrips_FileName(t *testing.T) {
	ag := files.Trips{}
	out := ag.FileName()

	if out != files.TripsFileName {
		t.Fatal("Wrong filename returned.")
	}
}

func TestTrips_FileHeaders(t *testing.T) {
	ag := files.Trips{}
	out := ag.FileHeaders()

	if reflect.DeepEqual(out, files.TripsFileHeaders) == false {
		t.Fatalf("File headers wrong!\nExpected:\n%v\nReceived:\n%v\n", files.TripsFileHeaders, out)
	}
}

func TestTrips_FileEntries(t *testing.T) {
	ag := &files.Trip{ID: "001"}
	ags := files.Trips{
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

func TestTrip_Flatten(t *testing.T) {
	fix, _ := fixtures.TestTripFlatten()

	out := fix.Input.Flatten()
	if ok := reflect.DeepEqual(out, fix.Output); ok == false {
		t.Fatalf("Expected output '%+v' to be '%+v'", out, fix.Output)
	}
}

func TestTrip_Flatten_WithoutParent(t *testing.T) {
	_, fix := fixtures.TestTripFlatten()

	out := fix.Input.Flatten()
	if ok := reflect.DeepEqual(out, fix.Output); ok == false {
		t.Fatalf("Expected output '%+v' to be '%+v'", out, fix.Output)
	}
}
