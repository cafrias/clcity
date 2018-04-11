package files

import (
	"net/mail"
	"net/url"

	"bitbucket.org/friasdesign/clcity/pkg/gtfs"
)

var _ gtfs.FeedFile = new(Shapes)
var _ gtfs.FeedFileEntry = &Shape{}

// Shapes is a map with all agencies represented on 'agency.txt' file of the GTFS feed.
type Shapes map[ShapeID]Shape

// FileName returns the GTFS filename.
func (a Shapes) FileName() string {
	return ShapesFileName
}

// FileHeaders returns the headers for this GTFS file
func (a Shapes) FileHeaders() []string {
	return ShapesFileHeaders
}

// FileEntries return all file entries for this GTFS file
func (a Shapes) FileEntries() []gtfs.FeedFileEntry {
	ret := []gtfs.FeedFileEntry{}

	for _, ag := range a {
		ret = append(ret, &ag)
	}

	return ret
}

// ShapeID represents the ID for an Shape
type ShapeID string

// Shape represents a single Shape that can be saved on the 'agency.txt' GTFS feed file
type Shape struct {
	ID       ShapeID
	Name     string
	URL      url.URL
	Timezone gtfs.Timezone
	Lang     gtfs.LanguageISO6391
	Phone    string
	FareURL  url.URL
	Email    mail.Address
}

// Validate validates the Shape struct is valid as of GTFS specification
func (a *Shape) Validate() (bool, *gtfs.ErrValidation) {
	ok := true
	err := new(gtfs.ErrValidation)
	err.File = "agency.txt"
	err.Fields = make(map[string]string)

	// TODO: refactor to more reusable code
	if a.Name == "" {
		ok = false
		err.Fields["agency_name"] = ""
	}
	if a.URL.String() == "" {
		ok = false
		err.Fields["agency_url"] = ""
	}
	if a.Timezone.Validate() == false {
		ok = false
		err.Fields["agency_timezone"] = string(a.Timezone)
	}
	if string(a.Lang) != "" && a.Lang.Validate() == false {
		ok = false
		err.Fields["agency_lang"] = string(a.Lang)
	}

	if ok == true {
		return ok, nil
	}

	return ok, err
}

// Flatten returns the struct flattened for passing it to CSV parser.
func (a *Shape) Flatten() []string {
	return []string{
		string(a.ID),
		a.Name,
		a.URL.String(),
		string(a.Timezone),
		string(a.Lang),
		a.Phone,
		a.FareURL.String(),
		a.Email.Address,
	}
}
