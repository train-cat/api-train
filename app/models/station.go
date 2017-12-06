package models

import (
	"aahframework.org/aah.v0"
	"github.com/train-cat/api-train/app/rest"
)

type (
	// StationData is all fields can be send by the client
	StationData struct {
		Name     *string `gorm:"column:name"        json:"name"        validate:"required,min=2,max=255"`
		UIC      *string `gorm:"column:UIC"         json:"uic"         validate:"required,len=8"`
		IsEnable *bool   `gorm:"column:is_enable"   json:"is_enable"   validate:"required"`
	}

	// Station of the SNCF
	Station struct {
		Entity
		StationData
		rest.Hateoas
	}

	// StationInput by the client
	StationInput StationData
)

// ToEntity transform StationInput to Station
func (i *StationInput) ToEntity() *Station {
	return &Station{StationData: StationData(*i)}
}

// GenerateHateoas content
func (s *Station) GenerateHateoas(ctx *aah.Context) error {
	s.Links = rest.Links{
		"self": rest.Link{
			Href: rest.GenerateURI(ctx, "get_station", s.ID),
		},
		"stops": rest.Link{
			Href: rest.GenerateURI(ctx, "cget_station_stops_time", s.ID),
		},
	}

	return nil
}
