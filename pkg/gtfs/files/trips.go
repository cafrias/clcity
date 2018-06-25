package files

import (
	"strconv"

	"github.com/cafrias/clcity/pkg/gtfs"
)

var _ gtfs.FeedFile = new(Trips)
var _ gtfs.FeedFileEntry = &Trip{}

// Trips is a map with all agencies represented on 'agency.txt' file of the GTFS feed.
type Trips map[TripID]*Trip

// FileName returns the GTFS filename.
func (a Trips) FileName() string {
	return TripsFileName
}

// FileHeaders returns the headers for this GTFS file
func (a Trips) FileHeaders() []string {
	return TripsFileHeaders
}

// FileEntries return all file entries for this GTFS file
func (a Trips) FileEntries() []gtfs.FeedFileEntry {
	ret := []gtfs.FeedFileEntry{}

	for _, ag := range a {
		ret = append(ret, ag)
	}

	return ret
}

// TripID represents the ID for an Trip
type TripID string

// Trip represents a single Trip that can be saved on the 'trips.txt' GTFS feed file
type Trip struct {
	Route        *Route
	Service      *Service
	ID           TripID
	Headsign     string
	ShortName    string
	DirectionID  int8
	BlockID      string
	Shape        *Shape
	Wheelchair   int8
	BikesAllowed int8
}

// Trip travel direction, can be outbound or inbound
const (
	TripTravelDirectionNone = -1
	TripTravelDirectionOut  = 0
	TripTravelDirectionIn   = 1
)

// Wheelchair information for a Trip
const (
	TripWheelchairNoInfo = iota
	TripWheelchairAllowed
	TripWheelchairNotAllowed
)

// Bike information for a Trip
const (
	TripBikesNoInfo = iota
	TripBikesAllowed
	TripBikesNotAllowed
)

// Validate validates the Trip struct is valid as of GTFS specification
func (a *Trip) Validate() (bool, *gtfs.ErrValidation) {
	return false, nil
}

// Flatten returns the struct flattened for passing it to CSV parser.
func (a *Trip) Flatten() []string {
	var roID, seID, shID string
	if a.Route != nil {
		roID = string(a.Route.ID)
	}
	if a.Service != nil {
		seID = string(a.Service.ID)
	}
	if a.Shape != nil {
		shID = string(a.Shape.ID)
	}
	return []string{
		// route_id
		roID,
		// service_id
		seID,
		// trip_id
		string(a.ID),
		// trip_headsign
		a.Headsign,
		// trip_short_name
		a.ShortName,
		// direction_id
		strconv.FormatInt(int64(a.DirectionID), 10),
		// block_id
		a.BlockID,
		// shape_id
		shID,
		// wheelchair_accessible
		strconv.FormatInt(int64(a.Wheelchair), 10),
		// bikes_allowed
		strconv.FormatInt(int64(a.BikesAllowed), 10),
	}
}
