package fixtures

import (
	"time"

	"bitbucket.org/friasdesign/clcity/pkg/gtfs/files"
)

type FrequencyFlattenTestCase struct {
	Input  files.Frequency
	Output []string
}

func TestFrequencyFlatten() (
	fix FrequencyFlattenTestCase,
	fixWithout FrequencyFlattenTestCase,
) {
	x := files.Frequency{
		Trip:        &files.Trip{ID: "001"},
		StartTime:   time.Date(2000, 1, 1, 20, 0, 0, 0, time.UTC),
		EndTime:     time.Date(2000, 1, 1, 22, 0, 0, 0, time.UTC),
		HeadwaySecs: 600,
		ExactTimes:  false,
	}
	fix = FrequencyFlattenTestCase{
		Input: x,
		Output: []string{
			// trip_id
			"001",
			// start_time
			"20:00:00",
			// end_time
			"22:00:00",
			// headway_secs
			"600",
			// exact_times
			"0",
		},
	}
	xWithout := x
	xWithout.Trip = nil
	xWithout.ExactTimes = true
	fixWithout = FrequencyFlattenTestCase{
		Input: xWithout,
		Output: []string{
			// trip_id
			"",
			// start_time
			"20:00:00",
			// end_time
			"22:00:00",
			// headway_secs
			"600",
			// exact_times
			"1",
		},
	}

	return
}
