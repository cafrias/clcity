package files_test

import (
	"reflect"
	"testing"

	"bitbucket.org/friasdesign/clcity/pkg/gtfs"
	"bitbucket.org/friasdesign/clcity/pkg/gtfs/files"
	"bitbucket.org/friasdesign/clcity/pkg/gtfs/files/fixtures"
)

func TestRoutes_FileName(t *testing.T) {
	ag := files.Routes{}
	out := ag.FileName()

	if out != files.RoutesFileName {
		t.Fatal("Wrong filename returned.")
	}
}

func TestRoutes_FileHeaders(t *testing.T) {
	ag := files.Routes{}
	out := ag.FileHeaders()

	if reflect.DeepEqual(out, files.RoutesFileHeaders) == false {
		t.Fatalf("File headers wrong!\nExpected:\n%v\nReceived:\n%v\n", files.RoutesFileHeaders, out)
	}
}

func TestRoutes_FileEntries(t *testing.T) {
	ag := &files.Route{ID: "001"}
	ags := files.Routes{
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

func TestRoute_Flatten(t *testing.T) {
	fix, _ := fixtures.TestRouteFlatten()

	out := fix.Input.Flatten()
	if ok := reflect.DeepEqual(out, fix.Output); ok == false {
		t.Fatalf("Expected output '%+v' to be '%+v'", out, fix.Output)
	}
}

func TestRoute_Flatten_WithoutAgency(t *testing.T) {
	_, fix := fixtures.TestRouteFlatten()

	out := fix.Input.Flatten()
	if ok := reflect.DeepEqual(out, fix.Output); ok == false {
		t.Fatalf("Expected output '%+v' to be '%+v'", out, fix.Output)
	}
}
