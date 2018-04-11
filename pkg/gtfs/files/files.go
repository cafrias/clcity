package files

// GTFS files names
const (
	AgencyFileName         = "agency.txt"
	StopsFileName          = "stops.txt"
	RoutesFileName         = "routes.txt"
	TripsFileName          = "trips.txt"
	StopTimesFileName      = "stop_times.txt"
	CalendarFileName       = "calendar.txt"
	CalendarDatesFileName  = "calendar_dates.txt"
	FareAttributesFileName = "fare_attributes.txt"
	FareRulesFileName      = "fare_rules.txt"
	ShapesFileName         = "shapes.txt"
	FrequenciesFileName    = "frequencies.txt"
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

// TripsFileHeaders contains all headers used by 'trips.txt'
var TripsFileHeaders = []string{
	"route_id", "service_id", "trip_id", "trip_headsign", "trip_short_name", "direction_id", "block_id", "shape_id", "wheelchair_accessible", "bikes_allowed",
}

// StopTimesFileHeaders contains all headers used by 'stop_times.txt'
var StopTimesFileHeaders = []string{
	"trip_id", "arrival_time", "departure_time", "stop_id", "stop_sequence", "stop_headsign", "pickup_type", "drop_off_type", "shape_dist_traveled", "timepoint",
}

// CalendarFileHeaders contains all headers used by 'calendar.txt'
var CalendarFileHeaders = []string{
	"service_id", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday", "start_date", "end_date",
}

// CalendarDatesFileHeaders contains all headers used by 'calendar_dates.txt'
var CalendarDatesFileHeaders = []string{
	"service_id", "date", "exception_type",
}

// FareAttributesFileHeaders contains all headers used by 'fare_attributes.txt'
var FareAttributesFileHeaders = []string{
	"fare_id", "price", "currency_type", "payment_method", "transfers", "agency_id", "transfer_duration",
}

// FareRulesFileHeaders contains all headers used by 'fare_attributes.txt'
var FareRulesFileHeaders = []string{
	"fare_id", "route_id", "origin_id", "destination_id", "contains_id",
}

// ShapesFileHeaders contains all headers used by 'shapes.txt'
var ShapesFileHeaders = []string{
	"shape_id", "shape_pt_lat", "shape_pt_lon", "shape_pt_sequence", "shape_dist_traveled",
}

// FrequenciesFileHeaders contains all headers used by 'shapes.txt'
var FrequenciesFileHeaders = []string{
	"trip_id", "start_time", "end_time", "headway_secs", "exact_times",
}
