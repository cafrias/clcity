package files_test

import (
	"reflect"
	"testing"

	"bitbucket.org/friasdesign/clcity/pkg/gtfs"
	"bitbucket.org/friasdesign/clcity/pkg/gtfs/files"
	"bitbucket.org/friasdesign/clcity/pkg/gtfs/files/fixtures"
)

func TestTransfers_FileName(t *testing.T) {
	ag := files.Transfers{}
	out := ag.FileName()

	if out != files.TransfersFileName {
		t.Fatal("Wrong filename returned.")
	}
}

func TestTransfers_FileHeaders(t *testing.T) {
	ag := files.Transfers{}
	out := ag.FileHeaders()

	if reflect.DeepEqual(out, files.TransfersFileHeaders) == false {
		t.Fatalf("File headers wrong!\nExpected:\n%v\nReceived:\n%v\n", files.TransfersFileHeaders, out)
	}
}

func TestTransfers_FileEntries(t *testing.T) {
	sf := files.Stop{ID: "ST001"}
	st := files.Stop{ID: "ST002"}
	ag := files.Transfer{
		From: &sf,
		To:   &st,
	}
	ags := files.Transfers{
		sf.ID: map[files.StopID]files.Transfer{
			st.ID: ag,
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

func TestTransfer_Flatten(t *testing.T) {
	fix, _ := fixtures.TestTransferFlatten()

	out := fix.Input.Flatten()
	if ok := reflect.DeepEqual(out, fix.Output); ok == false {
		t.Fatalf("Expected output '%+v' to be '%+v'", out, fix.Output)
	}
}

func TestTransfer_Flatten_WithoutParent(t *testing.T) {
	_, fix := fixtures.TestTransferFlatten()

	out := fix.Input.Flatten()
	if ok := reflect.DeepEqual(out, fix.Output); ok == false {
		t.Fatalf("Expected output '%+v' to be '%+v'", out, fix.Output)
	}
}
