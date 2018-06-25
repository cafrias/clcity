package parser

import (
	"encoding/csv"
	"os"
	"path"

	"github.com/cafrias/clcity/pkg/gtfs"
)

var _ gtfs.Parser = &Parser{}

// NewParser instanciates a new parser
func NewParser(p string) *Parser {
	return &Parser{
		Path: p,
	}
}

// Parser represents the GTFS parser that reads and writes to csv files
type Parser struct {
	Path string
}

// Write writes the given feed to csv files
func (p *Parser) Write(f *gtfs.Feed) error {
	done := make(chan struct{})
	defer close(done)

	// Generator. Iterates over each file on the Feed.
	gen := func(done <-chan struct{}, feed *gtfs.Feed) <-chan gtfs.FeedFile {
		outStream := make(chan gtfs.FeedFile)

		go func() {
			defer close(outStream)
			for _, file := range feed.Files() {
				select {
				case <-done:
					return
				case outStream <- file:
				}
			}
		}()

		return outStream
	}
	writer := func(done <-chan struct{}, inStream <-chan gtfs.FeedFile) chan error {
		outStream := make(chan error)

		go func() {
			defer close(outStream)

			for {
				select {
				case <-done:
					return
				case res, ok := <-inStream:
					// inStream channel got closed.
					if ok == false {
						return
					}
					file, err := os.Create(path.Join(p.Path, res.FileName()))
					if err != nil {
						outStream <- err
						return
					}

					writer := csv.NewWriter(file)

					writer.Write(res.FileHeaders())
					for _, entry := range res.FileEntries() {
						writer.Write(entry.Flatten())
					}
					writer.Flush()

					err = writer.Error()
					file.Close()
					if err != nil {
						outStream <- err
						return
					}
				}
			}
		}()

		return outStream
	}

	inStream := gen(done, f)

	for err := range writer(done, inStream) {
		if err != nil {
			return err
		}
	}

	return nil
}
