package files

import (
	"strconv"

	"bitbucket.org/friasdesign/clcity/pkg/gtfs"
	"golang.org/x/text/currency"
)

var _ gtfs.FeedFile = new(FareAttributes)
var _ gtfs.FeedFileEntry = &Fare{}

// FareAttributes is a map with all FareAttributes represented on 'FareAttributes.txt' file of the GTFS feed
type FareAttributes map[FareID]*Fare

// FileName returns the name for this GTFS file
func (a FareAttributes) FileName() string {
	return FareAttributesFileName
}

// FileHeaders returns the headers for this GTFS file
func (a FareAttributes) FileHeaders() []string {
	return FareAttributesFileHeaders
}

// FileEntries return all file entries for this GTFS file
func (a FareAttributes) FileEntries() []gtfs.FeedFileEntry {
	ret := []gtfs.FeedFileEntry{}

	for _, ag := range a {
		ret = append(ret, ag)
	}

	return ret
}

// FareID represents the ID for an Fare
type FareID string

// Fare represents a single Fare that can be saved on the 'FareAttributes.txt' GTFS feed file
type Fare struct {
	ID               FareID
	Price            float64
	CurrencyType     currency.Unit
	PaymentMethod    int8
	Transfers        int8
	Agency           *Agency
	TransferDuration int
}

// Payment methods
const (
	FarePaymentMethodOnBoard = iota
	FarePaymentMethodBefore
)

// Transfers
const (
	FareTransfersNotPermitted = 0
	FareTransfersOnce         = 1
	FareTransfersTwice        = 2
	FareTransfersUnlimited    = -1 // Maps to empty string on GTFS file
)

func formatTransfers(t int8) string {
	if t == -1 {
		return ""
	}

	return strconv.FormatInt(int64(t), 10)
}

func formatTransferDuration(t int) string {
	if t == 0 {
		return ""
	}

	return strconv.FormatInt(int64(t), 10)
}

// Validate validates the Fare struct is valid as of GTFS specification
func (a *Fare) Validate() (bool, *gtfs.ErrValidation) {
	return false, nil
}

// Flatten returns the struct flattened for passing it to CSV parser
func (a *Fare) Flatten() []string {
	var agID string
	if a.Agency != nil {
		agID = string(a.Agency.ID)
	}
	return []string{
		// fare_id
		string(a.ID),
		// price
		strconv.FormatFloat(a.Price, 'f', -1, 64),
		// currency_type
		a.CurrencyType.String(),
		// payment_method
		strconv.FormatInt(int64(a.PaymentMethod), 10),
		// transfers
		formatTransfers(a.Transfers),
		// agency_id
		agID,
		// transfer_duration
		formatTransferDuration(a.TransferDuration),
	}
}
