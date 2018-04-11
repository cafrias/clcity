package fixtures

import (
	"golang.org/x/text/currency"

	"bitbucket.org/friasdesign/clcity/pkg/gtfs/files"
)

type FareFlattenTestCase struct {
	Input  files.Fare
	Output []string
}

func TestFareFlatten() (
	fix FareFlattenTestCase,
	fixWithout FareFlattenTestCase,
) {
	x := files.Fare{
		ID:               "001",
		Price:            25.25,
		CurrencyType:     currency.USD,
		PaymentMethod:    1,
		Transfers:        3,
		Agency:           &files.Agency{ID: "AG001"},
		TransferDuration: 200,
	}
	fix = FareFlattenTestCase{
		Input: x,
		Output: []string{
			// fare_id
			"001",
			// price
			"25.25",
			// currency_type
			"USD",
			// payment_method
			"1",
			// transfers
			"3",
			// agency_id
			"AG001",
			// transfer_duration
			"200",
		},
	}
	xWithout := x
	xWithout.Agency = nil
	xWithout.Transfers = -1
	xWithout.TransferDuration = 0
	fixWithout = FareFlattenTestCase{
		Input: xWithout,
		Output: []string{
			// fare_id
			"001",
			// price
			"25.25",
			// currency_type
			"USD",
			// payment_method
			"1",
			// transfers
			"",
			// agency_id
			"",
			// transfer_duration
			"",
		},
	}

	return
}
