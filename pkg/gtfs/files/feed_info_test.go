package files_test

import (
	"reflect"
	"testing"

	"github.com/cafrias/clcity/pkg/gtfs"
	"github.com/cafrias/clcity/pkg/gtfs/files"
	"github.com/cafrias/clcity/pkg/gtfs/files/fixtures"
)

func TestFeedInfo_FileName(t *testing.T) {
	ag := files.FeedInfo{}
	out := ag.FileName()

	if out != files.FeedInfoFileName {
		t.Fatal("Wrong filename returned.")
	}
}

func TestFeedInfo_FileHeaders(t *testing.T) {
	ag := files.FeedInfo{}
	out := ag.FileHeaders()

	if reflect.DeepEqual(out, files.FeedInfoFileHeaders) == false {
		t.Fatalf("File headers wrong!\nExpected:\n%v\nReceived:\n%v\n", files.FeedInfoFileHeaders, out)
	}
}

func TestFeedInfo_FileEntries(t *testing.T) {
	ag := files.FeedInfo{
		PublisherName: "My name",
	}
	fOut := []gtfs.FeedFileEntry{
		&ag,
	}
	out := ag.FileEntries()

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

func TestFeedInfo_Flatten(t *testing.T) {
	fix := fixtures.TestFeedInfoFlatten()

	out := fix.Input.Flatten()
	if ok := reflect.DeepEqual(out, fix.Output); ok == false {
		t.Fatalf("Expected output '%+v' to be '%+v'", out, fix.Output)
	}
}
