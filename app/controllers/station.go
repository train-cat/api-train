package controllers

import (
	"github.com/train-sh/api-train/app/models"
	"github.com/train-sh/api-train/app/repositories"
	"github.com/train-sh/api-train/app/filters"
)

type StationController struct {
	Controller
}

func (c *StationController) Get(station_id int) {
	s, err := repositories.Station.FindOne(station_id)

	c.get(s, err)
}

func (c *StationController) CGet(f *filters.Station) {
	ss, err := repositories.Station.FindAll(f)

	c.get(ss, err)
}

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
