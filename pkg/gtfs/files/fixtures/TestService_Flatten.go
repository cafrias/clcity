package fixtures

import (
	"time"

	"github.com/cafrias/clcity/pkg/gtfs/files"
)

type ServiceFlattenTestCase struct {
	Input  files.Service
	Output []string
}

func TestServiceFlatten() (
	fix ServiceFlattenTestCase,
) {
	stDate := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	enDate := time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC)
	x := files.Service{
		ID:        "001",
		Mon:       true,
		Tue:       true,
		Wed:       true,
		Thu:       true,
		Fri:       true,
		Sat:       true,
		Sun:       false,
		StartDate: stDate,
		EndDate:   enDate,
	}
	fix = ServiceFlattenTestCase{
		Input: x,
		Output: []string{
			// service_id
			"001",
			// monday
			"1",
			// tuesday
			"1",
			// wednesday
			"1",
			// thursday
			"1",
			// friday
			"1",
			// saturday
			"1",
			// sunday
			"0",
			// start_date
			"20000101",
			// end_date
			"20010101",
		},
	}

	return
}
