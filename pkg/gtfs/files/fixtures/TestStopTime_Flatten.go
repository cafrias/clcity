package fixtures

import (
	"time"

	"bitbucket.org/friasdesign/clcity/pkg/gtfs/files"
)

type StopTimeFlattenTestCase struct {
	Input  files.StopTime
	Output []string
}

func TestStopTimeFlatten() (
	fix StopTimeFlattenTestCase,
	fixWithout StopTimeFlattenTestCase,
) {
	arrT := time.Date(2000, 1, 1, 22, 00, 00, 00, time.UTC)
	depT := time.Date(2000, 1, 1, 22, 10, 00, 00, time.UTC)
	stop := files.StopTime{
		Trip:              &files.Trip{ID: "TR001"},
		ArrivalTime:       arrT,
		DepartureTime:     depT,
		Stop:              &files.Stop{ID: "ST001"},
		StopSequence:      1,
		StopHeadsign:      "head",
		PickupType:        1,
		DropOffType:       1,
		ShapeDistTravaled: 1.5,
		Timepoint:         0,
	}
	stopWithout := stop
	stopWithout.Trip = nil
	stopWithout.Stop = nil
	fix = StopTimeFlattenTestCase{
		Input: stop,
		Output: []string{
			// trip_id
			"TR001",
			// arrival_time
			"22:00:00",
			// departure_time
			"22:10:00",
			// stop_id
			"ST001",
			// stop_sequence
			"1",
			// stop_headsign
			"head",
			// pickup_type
			"1",
			// drop_off_type
			"1",
			// shape_dist_traveled
			"1.5",
			// timepoint
			"0",
		},
	}
	fixWithout = StopTimeFlattenTestCase{
		Input: stopWithout,
		Output: []string{
			// trip_id
			"",
			// arrival_time
			"22:00:00",
			// departure_time
			"22:10:00",
			// stop_id
			"",
			// stop_sequence
			"1",
			// stop_headsign
			"head",
			// pickup_type
			"1",
			// drop_off_type
			"1",
			// shape_dist_traveled
			"1.5",
			// timepoint
			"0",
		},
	}

	return
}
