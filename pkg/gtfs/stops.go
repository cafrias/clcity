package gtfs

import "net/url"

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
