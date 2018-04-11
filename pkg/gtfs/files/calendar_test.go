package files_test

import (
	"reflect"
	"testing"

	"bitbucket.org/friasdesign/clcity/pkg/gtfs"
	"bitbucket.org/friasdesign/clcity/pkg/gtfs/files"
	"bitbucket.org/friasdesign/clcity/pkg/gtfs/files/fixtures"
)

func TestCalendar_FileName(t *testing.T) {
	ag := files.Calendar{}
	out := ag.FileName()

	if out != files.CalendarFileName {
		t.Fatal("Wrong filename returned.")
	}
}

func TestCalendar_FileHeaders(t *testing.T) {
	ag := files.Calendar{}
	out := ag.FileHeaders()

	if reflect.DeepEqual(out, files.CalendarFileHeaders) == false {
		t.Fatalf("File headers wrong!\nExpected:\n%v\nReceived:\n%v\n", files.CalendarFileHeaders, out)
	}
}

func TestCalendar_FileEntries(t *testing.T) {
	ag := files.Service{ID: "001"}
	ags := files.Calendar{
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

func TestService_Flatten(t *testing.T) {
	fix := fixtures.TestServiceFlatten()

	out := fix.Input.Flatten()
	if ok := reflect.DeepEqual(out, fix.Output); ok == false {
		t.Fatalf("Expected output '%+v' to be '%+v'", out, fix.Output)
	}
}