package models

import (
	"aahframework.org/aah.v0"
	"github.com/train-cat/api-train/app/rest"
)

const (
	ActionTypeMessenger = "messenger"
)

var AllActionTypes = []string{ActionTypeMessenger}

type (
	ActionData struct {
		Type   *string `gorm:"column:type"    json:"type" validate:"required,action_type"`
		Data   *Json   `gorm:"column:data"    json:"data" validate:"required"`
		UserID *uint   `gorm:"column:user_id" json:"user_id,omitempty"`
	}

	Action struct {
		Entity
		ActionData
		User *User  `gorm:"ForeignKey:UserID" json:"-"`
		UUID string `gorm:"column:uuid"       json:"-"` // Use for identify user without UserID
		rest.Hateoas
	}

	ActionInput ActionData
)

func (i *ActionInput) ToEntity() *Action {
	e := &Action{ActionData: ActionData(*i)}

	e.SetUUID()

	return e
}

func (i *ActionInput) TableName() string {
	return "action"
}

func (a *Action) SetUUID() string {
	switch *a.Type {
	case ActionTypeMessenger:
		a.UUID = a.Data.Get("messenger_id", "")
	}

	return a.UUID
}

func (a *Action) GenerateHateoas(ctx *aah.Context) error {
	a.Hateoas = rest.Hateoas{
		Embedded: rest.Embedded{},
		Links: rest.Links{
			"self": rest.Link{
			// Href: rest.GenerateURI(ctx, "get_action", a.ID),
			},
			// TODO link list alerts
		},
	}

	if a.User != nil {
		err := a.User.GenerateHateoas(ctx)

		if err != nil {
			return err
		}

		a.Hateoas.Embedded["user"] = a.User
	}

	return nil
}
