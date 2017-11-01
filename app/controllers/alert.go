package controllers

import (
	"github.com/train-cat/api-train/app/models"
	"github.com/train-cat/api-train/app/repositories"
)

// AlertController regroup all endpoints concern alert
type AlertController struct {
	Controller
}


// BeforePost verify if user is allowed to create an alert
func (c *AlertController) BeforePost() {
	c.needRole("bot")
}

// Post create a new alert
func (c *AlertController) Post(stationID int, code string, i *models.AlertInput) {
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
