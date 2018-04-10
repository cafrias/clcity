package files

// GTFS files names
const (
	AgencyFileName = "agency.txt"
	StopsFileName  = "stops.txt"
	RoutesFileName = "routes.txt"
)

// AgencyFileHeaders contains all headers used by 'agency.txt'
var AgencyFileHeaders = []string{
	"agency_id", "agency_name", "agency_url", "agency_timezone", "agency_lang", "agency_phone", "agency_fare_url", "agency_email",
}

// StopsFileHeaders contains all headers used by 'stops.txt'
var StopsFileHeaders = []string{
	"stop_id", "stop_code", "stop_name", "stop_desc", "stop_lat", "stop_lon", "zone_id", "stop_url", "location_type", "parent_station", "stop_timezone", "wheelchair_boarding",
}

// RoutesFileHeaders contains all headers used by 'routes.txt'
var RoutesFileHeaders = []string{
	"route_id", "agency_id", "route_short_name", "route_long_name", "route_desc", "route_type", "route_url", "route_color", "route_text_color", "route_sort_order",
}
