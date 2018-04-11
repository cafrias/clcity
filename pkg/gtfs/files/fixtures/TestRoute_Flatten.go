package fixtures

import (
	"image/color"
	"net/url"

	"bitbucket.org/friasdesign/clcity/pkg/gtfs/files"
)

type RouteFlattenTestCase struct {
	Input  files.Route
	Output []string
}

func TestRouteFlatten() (
	fix RouteFlattenTestCase,
	fixWag RouteFlattenTestCase,
) {
	fixURL, _ := url.Parse("https://github.com")
	r := files.Route{
		ID:        "001",
		Agency:    &files.Agency{ID: "AG001"},
		ShortName: "Short",
		LongName:  "Lon",
		Desc:      "desc",
		Type:      2,
		URL:       *fixURL,
		Color: color.RGBA{ // #4286f4
			R: 66,
			G: 134,
			B: 244,
			A: 1,
		},
		TextColor: color.RGBA{ // #2fa020
			R: 47,
			G: 160,
			B: 32,
			A: 1,
		},
		SortOrder: 0,
	}
	rw := r
	rw.Agency = nil
	fix = RouteFlattenTestCase{
		Input: r,
		Output: []string{
			// route_id
			"001",
			// agency_id
			"AG001",
			// route_short_name
			r.ShortName,
			// route_long_name
			r.LongName,
			// route_desc
			r.Desc,
			// route_type
			"2",
			// route_url
			"https://github.com",
			// route_color
			"#4286f4",
			// route_text_color
			"#2fa020",
			// route_sort_order
			"0",
		},
	}
	fixWag = RouteFlattenTestCase{
		Input: rw,
		Output: []string{
			// route_id
			"001",
			// agency_id
			"",
			// route_short_name
			r.ShortName,
			// route_long_name
			r.LongName,
			// route_desc
			r.Desc,
			// route_type
			"2",
			// route_url
			"https://github.com",
			// route_color
			"#4286f4",
			// route_text_color
			"#2fa020",
			// route_sort_order
			"0",
		},
	}

	return
}
