package models

import (
	"aahframework.org/aah.v0"
	"github.com/train-cat/api-train/app/rest"
)

const (
	Sunday = 1 << iota
	Saturday
	Friday
	Thursday
	Wednesday
	Tuesday
	Monday
)

type (
	// Calendar of the trips
	Calendar struct {
		Entity
		Start int     `gorm:"column:start" json:"start"`
		End   int     `gorm:"column:end"   json:"end"`
		Days  Bitmask `gorm:"column:days"  json:"days"`
		rest.Hateoas
	}
)

// GenerateHateoas content
func (c *Calendar) GenerateHateoas(ctx *aah.Context) error {
	c.Hateoas = rest.Hateoas{
		Embedded: rest.Embedded{
			"monday":    c.Days.HasFlag(Monday),
			"tuesday":   c.Days.HasFlag(Tuesday),
			"wednesday": c.Days.HasFlag(Wednesday),
			"thursday":  c.Days.HasFlag(Thursday),
			"friday":    c.Days.HasFlag(Friday),
			"saturday":  c.Days.HasFlag(Saturday),
			"sunday":    c.Days.HasFlag(Sunday),
		},
		Links: rest.Links{
			"self": rest.Link{
			// Href: rest.GenerateURI(ctx, "get_calendar", *c.ID),
			},
		},
	}

	return nil
}
