package files

import (
	"net/url"
	"strconv"

	"github.com/cafrias/clcity/pkg/gtfs"
)

var _ gtfs.FeedFile = new(Stops)
var _ gtfs.FeedFileEntry = &Stop{}

// Stops is a map with all Stops represented on 'stops.txt' file of the GTFS feed
type Stops map[StopID]*Stop

// FileName returns the name for this GTFS file
func (a Stops) FileName() string {
	return StopsFileName
}

// FileHeaders returns the headers for this GTFS file
func (a Stops) FileHeaders() []string {
	return StopsFileHeaders
}

// FileEntries return all file entries for this GTFS file
func (a Stops) FileEntries() []gtfs.FeedFileEntry {
	ret := []gtfs.FeedFileEntry{}

	for _, ag := range a {
		ret = append(ret, ag)
	}

	return ret
}

// StopID represents the ID for an Stop
type StopID string

// ZoneID defines a fare zone for a Stop
type ZoneID string

// Stop represents a single Stop that can be saved on the 'stops.txt' GTFS feed file
type Stop struct {
	ID            StopID
	Code          string
	Name          string
	Desc          string
	Lat           float64
	Lon           float64
	ZoneID        ZoneID
	URL           url.URL
	LocationType  int8
	ParentStation *Stop
	Timezone      gtfs.Timezone
	Wheelchair    int8
}

// LocationTypes for Stop
const (
	StopLocationTypeStop            = 0
	StopLocationTypeStation         = 1
	StopLocationtypeStationEntrance = 2
)

// Validate validates the Stop struct is valid as of GTFS specification
func (a *Stop) Validate() (bool, *gtfs.ErrValidation) {
	return false, nil
}

// Flatten returns the struct flattened for passing it to CSV parser
func (a *Stop) Flatten() []string {
	var parentID string
	if a.ParentStation != nil {
		parentID = string(a.ParentStation.ID)
	}

	return []string{
		// stop_id
		string(a.ID),
		// stop_code
		a.Code,
		// stop_name
		a.Name,
		// stop_desc
		a.Desc,
		// stop_lat
		strconv.FormatFloat(a.Lat, 'f', -1, 64),
		// stop_lon
		strconv.FormatFloat(a.Lon, 'f', -1, 64),
		// zone_id
		string(a.ZoneID),
		// stop_url
		a.URL.String(),
		// location_type
		strconv.FormatInt(int64(a.LocationType), 10),
		// parent_station
		parentID,
		// stop_timezone
		string(a.Timezone),
		// wheelchair_boarding
		strconv.FormatInt(int64(a.Wheelchair), 10),
	}
}
