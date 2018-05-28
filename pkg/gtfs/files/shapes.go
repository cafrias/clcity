package files

import (
	"strconv"

	"github.com/friasdesign/clcity/pkg/gtfs"
)

var _ gtfs.FeedFile = new(Shapes)
var _ gtfs.FeedFileEntry = &ShapePoint{}

// Shapes represents 'shapes.txt' GTFS files
type Shapes map[ShapeID]*Shape

// FileName returns the GTFS filename.
func (a Shapes) FileName() string {
	return ShapesFileName
}

// FileHeaders returns the headers for this GTFS file
func (a Shapes) FileHeaders() []string {
	return ShapesFileHeaders
}

// FileEntries return all file entries for this GTFS file
func (a Shapes) FileEntries() []gtfs.FeedFileEntry {
	ret := []gtfs.FeedFileEntry{}

	for _, ag := range a {
		for _, y := range ag.Points {
			// If I pass pointer to 'y' directly it's gonna change the value each iteration
			cl := y
			ret = append(ret, &cl)
		}
	}

	return ret
}

// ShapeID represents the ID for an Shape
type ShapeID string

// Shape represents a set of ShapePoints with same ID
type Shape struct {
	ID     ShapeID
	Points []ShapePoint
}

// ShapePoint represents a point that is part of a Shape
type ShapePoint struct {
	Shape        *Shape
	Lat          float64
	Lon          float64
	PtSequence   int
	DistTraveled float64
}

func formatDistTraveled(d float64) string {
	if d == -1 {
		return ""
	}

	return strconv.FormatFloat(d, 'f', -1, 64)
}

// Validate validates the ShapePoint struct is valid as of GTFS specification
func (a *ShapePoint) Validate() (bool, *gtfs.ErrValidation) {
	return false, nil
}

// Flatten returns the struct flattened for passing it to CSV parser.
func (a *ShapePoint) Flatten() []string {
	return []string{
		// shape_id
		string(a.Shape.ID),
		// shape_pt_lat
		strconv.FormatFloat(a.Lat, 'f', -1, 64),
		// shape_pt_lon
		strconv.FormatFloat(a.Lon, 'f', -1, 64),
		// shape_pt_sequence
		strconv.FormatInt(int64(a.PtSequence), 10),
		// shape_dist_traveled
		formatDistTraveled(a.DistTraveled),
	}
}
