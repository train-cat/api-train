package models

import (
	"encoding/json"

	"aahframework.org/aah.v0"
	"github.com/train-cat/api-train/app/rest"
)

type (
	// StopData they can be received by client
	StopData struct {
		Time   *Schedule `gorm:"column:schedule" json:"schedule" validate:"required"`
		IsWeek *bool     `gorm:"column:is_week"  json:"is_week"`
	}

	// Stop represent one train stop at one station in given time
	Stop struct {
		Entity
		StopData
		StationID uint     `gorm:"column:station_id"    json:"station_id"`
		TrainID   uint     `gorm:"column:train_id"      json:"-"`
		Station   *Station `gorm:"ForeignKey:StationID" json:"-"`
		Train     *Train   `gorm:"ForeignKey:TrainID"   json:"-"`
		rest.Hateoas
	}

	// StopOutput is json send by the API
	StopOutput struct {
		ID       uint   `json:"id"`
		Schedule string `json:"schedule"`
		OnWeek   bool   `json:"on_week"`
		rest.Hateoas
	}

	// StopInput by the client
	StopInput StopData
)

// ToEntity transform StopInput to Stop
func (i *StopInput) ToEntity() *Stop {
	return &Stop{StopData: StopData(*i)}
}

// GenerateHateoas content
func (s *Stop) GenerateHateoas(ctx *aah.Context) error {
	if s.Train.Terminus != nil {
		err := s.Train.Terminus.GenerateHateoas(ctx)

		if err != nil {
			return err
		}
	}

	s.Hateoas = rest.Hateoas{
		Links: rest.Links{
			"train": rest.Link{
				Href: rest.GenerateURI(ctx, "get_train", *s.Train.Code),
			},
			"station": rest.Link{
				Href: rest.GenerateURI(ctx, "get_station", s.StationID),
			},
		},
		Embedded: rest.Embedded{
			"mission":  s.Train.Mission,
			"terminus": s.Train.Terminus,
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

// MarshalJSON Stop to StopOutput
func (s Stop) MarshalJSON() ([]byte, error) {
	return json.Marshal(StopOutput{
		ID:       s.ID,
		Schedule: s.Time.String(),
		OnWeek:   *s.IsWeek,
		Hateoas:  s.Hateoas,
	})
}
