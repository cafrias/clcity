package files_test

import (
	"reflect"
	"testing"

	"bitbucket.org/friasdesign/clcity/pkg/gtfs"
	"bitbucket.org/friasdesign/clcity/pkg/gtfs/files"
	"bitbucket.org/friasdesign/clcity/pkg/gtfs/files/fixtures"
)

func TestShapes_FileName(t *testing.T) {
	ag := files.Shapes{}
	out := ag.FileName()

	if out != files.ShapesFileName {
		t.Fatal("Wrong filename returned.")
	}
}

func TestShapes_FileHeaders(t *testing.T) {
	ag := files.Shapes{}
	out := ag.FileHeaders()

	if reflect.DeepEqual(out, files.ShapesFileHeaders) == false {
		t.Fatalf("File headers wrong!\nExpected:\n%v\nReceived:\n%v\n", files.ShapesFileHeaders, out)
	}
}

func TestShapes_FileEntries(t *testing.T) {
	ag := files.Shape{ID: "001"}
	ags := files.Shapes{
		ag.ID: []files.Shape{
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

func TestShape_Flatten(t *testing.T) {
	fix, _ := fixtures.TestShapeFlatten()

	out := fix.Input.Flatten()
	if ok := reflect.DeepEqual(out, fix.Output); ok == false {
		t.Fatalf("Expected output '%+v' to be '%+v'", out, fix.Output)
	}
}

func TestShape_Flatten_Without(t *testing.T) {
	_, fix := fixtures.TestShapeFlatten()

	out := fix.Input.Flatten()
	if ok := reflect.DeepEqual(out, fix.Output); ok == false {
		t.Fatalf("Expected output '%+v' to be '%+v'", out, fix.Output)
	}
}
