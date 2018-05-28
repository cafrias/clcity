package files

import (
	"strconv"

	"github.com/friasdesign/clcity/pkg/gtfs"
)

var _ gtfs.FeedFile = new(Transfers)
var _ gtfs.FeedFileEntry = &Transfer{}

// Transfers represents 'transfers.txt' GTFS file
type Transfers map[StopID]map[StopID]*Transfer

// FileName returns the GTFS filename.
func (a Transfers) FileName() string {
	return TransfersFileName
}

// FileHeaders returns the headers for this GTFS file
func (a Transfers) FileHeaders() []string {
	return TransfersFileHeaders
}

// FileEntries return all file entries for this GTFS file
func (a Transfers) FileEntries() []gtfs.FeedFileEntry {
	ret := []gtfs.FeedFileEntry{}

	for _, ag := range a {
		for _, y := range ag {
			ret = append(ret, y)
		}
	}

	return ret
}

// Transfer represents a single Transfer that can be saved on the 'transfers.txt' GTFS feed file
type Transfer struct {
	From            *Stop
	To              *Stop
	Type            int8
	MinTransferTime int
}

func formatMinTransferTime(t int) string {
	if t == 0 {
		return ""
	}

	return strconv.FormatInt(int64(t), 10)
}

// Transfer types
const (
	TransferTypeRecommended = 0
	TransferTypeTimed       = 1
	TransferTypeMinTime     = 2
	TransferTypeNoTransfer  = 3
)

// Validate validates the Transfer struct is valid as of GTFS specification
func (a *Transfer) Validate() (bool, *gtfs.ErrValidation) {
	return false, nil
}

// Flatten returns the struct flattened for passing it to CSV parser.
func (a *Transfer) Flatten() []string {
	var fsID, tsID string
	if a.From != nil {
		fsID = string(a.From.ID)
	}
	if a.To != nil {
		tsID = string(a.To.ID)
	}

	return []string{
		// from_stop_id
		fsID,
		// to_stop_id
		tsID,
		// transfer_type
		strconv.FormatInt(int64(a.Type), 10),
		// min_transfer_time
		formatMinTransferTime(a.MinTransferTime),
	}
}
