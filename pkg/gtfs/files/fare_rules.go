package files

import (
	"bitbucket.org/friasdesign/clcity/pkg/gtfs"
)

var _ gtfs.FeedFile = new(FareRules)
var _ gtfs.FeedFileEntry = &FareRule{}

// FareRules is a map with all FareRules represented on 'fare_rules.txt' file of the GTFS feed.
type FareRules map[FareID][]FareRule

// I used a slice above because I didn't find a way to uniquely identify a single rule, without having to use all fields, and in any case, it's simply pointless.

// FileName returns the name for this GTFS file
func (a FareRules) FileName() string {
	return FareRulesFileName
}

// FileHeaders returns the headers for this GTFS file
func (a FareRules) FileHeaders() []string {
	return FareRulesFileHeaders
}

// FileEntries return all file entries for this GTFS file
func (a FareRules) FileEntries() []gtfs.FeedFileEntry {
	ret := []gtfs.FeedFileEntry{}

	for _, ag := range a {
		for _, y := range ag {
			ret = append(ret, &y)
		}
	}

	return ret
}

// FareRule represents a single FareRule that can be saved on the 'FareRules.txt' GTFS feed file
type FareRule struct {
	Fare        *Fare
	Route       *Route
	Origin      ZoneID
	Destination ZoneID
	Contains    ZoneID
}

// Validate validates the FareRule struct is valid as of GTFS specification
func (a *FareRule) Validate() (bool, *gtfs.ErrValidation) {
	return false, nil
}

// Flatten returns the struct flattened for passing it to CSV parser
func (a *FareRule) Flatten() []string {
	var faID, roID string
	if a.Fare != nil {
		faID = string(a.Fare.ID)
	}
	if a.Route != nil {
		roID = string(a.Route.ID)
	}
	return []string{
		// fare_id
		faID,
		// route_id
		roID,
		// origin_id
		string(a.Origin),
		// destination_id
		string(a.Destination),
		// contains_id
		string(a.Contains),
	}
}
