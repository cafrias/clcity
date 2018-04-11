package files

import (
	"time"

	"bitbucket.org/friasdesign/clcity/pkg/gtfs"
	"bitbucket.org/friasdesign/clcity/pkg/gtfs/date"
)

var _ gtfs.FeedFile = new(Calendar)
var _ gtfs.FeedFileEntry = &Service{}

// Calendar is a map with all agencies represented on 'agency.txt' file of the GTFS feed.
type Calendar map[ServiceID]Service

// FileName returns the GTFS filename.
func (a Calendar) FileName() string {
	return CalendarFileName
}

// FileHeaders returns the headers for this GTFS file
func (a Calendar) FileHeaders() []string {
	return CalendarFileHeaders
}

// FileEntries return all file entries for this GTFS file
func (a Calendar) FileEntries() []gtfs.FeedFileEntry {
	ret := []gtfs.FeedFileEntry{}

	for _, ag := range a {
		ret = append(ret, &ag)
	}

	return ret
}

// ServiceID represents the ID for an Service
type ServiceID string

// Service represents a single Service that can be saved on the 'calendar.txt' GTFS feed file
type Service struct {
	ID        ServiceID
	Mon       bool
	Tue       bool
	Wed       bool
	Thu       bool
	Fri       bool
	Sat       bool
	Sun       bool
	StartDate time.Time
	EndDate   time.Time
}

// Validate validates the Service struct is valid as of GTFS specification
func (a *Service) Validate() (bool, *gtfs.ErrValidation) {
	return false, nil
}

// Flatten returns the struct flattened for passing it to CSV parser.
func (a *Service) Flatten() []string {
	return []string{
		// service_id
		string(a.ID),
		// monday
		parseBool(a.Mon),
		// tuesday
		parseBool(a.Tue),
		// wednesday
		parseBool(a.Wed),
		// thursday
		parseBool(a.Thu),
		// friday
		parseBool(a.Fri),
		// saturday
		parseBool(a.Sat),
		// sunday
		parseBool(a.Sun),
		// start_date
		string(date.FormatDate(a.StartDate)),
		// end_date
		string(date.FormatDate(a.EndDate)),
	}
}

func parseBool(x bool) string {
	if x {
		return "1"
	}

	return "0"
}
