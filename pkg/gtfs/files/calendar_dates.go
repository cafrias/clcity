package files

import (
	"strconv"
	"time"

	"bitbucket.org/friasdesign/clcity/pkg/gtfs/date"

	"bitbucket.org/friasdesign/clcity/pkg/gtfs"
)

var _ gtfs.FeedFile = new(CalendarDates)
var _ gtfs.FeedFileEntry = &CalendarDate{}

// CalendarDates is a map with all agencies represented on 'agency.txt' file of the GTFS feed.
type CalendarDates map[ServiceID]map[date.Date]CalendarDate

// FileName returns the GTFS filename.
func (a CalendarDates) FileName() string {
	return CalendarDatesFileName
}

// FileHeaders returns the headers for this GTFS file
func (a CalendarDates) FileHeaders() []string {
	return CalendarDatesFileHeaders
}

// FileEntries return all file entries for this GTFS file
func (a CalendarDates) FileEntries() []gtfs.FeedFileEntry {
	ret := []gtfs.FeedFileEntry{}

	for _, ag := range a {
		for _, y := range ag {
			ret = append(ret, &y)
		}
	}

	return ret
}

// CalendarDate represents a single CalendarDate that can be saved on the 'calendar.txt' GTFS feed file
type CalendarDate struct {
	Service       *Service
	Date          time.Time
	ExceptionType int8
}

// Validate validates the CalendarDate struct is valid as of GTFS specification
func (a *CalendarDate) Validate() (bool, *gtfs.ErrValidation) {
	return false, nil
}

// Flatten returns the struct flattened for passing it to CSV parser.
func (a *CalendarDate) Flatten() []string {
	var sID string
	if a.Service != nil {
		sID = string(a.Service.ID)
	}
	return []string{
		// service_id
		sID,
		// date
		string(date.FormatDate(a.Date)),
		// exception_type
		strconv.FormatInt(int64(a.ExceptionType), 10),
	}
}
