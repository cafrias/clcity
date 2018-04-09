package gtfs

import (
	"net/url"
	"time"

	"github.com/johngb/langreg"
)

// Feed represents a GTFS feed.
type Feed struct {
	Agencies map[AgencyID]Agency
	Stops    map[StopID]Stop
	Routes   map[RouteID]Route
	// Trips     map[TripID]Trip
	// StopTimes map[TripID]StopTime
	// Shapes    map[ShapeID]Shape
	// FeedInfo  FeedInfo
}

// FeedFileEntry represents an entry on any GTFS Feed file. It contains required methods for them to implement.
type FeedFileEntry interface {
	Validate() bool
	Flatten() []string
}

// StopID represents the ID for a Stop
type StopID string

// Stop represents a single Stop that can be saved on a 'stops.txt' GTFS feed file
type Stop struct {
	ID            StopID
	Code          string
	Name          string
	Desc          string
	Lat           float64
	Lon           float64
	ZoneID        string
	URL           *url.URL
	LocationType  int8
	ParentStation *Stop
	Timezone      Timezone
	Wheelchair    int8
}

type RouteID string

type Route struct {
	ID        RouteID
	Agency    *Agency
	ShortName string
}

// Timezone represents a timezone string
type Timezone string

// Validate validates a Timezone
func (t Timezone) Validate() (bool, error) {
	_, err := time.LoadLocation(string(t))
	if err != nil {
		return false, err
	}

	return true, nil
}

// LanguageISO6391 represents a language as defined by BCP 47 specification
type LanguageISO6391 string

// Validate validates a LanguageISO6391
func (l LanguageISO6391) Validate() bool {
	return langreg.IsValidLanguageCode(string(l))
}
