package models

import (
	"aahframework.org/aah.v0"
	"github.com/train-cat/api-train/app/rest"
)

type (
	// AlertData send by client
	AlertData struct {
		ActionID *uint `gorm:"column:action_id" json:"action_id" validate:"foreign_key=action"`
	}

	// Alert is a watcher for one stop (if issue is detected, action is triggered)
	Alert struct {
		Entity
		AlertData
		CodeTrain string   `gorm:"column:code_train"    json:"-"`
		StationID uint     `gorm:"column:station_id"    json:"-"`
		Station   *Station `gorm:"ForeignKey:StationID" json:"-"`
		Action    *Action  `gorm:"ForeignKey:ActionID"  json:"-"`
		rest.Hateoas
	}

	// AlertInput send by client
	AlertInput AlertData
)

// ToEntity transform AlertInput to Alert
func (i *AlertInput) ToEntity() *Alert {
	return &Alert{AlertData: AlertData(*i)}
}

// GenerateHateoas content
func (a *Alert) GenerateHateoas(ctx *aah.Context) error {
	if err := a.Station.GenerateHateoas(ctx); err != nil {
		return err
	}

	if err := a.Action.GenerateHateoas(ctx); err != nil {
		return err
	}

	a.Hateoas = rest.Hateoas{
		Embedded: rest.Embedded{
			"station": a.Station,
			"action":  a.Action,
		},
		Links: rest.Links{
			"self": rest.Link{
			// TODO Href: rest.GenerateURI(ctx, "get_alert"),
			},
			"train": rest.Link{
				Href: rest.GenerateURI(ctx, "get_trip", a.CodeTrain),
			},
		},
	}

	return nil
}
