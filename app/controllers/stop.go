package controllers

import (
	"github.com/train-cat/api-train/app/filters"
	"github.com/train-cat/api-train/app/repositories"
)

type StopController struct {
	Controller
}

func (c *StopController) CGet(stationID int, f *filters.Stop) {
	ss, err := repositories.Stop.FindAllByStation(stationID, f)

	c.get(ss, err)
}
