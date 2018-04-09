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
	URL      *url.URL
	Timezone Timezone
	Lang     LanguageISO6391
	Phone    string
	FareURL  *url.URL
	Email    *mail.Address
}

func (a *Agency) Validate() (bool, ErrValidation) {
	err := ErrValidation{
		File: "agency.txt",
	}

	if a.ID == "" {
		err.Fields["agency_id"] = "(empty)"
	}

}
