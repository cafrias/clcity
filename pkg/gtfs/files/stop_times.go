package files

import (
	"fmt"
	"strconv"
	"time"

	"bitbucket.org/friasdesign/clcity/pkg/gtfs"
)

var _ gtfs.FeedFile = new(StopTimes)
var _ gtfs.FeedFileEntry = &StopTime{}

// StopTimes is a map with all agencies represented on 'agency.txt' file of the GTFS feed.
type StopTimes map[TripID]map[StopSequence]StopTime

// FileName returns the GTFS filename.
func (a StopTimes) FileName() string {
	return StopTimesFileName
}

// FileHeaders returns the headers for this GTFS file
func (a StopTimes) FileHeaders() []string {
	return StopTimesFileHeaders
}

// FileEntries return all file entries for this GTFS file
func (a StopTimes) FileEntries() []gtfs.FeedFileEntry {
	ret := []gtfs.FeedFileEntry{}

	for _, ag := range a {
		for _, y := range ag {
			ret = append(ret, &y)
		}
	}

	return ret
}

// StopSequence represents the sequence of a stop during a given trip
type StopSequence int

// StopTime represents a single StopTime that can be saved on the 'stop_times.txt' GTFS feed file
type StopTime struct {
	Trip              *Trip
	ArrivalTime       time.Time
	DepartureTime     time.Time
	Stop              *Stop
	StopSequence      StopSequence
	StopHeadsign      string
	PickupType        int8
	DropOffType       int8
	ShapeDistTravaled float64
	Timepoint         int8
}

// Validate validates the StopTime struct is valid as of GTFS specification
func (a *StopTime) Validate() (bool, *gtfs.ErrValidation) {
	return false, nil
}

// Flatten returns the struct flattened for passing it to CSV parser.
func (a *StopTime) Flatten() []string {
	var trID, stID string
	if a.Trip != nil {
		trID = string(a.Trip.ID)
	}
	if a.Stop != nil {
		stID = string(a.Stop.ID)
	}
	return []string{
		// trip_id
		trID,
		// arrival_time
		formatTime(a.ArrivalTime),
		// departure_time
		formatTime(a.DepartureTime),
		// stop_id
		stID,
		// stop_sequence
		strconv.FormatInt(int64(a.StopSequence), 10),
		// stop_headsign
		a.StopHeadsign,
		// pickup_type
		strconv.FormatInt(int64(a.PickupType), 10),
		// drop_off_type
		strconv.FormatInt(int64(a.DropOffType), 10),
		// shape_dist_traveled
		strconv.FormatFloat(a.ShapeDistTravaled, 'f', -1, 64),
		// timepoint
		strconv.FormatInt(int64(a.Timepoint), 10),
	}
}

func formatTime(t time.Time) string {
	return fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second())
}
