package controllers

import (
	"github.com/train-cat/api-train/app/filters"
	"github.com/train-cat/api-train/app/models"
	"github.com/train-cat/api-train/app/repositories"
)

// AlertController regroup all endpoints concern alert
type AlertController struct {
	Controller
}

// BeforeCGet verify user is allowed to get all alerts
func (c *AlertController) BeforeCGet() {
	c.needRole("notifier")
}

// CGet returns alerts filtered
func (c *AlertController) CGet(f *filters.Alert) {
	as, err := repositories.Alert.FindAll(f)

	c.get(as, err)
}

// BeforePost verify if user is allowed to create an alert
func (c *AlertController) BeforePost() {
	c.needRole("bot")
}

// Post create a new alert
func (c *AlertController) Post(stationID int, code string, i *models.AlertInput) {
	// TODO add uniq constraint (stationID + code + i.actionID)

	s, err := repositories.Stop.FindOneByStationAndCode(stationID, code)

	if c.notFound(s) || c.serverError(err) || !c.validatePost(i) {
		return
	}

	a, err := repositories.Alert.Persist(s, i)

	if c.serverError(err) || c.serverError(c.hateoas(a)) {
		return
	}

	c.Reply().Created().JSON(a)
}
