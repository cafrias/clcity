package fixtures

import (
	"bitbucket.org/friasdesign/clcity/pkg/gtfs/files"
)

type TransferFlattenTestCase struct {
	Input  files.Transfer
	Output []string
}

func TestTransferFlatten() (
	fix TransferFlattenTestCase,
	fixW TransferFlattenTestCase,
) {
	r := files.Transfer{
		From:            &files.Stop{ID: "ST001"},
		To:              &files.Stop{ID: "ST002"},
		Type:            files.TransferTypeMinTime,
		MinTransferTime: 600,
	}
	fix = TransferFlattenTestCase{
		Input: r,
		Output: []string{
			// from_stop_id
			"ST001",
			// to_stop_id
			"ST002",
			// transfer_type
			"2",
			// min_transfer_time
			"600",
		},
	}
	rw := r
	rw.From = nil
	rw.To = nil
	rw.MinTransferTime = 0
	fixW = TransferFlattenTestCase{
		Input: rw,
		Output: []string{
			// from_stop_id
			"",
			// to_stop_id
			"",
			// transfer_type
			"2",
			// min_transfer_time
			"",
		},
	}

	return
}
