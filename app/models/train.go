package models

import (
	"aahframework.org/aah.v0"
	"github.com/train-cat/api-train/app/rest"
)

type (
	TrainData struct {
		Code       *string `gorm:"column:code"        json:"code"        validate:"required,len=6,unique=code"`
		Mission    *string `gorm:"column:mission"     json:"mission"     validate:"required,len=4"`
		TerminusID *int    `gorm:"column:terminus_id" json:"terminus_id" validate:"omitempty,foreign_key=station"`
	}

	Train struct {
		Entity
		TrainData
		Terminus *Station `gorm:"ForeignKey:TerminusID" json:"-"`
		rest.Hateoas
	}

	TrainInput TrainData
)

func (i *TrainInput) ToEntity() *Train {
	return &Train{TrainData: TrainData(*i)}
}

// TableName return the table for gorm, used for 'unique' validation
func (i *TrainInput) TableName() string {
	return "train"
}

func (t *Train) GenerateHateoas(ctx *aah.Context) error {
	if t.Terminus != nil {
		if err := t.Terminus.GenerateHateoas(ctx); err != nil {
			return err
		}
	}

	t.Hateoas = rest.Hateoas{
		Embedded: rest.Embedded{
			"terminus": t.Terminus,
		},
		Links: rest.Links{
			"self": rest.Link{
				Href: rest.GenerateURI(ctx, "get_train", *t.Code),
			},
		},
	}

	return nil
}
