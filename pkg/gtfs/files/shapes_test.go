package files_test

import (
	"reflect"
	"testing"

	"bitbucket.org/friasdesign/clcity/pkg/gtfs"
	"bitbucket.org/friasdesign/clcity/pkg/gtfs/files"
	"bitbucket.org/friasdesign/clcity/pkg/gtfs/files/fixtures"
)

func TestShapePoints_FileName(t *testing.T) {
	ag := files.Shapes{}
	out := ag.FileName()

	if out != files.ShapesFileName {
		t.Fatal("Wrong filename returned.")
	}
}

func TestShapePoints_FileHeaders(t *testing.T) {
	ag := files.Shapes{}
	out := ag.FileHeaders()

	if reflect.DeepEqual(out, files.ShapesFileHeaders) == false {
		t.Fatalf("File headers wrong!\nExpected:\n%v\nReceived:\n%v\n", files.ShapesFileHeaders, out)
	}
}

func TestShapePoints_FileEntries(t *testing.T) {
	p := files.ShapePoint{}
	ag := &files.Shape{ID: "001", Points: []files.ShapePoint{p}}
	ags := files.Shapes{
		ag.ID: ag,
	}
	fOut := []gtfs.FeedFileEntry{
		&p,
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

func TestShapePoint_Flatten(t *testing.T) {
	fix, _ := fixtures.TestShapePointFlatten()

	out := fix.Input.Flatten()
	if ok := reflect.DeepEqual(out, fix.Output); ok == false {
		t.Fatalf("Expected output '%+v' to be '%+v'", out, fix.Output)
	}
}

func TestShapePoint_Flatten_Without(t *testing.T) {
	_, fix := fixtures.TestShapePointFlatten()

	out := fix.Input.Flatten()
	if ok := reflect.DeepEqual(out, fix.Output); ok == false {
		t.Fatalf("Expected output '%+v' to be '%+v'", out, fix.Output)
	}
}
