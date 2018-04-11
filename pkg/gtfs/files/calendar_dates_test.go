package files_test

import (
	"reflect"
	"testing"
	"time"

	"bitbucket.org/friasdesign/clcity/pkg/gtfs/date"
	"bitbucket.org/friasdesign/clcity/pkg/gtfs/files/fixtures"

	"bitbucket.org/friasdesign/clcity/pkg/gtfs"
	"bitbucket.org/friasdesign/clcity/pkg/gtfs/files"
)

func TestCalendarDates_FileName(t *testing.T) {
	ag := files.Calendar{}
	out := ag.FileName()

	if out != files.CalendarFileName {
		t.Fatal("Wrong filename returned.")
	}
}

func TestCalendarDates_FileHeaders(t *testing.T) {
	ag := files.Calendar{}
	out := ag.FileHeaders()

	if reflect.DeepEqual(out, files.CalendarFileHeaders) == false {
		t.Fatalf("File headers wrong!\nExpected:\n%v\nReceived:\n%v\n", files.CalendarFileHeaders, out)
	}
}

func TestCalendarDates_FileEntries(t *testing.T) {
	ag := files.CalendarDate{
		Service: &files.Service{ID: "SE001"},
		Date:    time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
	}
	ags := files.CalendarDates{
		ag.Service.ID: map[date.Date]files.CalendarDate{
			date.FormatDate(ag.Date): ag,
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

func TestCalendarDate_Flatten(t *testing.T) {
	fix, _ := fixtures.TestCalendarDateFlatten()

	out := fix.Input.Flatten()
	if ok := reflect.DeepEqual(out, fix.Output); ok == false {
		t.Fatalf("Expected output '%+v' to be '%+v'", out, fix.Output)
	}
}

func TestRoute_Flatten_Without(t *testing.T) {
	_, fix := fixtures.TestCalendarDateFlatten()

	out := fix.Input.Flatten()
	if ok := reflect.DeepEqual(out, fix.Output); ok == false {
		t.Fatalf("Expected output '%+v' to be '%+v'", out, fix.Output)
	}
}
