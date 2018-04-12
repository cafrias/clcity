package fixtures

import (
	"bitbucket.org/friasdesign/clcity/pkg/gtfs/files"
)

type ShapePointFlattenTestCase struct {
	Input  files.ShapePoint
	Output []string
}

func TestShapePointFlatten() (
	fix ShapePointFlattenTestCase,
	fixWithout ShapePointFlattenTestCase,
) {
	x := files.ShapePoint{
		Shape:        &files.Shape{ID: "001"},
		Lat:          20.21,
		Lon:          21.22,
		PtSequence:   0,
		DistTraveled: 500,
	}
	fix = ShapePointFlattenTestCase{
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
	fixWithout = ShapePointFlattenTestCase{
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
