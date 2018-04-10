package files

import (
	"bytes"
	"image/color"
	"net/url"
	"strconv"

	"bitbucket.org/friasdesign/clcity/pkg/gtfs"
)

var _ gtfs.FeedFile = new(Routes)
var _ gtfs.FeedFileEntry = &Route{}

// Routes is a map with all Routes represented on 'Routes.txt' file of the GTFS feed
type Routes map[RouteID]Route

// FileName returns the name for this GTFS file
func (a Routes) FileName() string {
	return RoutesFileName
}

// FileHeaders returns the headers for this GTFS file
func (a Routes) FileHeaders() []string {
	return RoutesFileHeaders
}

// FileEntries return all file entries for this GTFS file
func (a Routes) FileEntries() []gtfs.FeedFileEntry {
	ret := []gtfs.FeedFileEntry{}

	for _, ag := range a {
		ret = append(ret, &ag)
	}

	return ret
}

// RouteID represents the ID for an Route
type RouteID string

// Route represents a single Route that can be saved on the 'Routes.txt' GTFS feed file
type Route struct {
	ID        RouteID
	Agency    *Agency
	ShortName string
	LongName  string
	Desc      string
	Type      int8
	URL       url.URL
	Color     color.RGBA
	TextColor color.RGBA
	SortOrder int
}

// Validate validates the Route struct is valid as of GTFS specification
func (a *Route) Validate() (bool, *gtfs.ErrValidation) {
	return false, nil
}

// Flatten returns the struct flattened for passing it to CSV parser
func (a *Route) Flatten() []string {
	var agID string
	if a.Agency != nil {
		agID = string(a.Agency.ID)
	}
	return []string{
		// route_id
		string(a.ID),
		// agency_id
		agID,
		// route_short_name
		a.ShortName,
		// route_long_name
		a.LongName,
		// route_desc
		a.Desc,
		// route_type
		strconv.FormatInt(int64(a.Type), 10),
		// route_url
		a.URL.String(),
		// route_color
		rgbaToHex(a.Color),
		// route_text_color
		rgbaToHex(a.TextColor),
		// route_sort_order
		strconv.FormatInt(int64(a.SortOrder), 10),
	}
}

func rgbaToHex(c color.RGBA) string {
	buf := bytes.NewBufferString("#")

	// Write red
	rHex := strconv.FormatInt(int64(c.R), 16)
	buf.WriteString(rHex)

	// Write green
	gHex := strconv.FormatInt(int64(c.G), 16)
	buf.WriteString(gHex)

	// Write blue
	bHex := strconv.FormatInt(int64(c.B), 16)
	buf.WriteString(bHex)

	return buf.String()
}
