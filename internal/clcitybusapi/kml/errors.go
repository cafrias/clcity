package kml

// KML parsing errors
const (
	ErrNoLineas = Error("Passed Empresa doesn't have any Linea associated")
)

// Error represents a single error.
type Error string

func (e Error) Error() string { return string(e) }
