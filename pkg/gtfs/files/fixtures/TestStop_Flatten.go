package fixtures

import (
	"net/url"

	"github.com/cafrias/clcity/pkg/gtfs/files"
)

type StopFlattenTestCase struct {
	Input  files.Stop
	Output []string
}

func TestStopFlatten() (
	fix StopFlattenTestCase,
	fixWithout StopFlattenTestCase,
) {
	fixURL, _ := url.Parse("https://github.com")
	stop := files.Stop{
		ID:            "ASD001",
		Code:          "RG001",
		Name:          "Parada",
		Desc:          "Entre esquinas",
		Lat:           -21.65,
		Lon:           -22.65,
		ZoneID:        "ASD1234",
		URL:           *fixURL,
		LocationType:  0,
		ParentStation: &files.Stop{ID: "002"},
		Timezone:      "America/Argentina/Ushuaia",
		Wheelchair:    0,
	}
	stopWithout := stop
	stopWithout.ParentStation = nil
	fix = StopFlattenTestCase{
		Input: stop,
		Output: []string{
			// stop_id
			"ASD001",
			// stop_code
			stop.Code,
			// stop_name
			stop.Name,
			// stop_desc
			stop.Desc,
			// stop_lat
			"-21.65",
			// stop_lon
			"-22.65",
			// zone_id
			"ASD1234",
			// stop_url
			"https://github.com",
			// location_type
			"0",
			// parent_station
			"002",
			// stop_timezone
			"America/Argentina/Ushuaia",
			// wheelchair_boarding
			"0",
		},
	}
	fixWithout = StopFlattenTestCase{
		Input: stopWithout,
		Output: []string{
			// stop_id
			"ASD001",
			// stop_code
			stop.Code,
			// stop_name
			stop.Name,
			// stop_desc
			stop.Desc,
			// stop_lat
			"-21.65",
			// stop_lon
			"-22.65",
			// zone_id
			"ASD1234",
			// stop_url
			"https://github.com",
			// location_type
			"0",
			// parent_station
			"",
			// stop_timezone
			"America/Argentina/Ushuaia",
			// wheelchair_boarding
			"0",
		},
	}

	return
}
