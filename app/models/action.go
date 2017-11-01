package models

import (
	"aahframework.org/aah.v0"
	"github.com/train-cat/api-train/app/rest"
)

// List of actions available
const (
	ActionTypeMessenger = "messenger"
)

// AllActionTypes regroup all types available, used for validation
var AllActionTypes = []string{ActionTypeMessenger}

type (
	// ActionData send by client
	ActionData struct {
		Type   *string `gorm:"column:type"    json:"type" validate:"required,action_type"`
		Data   *JSON   `gorm:"column:data"    json:"data" validate:"required"`
	}

	// Action to perform when alert is triggered
	Action struct {
		Entity
		ActionData
		User *User  `gorm:"ForeignKey:UserID" json:"-"`
		UUID string `gorm:"column:uuid"       json:"-"` // Use for identify user without UserID
		rest.Hateoas
	}

	// ActionInput send by client
	ActionInput ActionData
)

// ToEntity transform ActionInput to Action
func (i *ActionInput) ToEntity() *Action {
	e := &Action{ActionData: ActionData(*i)}

	e.SetUUID()

	return e
}

// TableName for ActionInput, used by unique validation on ActionInput
func (i *ActionInput) TableName() string {
	return "action"
}

// SetUUID for ensure only one action is set per type and user
func (a *Action) SetUUID() string {
	switch *a.Type {
	case ActionTypeMessenger:
		a.UUID = a.Data.Get("messenger_id", "")
	}

	return a.UUID
}

// GenerateHateoas content
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

	return nil
}
