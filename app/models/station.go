package models

import (
	"aahframework.org/aah.v0"
	"github.com/train-cat/api-train/app/rest"
)

type (
	StationData struct {
		Name       *string `gorm:"column:name"        json:"name"        validate:"required,min=2,max=255"`
		UIC        *string `gorm:"column:UIC"         json:"uic"         validate:"required,len=8"`
		IsRealTime *bool   `gorm:"column:is_realtime" json:"is_realtime" validate:"required"`
	}

	Station struct {
		Entity
		StationData
		rest.Hateoas
	}

	StationInput StationData
)

func (i *StationInput) ToEntity() *Station {
	return &Station{StationData: StationData(*i)}
}

func (s *Station) GenerateHateoas(ctx *aah.Context) error {
	s.Links = rest.Links{
		"self": rest.Link{
			Href: rest.GenerateURI(ctx, "get_station", s.ID),
		},
		"stops": rest.Link{
			Href: rest.GenerateURI(ctx, "cget_station_stops", s.ID),
		},
	}

	return nil
}
