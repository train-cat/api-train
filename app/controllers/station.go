package controllers

import (
	"github.com/train-cat/api-train/app/models"
	"github.com/train-cat/api-train/app/repositories"
	"github.com/train-cat/api-train/app/filters"
)

// StationController regroup all endpoints concern station
type StationController struct {
	Controller
}

// Get return one station
func (c *StationController) Get(stationID uint) {
	s, err := repositories.Station.FindOne(stationID)

	c.get(s, err)
}

// CGet regroup list of stations
func (c *StationController) CGet(f *filters.Station) {
	ss, err := repositories.Station.FindAll(f)

	c.get(ss, err)
}

// BeforePost verify if user is allowed to create a train
func (c *StationController) BeforePost() {
	c.needRole("admin")
}

// Post create a new train
func (c *StationController) Post(i *models.StationInput) {
	if !c.validatePost(i) {
		return
	}

	s, err := repositories.Station.Persist(i)

	if c.serverError(err) || c.serverError(c.hateoas(s)) {
		return
	}

	c.Reply().Created().JSON(s)
}
