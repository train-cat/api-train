package models

import (
	"aahframework.org/aah.v0"
	"github.com/train-cat/api-train/app/rest"
)

type (
	// TrainData they can be received by client
	TrainData struct {
		Code       *string `gorm:"column:code"        json:"code"        validate:"required,max=6,unique=code"`
		Mission    *string `gorm:"column:mission"     json:"mission"     validate:"required,len=4"`
		TerminusID *uint   `gorm:"column:terminus_id" json:"terminus_id" validate:"omitempty,foreign_key=station"`
	}

	// Train of the SNCF
	Train struct {
		Entity
		TrainData
		Terminus *Station `gorm:"ForeignKey:TerminusID" json:"-"`
		rest.Hateoas
	}

	// TrainInput by client
	TrainInput TrainData
)

// ToEntity transform TrainInput to Train
func (i *TrainInput) ToEntity() *Train {
	return &Train{TrainData: TrainData(*i)}
}

// TableName for TrainInput, used for 'unique' validation
func (i *TrainInput) TableName() string {
	return "train"
}

// GenerateHateoas content
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
