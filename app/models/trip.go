package models

import (
	"aahframework.org/aah.v0"
	"github.com/train-cat/api-train/app/rest"
)

type (
	// Trip of the SNCF
	Trip struct {
		Entity
		Code       string    `gorm:"column:code"           json:"code"`
		Mission    string    `gorm:"column:mission"        json:"mission"`
		CalendarID uint      `gorm:"column:calendar_id"    json:"-"`
		Calendar   *Calendar `gorm:"ForeignKey:CalendarID" json:"-"`
		rest.Hateoas
	}
)

// GenerateHateoas content
func (t *Trip) GenerateHateoas(ctx *aah.Context) error {
	if t.Calendar != nil {
		if err := t.Calendar.GenerateHateoas(ctx); err != nil {
			return err
		}
	}

	t.Hateoas = rest.Hateoas{
		Embedded: rest.Embedded{
			"calendar": t.Calendar,
		},
		Links: rest.Links{
			"self": rest.Link{
				Href: rest.GenerateURI(ctx, "get_trip", t.ID),
			},
		},
	}

	return nil
}
