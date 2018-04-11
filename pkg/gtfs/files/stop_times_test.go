package files_test

import (
	"reflect"
	"testing"

	"bitbucket.org/friasdesign/clcity/pkg/gtfs"
	"bitbucket.org/friasdesign/clcity/pkg/gtfs/files"
	"bitbucket.org/friasdesign/clcity/pkg/gtfs/files/fixtures"
)

func TestStopTimes_FileName(t *testing.T) {
	ag := files.StopTimes{}
	out := ag.FileName()

	if out != files.StopTimesFileName {
		t.Fatal("Wrong filename returned.")
	}
}

func TestStopTimes_FileHeaders(t *testing.T) {
	ag := files.StopTimes{}
	out := ag.FileHeaders()

	if reflect.DeepEqual(out, files.StopTimesFileHeaders) == false {
		t.Fatalf("File headers wrong!\nExpected:\n%v\nReceived:\n%v\n", files.StopsFileHeaders, out)
	}
}

func TestStopTimes_FileEntries(t *testing.T) {
	tr := files.Trip{ID: "TR001"}
	ag := files.StopTime{Trip: &tr, StopSequence: 1}
	ags := files.StopTimes{
		tr.ID: map[files.StopSequence]files.StopTime{
			ag.StopSequence: ag,
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

func TestStopTime_Flatten(t *testing.T) {
	fix, _ := fixtures.TestStopTimeFlatten()

	out := fix.Input.Flatten()
	if ok := reflect.DeepEqual(out, fix.Output); ok == false {
		t.Fatalf("Expected output '%+v' to be '%+v'", out, fix.Output)
	}
}

func TestStopTime_Flatten_WithoutParent(t *testing.T) {
	_, fix := fixtures.TestStopTimeFlatten()

	out := fix.Input.Flatten()
	if ok := reflect.DeepEqual(out, fix.Output); ok == false {
		t.Fatalf("Expected output '%+v' to be '%+v'", out, fix.Output)
	}
}
