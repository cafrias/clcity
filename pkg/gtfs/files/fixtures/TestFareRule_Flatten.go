package fixtures

import (
	"github.com/cafrias/clcity/pkg/gtfs/files"
)

type FareRuleFlattenTestCase struct {
	Input  files.FareRule
	Output []string
}

func TestFareRuleFlatten() (
	fix FareRuleFlattenTestCase,
	fixWithout FareRuleFlattenTestCase,
) {
	x := files.FareRule{
		Fare:        &files.Fare{ID: "FA001"},
		Route:       &files.Route{ID: "RO001"},
		Origin:      "Z001",
		Destination: "Z002",
		Contains:    "Z003",
	}
	fix = FareRuleFlattenTestCase{
		Input: x,
		Output: []string{
			// fare_id
			"FA001",
			// route_id
			"RO001",
			// origin_id
			"Z001",
			// destination_id
			"Z002",
			// contains_id
			"Z003",
		},
	}
	xWithout := x
	xWithout.Fare = nil
	xWithout.Route = nil
	fixWithout = FareRuleFlattenTestCase{
		Input: xWithout,
		Output: []string{
			// fare_id
			"",
			// route_id
			"",
			// origin_id
			"Z001",
			// destination_id
			"Z002",
			// contains_id
			"Z003",
		},
	}

	return
}
