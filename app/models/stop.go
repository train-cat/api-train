package models

import (
	"encoding/json"

	"aahframework.org/aah.v0"
	"github.com/train-cat/api-train/app/rest"
)

type (
	// TODO validation
	StopData struct {
		Time   *Schedule `gorm:"column:schedule" json:"schedule" validate:"required"`
		IsWeek *bool     `gorm:"column:is_week"  json:"is_week"`
	}

	Stop struct {
		Entity
		StopData
		StationID uint     `gorm:"column:station_id"    json:"station_id"`
		TrainID   uint     `gorm:"column:train_id"      json:"-"`
		Station   *Station `gorm:"ForeignKey:StationID" json:"-"`
		Train     *Train   `gorm:"ForeignKey:TrainID"   json:"-"`
		rest.Hateoas
	}

	StopOutput struct {
		ID       uint   `json:"id"`
		Schedule string `json:"schedule"`
		OnWeek   bool   `json:"on_week"`
		rest.Hateoas
	}

	StopInput StopData
)

func (i *StopInput) ToEntity() *Stop {
	return &Stop{StopData: StopData(*i)}
}

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

	return nil
}

func (s Stop) MarshalJSON() ([]byte, error) {
	return json.Marshal(StopOutput{
		ID:       s.ID,
		Schedule: s.Time.String(),
		OnWeek:   *s.IsWeek,
		Hateoas:  s.Hateoas,
	})
}
