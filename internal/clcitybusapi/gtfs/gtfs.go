package gtfs

import (
	"net/url"
	"strconv"
	"time"

	"github.com/cafrias/clcity/pkg/gtfs/parser"

	"github.com/cafrias/clcity/internal/clcitybusapi"
	"github.com/cafrias/clcity/pkg/gtfs"
	"github.com/cafrias/clcity/pkg/gtfs/files"
)

// Generate generates GTFS feed from data provided by 'Cuando Llega City Bus' API
func Generate(e *clcitybusapi.Empresa, path string) error {
	feed := gtfs.NewFeed()

	// agency.txt
	cityID := files.AgencyID(strconv.Itoa(e.Codigo))
	cityURL, _ := url.Parse(e.URL)
	agencies := files.Agencies{
		cityID: &files.Agency{
			ID:       cityID,
			Name:     e.Nombre,
			URL:      *cityURL,
			Timezone: gtfs.Timezone(e.TZ),
			Lang:     gtfs.LanguageISO6391(e.Lang),
		},
	}
	feed.AddFile(agencies)

	// stops.txt
	stops := make(files.Stops)
	for _, st := range e.Paradas {
		stID := files.StopID(st.Codigo)
		stops[stID] = &files.Stop{
			ID:           stID,
			Code:         st.Codigo,
			Lat:          st.Punto.Lat,
			Lon:          st.Punto.Lon,
			LocationType: files.StopLocationTypeStop,
		}
	}
	feed.AddFile(stops)

	// calendar.txt
	calendar := make(files.Calendar)
	service := &files.Service{
		ID:  "SE001",
		Mon: true,
		Tue: true,
		Wed: true,
		Thu: true,
		Fri: true,
		Sat: true,
		Sun: true,
		// TODO: Change to not fixed dates.
		StartDate: time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
		EndDate:   time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
	}
	calendar[service.ID] = service
	feed.AddFile(calendar)

	// routes.txt, trips.txt, shapes.txt, stop_times.txt
	routes := make(files.Routes)
	trips := make(files.Trips)
	shapes := make(files.Shapes)
	stopTimes := make(files.StopTimes)
	for idx, rt := range e.Lineas {
		rtID := files.RouteID(strconv.Itoa(rt.Codigo))
		route := &files.Route{
			ID:        rtID,
			Agency:    agencies[cityID],
			ShortName: rt.Descripcion,
			LongName:  rt.Descripcion,
			Type:      files.RouteTypeBus,
			Color:     rt.Color,
			SortOrder: idx,
		}

		shID := files.ShapeID(rtID)
		shape := &files.Shape{ID: shID}
		shape.Points = []files.ShapePoint{}

		for seq, rpt := range rt.Recorrido.Puntos {
			spt := files.ShapePoint{
				Shape:      shape,
				Lat:        rpt.Lat,
				Lon:        rpt.Lon,
				PtSequence: seq,
			}
			shape.Points = append(shape.Points, spt)
		}

		// TODO: check if I should create two trips: one outbound, other inbound.
		trID := files.TripID(rtID)
		trip := &files.Trip{
			ID:      trID,
			Service: service,
			Route:   route,
		}

		for idx, par := range rt.Paradas {
			seq := files.StopSequence(idx)
			var arrTime, depTime time.Time
			var tpoint int8
			if idx == 0 {
				// If first
				arrTime = time.Date(2000, 1, 1, 5, 30, 0, 0, time.UTC)
				depTime = arrTime
				tpoint = files.StopTimeTimepointExact
			}

			if stopTimes[trID] == nil {
				stopTimes[trID] = make(map[files.StopSequence]*files.StopTime)
			}

			stopTimes[trID][seq] = &files.StopTime{
				Trip:          trip,
				ArrivalTime:   arrTime,
				DepartureTime: depTime,
				Stop:          stops[files.StopID(par.Identificador)],
				StopSequence:  seq,
				PickupType:    files.StopTimePickTypeDriver,
				DropOffType:   files.StopTimePickTypeDriver,
				Timepoint:     tpoint,
			}
		}

		routes[rtID] = route
		trips[trID] = trip
		shapes[shID] = shape
	}
	feed.AddFile(routes)
	feed.AddFile(trips)
	feed.AddFile(shapes)
	feed.AddFile(stopTimes)

	err := parser.NewParser(path).Write(feed)

	return err
}
