package fixtures

import (
	"github.com/cafrias/clcity/pkg/gtfs/files"
)

type TripFlattenTestCase struct {
	Input  files.Trip
	Output []string
}

func TestTripFlatten() (
	fix TripFlattenTestCase,
	fixWag TripFlattenTestCase,
) {
	r := files.Trip{
		Route:        &files.Route{ID: "RO001"},
		Service:      &files.Service{ID: "SE001"},
		ID:           "001",
		Headsign:     "head",
		ShortName:    "short",
		DirectionID:  1,
		BlockID:      "BL001",
		Shape:        &files.Shape{ID: "SH001"},
		Wheelchair:   0,
		BikesAllowed: 1,
	}
	rw := r
	rw.Route = nil
	rw.Service = nil
	rw.Shape = nil
	fix = TripFlattenTestCase{
		Input: r,
		Output: []string{
			// route_id
			"RO001",
			// service_id
			"SE001",
			// trip_id
			"001",
			// trip_headsign
			"head",
			// trip_short_name
			"short",
			// direction_id
			"1",
			// block_id
			"BL001",
			// shape_id
			"SH001",
			// wheelchair_accessible
			"0",
			// bikes_allowed
			"1",
		},
	}
	fixWag = TripFlattenTestCase{
		Input: rw,
		Output: []string{
			// route_id
			"",
			// service_id
			"",
			// trip_id
			"001",
			// trip_headsign
			"head",
			// trip_short_name
			"short",
			// direction_id
			"1",
			// block_id
			"BL001",
			// shape_id
			"",
			// wheelchair_accessible
			"0",
			// bikes_allowed
			"1",
		},
	}

	return
}
