package files_test

import (
	"reflect"
	"testing"

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

func TestStops_Flatten(t *testing.T) {
	ag := files.Stops{
		"001": files.Stop{
			ID: "001",
		},
	}
	fOut := [][]string{
		{
			"stop_id", "stop_code", "stop_name", "stop_desc", "stop_lat", "stop_lon", "zone_id", "stop_url", "location_type", "parent_station", "stop_timezone", "wheelchair_boarding",
		},
		{
			"001", "", "", "", "0", "0", "", "", "0", "", "", "0",
		},
	}
	out := ag.Flatten()

	if reflect.DeepEqual(out, fOut) == false {
		t.Fatalf("Invalid output. Expected:\n%+v\nReceived:\n%+v\n", fOut, out)
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
