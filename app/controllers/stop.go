package controllers

import (
	"github.com/train-cat/api-train/app/filters"
	"github.com/train-cat/api-train/app/models"
	"github.com/train-cat/api-train/app/repositories"
)

type StopController struct {
	Controller
}

func (c *StopController) CGetByStation(stationID int, f *filters.Stop) {
	ss, err := repositories.Stop.FindAllByStation(stationID, f)

	c.get(ss, err)
}

func (c *StopController) CGetByTrain(code string, f *filters.Pagination) {
	ss, err := repositories.Stop.FindAllByTrain(code, f)

	c.get(ss, err)
}

func (c *StopController) BeforeLink() {
	c.needRole("admin")
}

func (c *StopController) Link(stationID int, code string, i *models.StopInput) {
	s, sErr := repositories.Station.FindOne(stationID)
	t, tErr := repositories.Train.FindOneByCode(code)

	if c.serverError(sErr) || c.serverError(tErr) || // 500
		c.notFound(s) || c.notFound(t) || // 404
		!c.validatePost(i) { // 400
		return
	}

	stop := i.ToEntity()
	stop.Station = s
	stop.Train = t

	err := repositories.Stop.Persist(stop)

	if c.serverError(err) || c.serverError(c.hateoas(stop)) {
		return
	}

	c.Reply().Created().JSON(stop)
}
