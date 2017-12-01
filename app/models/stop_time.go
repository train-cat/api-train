package models

import (
	"encoding/json"

	"aahframework.org/aah.v0"
	"github.com/train-cat/api-train/app/rest"
)

type (
	// StopTime represent one train stop at one station in given time
	StopTime struct {
		Entity
		Time      *Schedule `gorm:"column:schedule"      json:"schedule"`
		StationID uint      `gorm:"column:station_id"    json:"station_id"`
		TripID    uint      `gorm:"column:trip_id"       json:"-"`
		Station   *Station  `gorm:"ForeignKey:StationID" json:"-"`
		Trip      *Trip     `gorm:"ForeignKey:TripID"    json:"-"`
		rest.Hateoas
	}

	// StopTimeOutput is json send by the API
	StopTimeOutput struct {
		ID       uint   `json:"id"`
		Schedule string `json:"schedule"`
		rest.Hateoas
	}
)

// GenerateHateoas content
func (s *StopTime) GenerateHateoas(ctx *aah.Context) error {
	s.Hateoas = rest.Hateoas{
		Links: rest.Links{
			"trip": rest.Link{
				Href: rest.GenerateURI(ctx, "get_trip", s.Trip.Code),
			},
			"station": rest.Link{
				Href: rest.GenerateURI(ctx, "get_station", s.StationID),
			},
		},
		Embedded: rest.Embedded{
			"mission": s.Trip.Mission,
			"days":    s.Trip.Calendar.Days,
		},
	}

	if s.Station != nil {
		err := s.Station.GenerateHateoas(ctx)

		if err != nil {
			return err
		}

		s.Hateoas.Embedded["station"] = s.Station
	}

	return nil
}

// MarshalJSON StopTime to StopTimeOutput
func (s StopTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(StopTimeOutput{
		ID:       s.ID,
		Schedule: s.Time.String(),
		Hateoas:  s.Hateoas,
	})
}
