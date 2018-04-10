package gtfs

import (
	"time"

	"github.com/johngb/langreg"
)

// Feed represents a GTFS feed.
type Feed struct {
	files map[string]FeedFile
}

// NewFeed creates a new feed.
func NewFeed() *Feed {
	return &Feed{
		files: make(map[string]FeedFile),
	}
}

// Files returns all files associated with this Feed.
func (f *Feed) Files() map[string]FeedFile {
	return f.files
}

// SetFile set a file for the current feed.
func (f *Feed) SetFile(a FeedFile) {
	f.files[a.FileName()] = a
}

// FeedFile represents a file on a GTFS feed.
type FeedFile interface {
	// Flatten returns a flattened version of this map ready to be passed to CSV parser.
	Flatten() [][]string
	// FileName returns GTFS filename that this interface represents.
	FileName() string
}

// FeedFileEntry represents an entry on any GTFS Feed file. It contains required methods for them to implement
type FeedFileEntry interface {
	Validate() (bool, *ErrValidation)
	Flatten() []string
}

// Parser represents the parser which will read and write GTFS feed files
type Parser interface {
	// Read(path string) error
	Write(f *Feed) error
}

// Timezone represents a timezone string
type Timezone string

// Validate validates a Timezone
func (t Timezone) Validate() bool {
	_, err := time.LoadLocation(string(t))
	if err != nil {
		return false
	}

	return true
}

// LanguageISO6391 represents a language as defined by BCP 47 specification
type LanguageISO6391 string

// Validate validates a LanguageISO6391
func (l LanguageISO6391) Validate() bool {
	return langreg.IsValidLanguageCode(string(l))
}
