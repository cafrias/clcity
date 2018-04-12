package files

import (
	"strconv"
	"time"

	"bitbucket.org/friasdesign/clcity/pkg/gtfs"
)

var _ gtfs.FeedFile = new(Frequencies)
var _ gtfs.FeedFileEntry = &Frequency{}

// Frequencies represents 'frequencies.txt' GTFS file
type Frequencies map[TripID][]*Frequency

// FileName returns the GTFS filename.
func (a Frequencies) FileName() string {
	return FrequenciesFileName
}

// FileHeaders returns the headers for this GTFS file
func (a Frequencies) FileHeaders() []string {
	return FrequenciesFileHeaders
}

// FileEntries return all file entries for this GTFS file
func (a Frequencies) FileEntries() []gtfs.FeedFileEntry {
	ret := []gtfs.FeedFileEntry{}

	for _, ag := range a {
		for _, y := range ag {
			ret = append(ret, y)
		}
	}

	return ret
}

// Frequency represents a single Frequency that can be saved on the 'shapes.txt' GTFS feed file
type Frequency struct {
	Trip        *Trip
	StartTime   time.Time
	EndTime     time.Time
	HeadwaySecs int
	ExactTimes  bool
}

// Validate validates the Frequency struct is valid as of GTFS specification
func (a *Frequency) Validate() (bool, *gtfs.ErrValidation) {
	return false, nil
}

// Flatten returns the struct flattened for passing it to CSV parser.
func (a *Frequency) Flatten() []string {
	var trID string
	if a.Trip != nil {
		trID = string(a.Trip.ID)
	}
	return []string{
		// trip_id
		trID,
		// start_time
		formatTime(a.StartTime),
		// end_time
		formatTime(a.EndTime),
		// headway_secs
		strconv.FormatInt(int64(a.HeadwaySecs), 10),
		// exact_times
		formatBool(a.ExactTimes),
	}
}

func formatBool(a bool) string {
	if a {
		return "1"
	}

	return "0"
}
