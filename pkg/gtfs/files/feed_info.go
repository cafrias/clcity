package files

import (
	"net/url"
	"time"

	"github.com/friasdesign/clcity/pkg/gtfs/date"

	"golang.org/x/text/language"

	"github.com/friasdesign/clcity/pkg/gtfs"
)

var _ gtfs.FeedFile = new(FeedInfo)
var _ gtfs.FeedFileEntry = &FeedInfo{}

// FeedInfo represents 'feed_info.txt'
type FeedInfo struct {
	PublisherName string
	PublisherURL  url.URL
	Lang          language.Tag
	StartDate     time.Time
	EndDate       time.Time
	Version       string
}

// FileName returns the GTFS filename.
func (a FeedInfo) FileName() string {
	return FeedInfoFileName
}

// FileHeaders returns the headers for this GTFS file
func (a FeedInfo) FileHeaders() []string {
	return FeedInfoFileHeaders
}

// FileEntries return all file entries for this GTFS file
func (a FeedInfo) FileEntries() []gtfs.FeedFileEntry {
	return []gtfs.FeedFileEntry{
		&a,
	}
}

// Validate validates the FeedInfo struct is valid as of GTFS specification
func (a FeedInfo) Validate() (bool, *gtfs.ErrValidation) {
	return false, nil
}

// Flatten returns the struct flattened for passing it to CSV parser.
func (a FeedInfo) Flatten() []string {
	return []string{
		// feed_publisher_name
		a.PublisherName,
		// feed_publisher_url
		a.PublisherURL.String(),
		// feed_lang
		a.Lang.String(),
		// feed_start_date
		string(date.FormatDate(a.StartDate)),
		// feed_end_date
		string(date.FormatDate(a.EndDate)),
		// feed_version
		a.Version,
	}
}
