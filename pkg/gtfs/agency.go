package gtfs

import (
	"net/mail"
	"net/url"
)

var _ FeedFileEntry = &Agency{}

// AgencyID represents the ID for an Agency
type AgencyID string

// Agency represents a single Agency that can be saved on a 'agency.txt' GTFS feed file
type Agency struct {
	ID       AgencyID
	Name     string
	URL      url.URL
	Timezone Timezone
	Lang     LanguageISO6391
	Phone    string
	FareURL  url.URL
	Email    mail.Address
}

// Validate validates the Agency struct is valid as of GTFS specification
func (a *Agency) Validate() (bool, *ErrValidation) {
	ok := true
	err := new(ErrValidation)
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
func (a *Agency) Flatten() []string {
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
