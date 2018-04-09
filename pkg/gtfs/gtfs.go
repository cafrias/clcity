package gtfs

import (
	"time"

	"github.com/johngb/langreg"
)

// Feed represents a GTFS feed.
type Feed struct {
	Agencies map[AgencyID]Agency
	// Stops    map[StopID]Stop
	// Routes   map[RouteID]Route
	// Trips     map[TripID]Trip
	// StopTimes map[TripID]StopTime
	// Shapes    map[ShapeID]Shape
	// FeedInfo  FeedInfo
}

// FeedFileEntry represents an entry on any GTFS Feed file. It contains required methods for them to implement
type FeedFileEntry interface {
	Validate() (bool, *ErrValidation)
	Flatten() []string
}

// Parser represents the parser which will read and write GTFS feed files
type Parser interface {
	Read(path string) error
	Write(path string) error
}

// Timezone represents a timezone string
type Timezone string

// Validate validates a Timezone
func (t Timezone) Validate() bool {
	_, err := time.LoadLocation(string(t))
	if err != nil {
		return false
	}

	return true
}

// LanguageISO6391 represents a language as defined by BCP 47 specification
type LanguageISO6391 string

// Validate validates a LanguageISO6391
func (l LanguageISO6391) Validate() bool {
	return langreg.IsValidLanguageCode(string(l))
}
