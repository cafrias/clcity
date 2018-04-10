package parser

import (
	"encoding/csv"
	"os"
	"path"

	"bitbucket.org/friasdesign/clcity/pkg/gtfs"
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

	// Write agency.txt
	file, err := os.Create(path.Join(p.Path, "agency.txt"))
	if err != nil {
		return err
	}
	defer file.Close()
	writer := csv.NewWriter(file)

	records := [][]string{
		f.Agencies.Headers(),
	}

	for _, agency := range f.Agencies {
		records = append(records, agency.Flatten())
	}

	writer.WriteAll(records)
	writer.Flush()

	err = writer.Error()

	return err
}
