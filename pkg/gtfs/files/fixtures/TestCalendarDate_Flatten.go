package fixtures

import (
	"time"

	"bitbucket.org/friasdesign/clcity/pkg/gtfs/files"
)

type CalendarDateFlattenTestCase struct {
	Input  files.CalendarDate
	Output []string
}

func TestCalendarDateFlatten() (
	fix CalendarDateFlattenTestCase,
	fixWithout CalendarDateFlattenTestCase,
) {
	x := files.CalendarDate{
		Service:       &files.Service{ID: "SE001"},
		Date:          time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
		ExceptionType: 2,
	}
	fix = CalendarDateFlattenTestCase{
		Input: x,
		Output: []string{
			// service_id
			"SE001",
			// date
			"20000101",
			// exception_type
			"2",
		},
	}
	xWithout := x
	xWithout.Service = nil
	fixWithout = CalendarDateFlattenTestCase{
		Input: xWithout,
		Output: []string{
			// service_id
			"",
			// date
			"20000101",
			// exception_type
			"2",
		},
	}

	return
}
