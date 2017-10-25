package models

import (
	"time"

	"encoding/json"

	"aahframework.org/aah.v0"
	"github.com/train-cat/api-train/app/rest"
)

type (
	// Todo validation + StopInput
	StopData struct {
		Time      *time.Time `gorm:"column:time"    json:"schedule"`
		IsWeek    *bool      `gorm:"column:is_week" json:"is_week"`
		StationID *uint      `gorm:"-"              json:"station_id"`
		TrainID   *uint      `gorm:"-"              json:"-"`
	}

	Stop struct {
		Entity
		StopData
		Station *Station `gorm:"ForeignKey:StationID" json:"-"`
		Train   *Train   `gorm:"ForeignKey:TrainID"   json:"-"`
		rest.Hateoas
	}

	StopOutput struct {
		ID       uint   `json:"id"`
		Schedule string `json:"schedule"`
		OnWeek   bool   `json:"on_week"`
		rest.Hateoas
	}
)

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
		Schedule: s.Time.Format("15:04"),
		OnWeek:   *s.IsWeek,
		Hateoas:  s.Hateoas,
	})
}

func (s Stop) TableName() string {
	return "passage"
}
