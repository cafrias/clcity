package fixtures

import (
	"bitbucket.org/friasdesign/clcity/pkg/gtfs/files"
)

type ShapeFlattenTestCase struct {
	Input  files.Shape
	Output []string
}

func TestShapeFlatten() (
	fix ShapeFlattenTestCase,
	fixWithout ShapeFlattenTestCase,
) {
	x := files.Shape{
		ID:           "001",
		Lat:          20.21,
		Lon:          21.22,
		PtSequence:   0,
		DistTraveled: 500,
	}
	fix = ShapeFlattenTestCase{
		Input: x,
		Output: []string{
			// shape_id
			"001",
			// shape_pt_lat
			"20.21",
			// shape_pt_lon
			"21.22",
			// shape_pt_sequence
			"0",
			// shape_dist_traveled
			"500",
		},
	}
	xWithout := x
	xWithout.DistTraveled = -1
	fixWithout = ShapeFlattenTestCase{
		Input: xWithout,
		Output: []string{
			// shape_id
			"001",
			// shape_pt_lat
			"20.21",
			// shape_pt_lon
			"21.22",
			// shape_pt_sequence
			"0",
			// shape_dist_traveled
			"",
		},
	}

	return
}
