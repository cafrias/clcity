package fixtures

import (
	"image/color"
	"net/mail"
	"net/url"
	"time"

	"golang.org/x/text/language"

	"golang.org/x/text/currency"

	"github.com/cafrias/clcity/pkg/gtfs/date"

	"github.com/cafrias/clcity/pkg/gtfs"
	"github.com/cafrias/clcity/pkg/gtfs/files"
)

func Feed() *gtfs.Feed {
	feed := gtfs.NewFeed()

	// "agency.txt"
	agID := files.AgencyID("AG001")
	fURL, _ := url.Parse("https://github.com")
	agencies := files.Agencies{
		agID: &files.Agency{
			ID: agID,
			Email: mail.Address{
				Address: "pepe@pepe.com",
			},
			URL:      *fURL,
			Name:     "City Bus",
			Timezone: "America/Argentina/Ushuaia",
			Lang:     "es",
		},
	}
	feed.AddFile(agencies)

	// "stops.txt"
	stops := files.Stops{
		"ST001": &files.Stop{
			ID:           "ST001",
			Code:         "ST001",
			Name:         "stop 1",
			Desc:         "between st 1 and st 2",
			Lat:          -53.78286,
			Lon:          -67.69781,
			ZoneID:       "ZO001",
			URL:          *fURL,
			LocationType: files.StopLocationTypeStop,
		},
		"ST002": &files.Stop{
			ID:           "ST002",
			Code:         "ST002",
			Name:         "stop 2",
			Desc:         "between st 1 and st 2",
			Lat:          -53.78307,
			Lon:          -67.70773,
			ZoneID:       "ZO002",
			LocationType: files.StopLocationTypeStop,
		},
	}
	feed.AddFile(stops)

	// "routes.txt"
	routes := files.Routes{
		"RO001": &files.Route{
			ID:        "RO001",
			Agency:    agencies[agID],
			ShortName: "A",
			LongName:  "Linea A",
			Desc:      "Pasa por barrios",
			Type:      files.RouteTypeBus,
			Color: color.RGBA{
				R: 66,
				G: 134,
				B: 244,
			},
			TextColor: color.RGBA{
				R: 0,
				G: 0,
				B: 0,
			},
			SortOrder: 0,
		},
		"RO002": &files.Route{
			ID:        "RO002",
			Agency:    agencies[agID],
			ShortName: "B",
			LongName:  "Linea B",
			Type:      files.RouteTypeBus,
			TextColor: color.RGBA{
				R: 255,
				G: 255,
				B: 255,
			},
		},
	}
	feed.AddFile(routes)

	// "trips.txt"
	se := &files.Service{
		ID:        "SE001",
		Mon:       true,
		Tue:       true,
		Wed:       true,
		Thu:       true,
		Fri:       true,
		Sat:       true,
		Sun:       false,
		StartDate: time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
		EndDate:   time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
	}
	sh := &files.Shape{
		ID: "SH001",
	}
	sh.Points = []files.ShapePoint{
		files.ShapePoint{
			Shape:        sh,
			Lat:          -53.78263,
			Lon:          -67.69747,
			PtSequence:   0,
			DistTraveled: 0,
		},
		files.ShapePoint{
			Shape:        sh,
			Lat:          -53.78357,
			Lon:          -67.70607,
			PtSequence:   1,
			DistTraveled: 800,
		},
	}
	trips := files.Trips{
		"TR001": &files.Trip{
			ID:           "TR001",
			Route:        routes["RO001"],
			Service:      se,
			Headsign:     "head",
			ShortName:    "short",
			DirectionID:  files.TripTravelDirectionIn,
			Shape:        sh,
			Wheelchair:   files.TripWheelchairNoInfo,
			BikesAllowed: files.TripBikesAllowed,
		},
	}
	feed.AddFile(trips)

	// "stop_times.txt"
	stopTimes := files.StopTimes{
		"TR001": {
			0: &files.StopTime{
				Trip:              trips["TR001"],
				ArrivalTime:       time.Date(2000, 1, 1, 20, 5, 0, 0, time.UTC),
				DepartureTime:     time.Date(2000, 1, 1, 20, 10, 0, 0, time.UTC),
				Stop:              stops["ST001"],
				StopSequence:      0,
				DropOffType:       files.StopTimePickTypeDriver,
				ShapeDistTravaled: 150,
				Timepoint:         files.StopTimeTimepointApprox,
			},
			1: &files.StopTime{
				Trip:              trips["TR001"],
				ArrivalTime:       time.Date(2000, 1, 1, 20, 15, 0, 0, time.UTC),
				DepartureTime:     time.Date(2000, 1, 1, 20, 20, 0, 0, time.UTC),
				Stop:              stops["ST002"],
				StopSequence:      1,
				DropOffType:       files.StopTimePickTypeRegular,
				ShapeDistTravaled: 152,
				Timepoint:         files.StopTimeTimepointApprox,
			},
		},
	}
	feed.AddFile(stopTimes)

	// "calendar.txt"
	calendar := files.Calendar{
		se.ID: se,
	}
	feed.AddFile(calendar)

	// "calendar_dates.txt"
	calendarDates := files.CalendarDates{
		se.ID: map[date.Date]*files.CalendarDate{
			"20180101": &files.CalendarDate{
				Service:       se,
				Date:          time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
				ExceptionType: files.CalendarDateExceptionTypeRemoved,
			},
		},
	}
	feed.AddFile(calendarDates)

	// "fare_attributes.txt"
	fareAttr := files.FareAttributes{
		"FA001": &files.Fare{
			ID:               "FA001",
			Price:            8.5,
			CurrencyType:     currency.USD,
			PaymentMethod:    files.FarePaymentMethodOnBoard,
			Transfers:        files.FareTransfersOnce,
			Agency:           agencies[agID],
			TransferDuration: 6500,
		},
	}
	feed.AddFile(fareAttr)

	// "fare_rules.txt"
	fareRules := files.FareRules{
		"FA001": []*files.FareRule{
			&files.FareRule{
				Fare:  fareAttr["FA001"],
				Route: routes["RO001"],
			},
			&files.FareRule{
				Fare:  fareAttr["FA001"],
				Route: routes["RO002"],
			},
		},
	}
	feed.AddFile(fareRules)

	// "shapes.txt"
	shapes := files.Shapes{
		sh.ID: sh,
	}
	feed.AddFile(shapes)

	// "frequencies.txt"
	frequencies := files.Frequencies{
		"TR001": []*files.Frequency{
			&files.Frequency{
				Trip:        trips["TR001"],
				StartTime:   time.Date(2000, 1, 1, 7, 30, 0, 0, time.UTC),
				EndTime:     time.Date(2000, 1, 1, 23, 40, 0, 0, time.UTC),
				HeadwaySecs: 600,
				ExactTimes:  false,
			},
		},
	}
	feed.AddFile(frequencies)

	// "transfers.txt"
	transfers := files.Transfers{
		"ST001": map[files.StopID]*files.Transfer{
			"ST002": &files.Transfer{
				From: stops["ST001"],
				To:   stops["ST002"],
				Type: files.TransferTypeRecommended,
			},
		},
	}
	feed.AddFile(transfers)

	// "feed_info.txt"
	feedInfo := files.FeedInfo{
		PublisherName: "Muni",
		PublisherURL:  *fURL,
		Lang:          language.Spanish,
		StartDate:     time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
		EndDate:       time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		Version:       "v1",
	}
	feed.AddFile(feedInfo)

	return feed
}
