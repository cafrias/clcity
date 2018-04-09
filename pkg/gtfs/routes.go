package gtfs

type RouteID string

type Route struct {
	ID        RouteID
	Agency    *Agency
	ShortName string
}
