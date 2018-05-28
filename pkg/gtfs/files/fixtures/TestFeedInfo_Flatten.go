package fixtures

import (
	"net/url"
	"time"

	"github.com/friasdesign/clcity/pkg/gtfs/files"
	"golang.org/x/text/language"
)

type FeedInfoFlattenTestCase struct {
	Input  files.FeedInfo
	Output []string
}

func TestFeedInfoFlatten() (
	fix FeedInfoFlattenTestCase,
) {
	fURL, _ := url.Parse("https://github.com")
	r := files.FeedInfo{
		PublisherName: "Name",
		PublisherURL:  *fURL,
		Lang:          language.Spanish,
		StartDate:     time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
		EndDate:       time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC),
		Version:       "v2",
	}
	fix = FeedInfoFlattenTestCase{
		Input: r,
		Output: []string{
			// feed_publisher_name
			"Name",
			// feed_publisher_url
			"https://github.com",
			// feed_lang
			"es",
			// feed_start_date
			"20000101",
			// feed_end_date
			"20010101",
			// feed_version
			"v2",
		},
	}

	return
}
