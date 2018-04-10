package files

// GTFS files names
const (
	AgencyFileName = "agency.txt"
	StopsFileName  = "stops.txt"
)

// AgencyFileHeaders contains all headers used by 'agency.txt'
var AgencyFileHeaders = []string{
	"agency_id", "agency_name", "agency_url", "agency_timezone", "agency_lang", "agency_phone", "agency_fare_url", "agency_email",
}

// StopsFileHeaders contains all headers used by 'stops.txt'
var StopsFileHeaders = []string{
	"stop_id", "stop_code", "stop_name", "stop_desc", "stop_lat", "stop_lon", "zone_id", "stop_url", "location_type", "parent_station", "stop_timezone", "wheelchair_boarding",
}
