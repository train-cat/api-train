package controllers

import (
	"aahframework.org/aah.v0"
	"github.com/train-cat/api-train/app/filters"
	"github.com/train-cat/api-train/app/models"
	"github.com/train-cat/api-train/app/repositories"
)

// StopController regroup all endpoints concern stop
type StopController struct {
	Controller
}

// CGetByStation return all stops for one given station
func (c *StopController) CGetByStation(stationID int, f *filters.Stop) {
	ss, err := repositories.Stop.FindAllByStation(stationID, f)

	c.get(ss, err)
}

// CGetByTrain return all stops for one given train
func (c *StopController) CGetByTrain(code string, f *filters.Pagination) {
	t, err := repositories.Train.FindOneByCode(code)

	if c.serverError(err) || c.notFound(t) {
		return
	}

	ss, err := repositories.Stop.FindAllByTrain(t, f)

	c.get(ss, err)
}

// Head allow to know is specific station is deserve by a train
func (c *StopController) Head(stationID uint, code string) {
	exist, err := repositories.Stop.IsExist(stationID, code)

	if c.serverError(err) {
		return
	}

	if !exist {
		c.notFoundResponse()
		return
	}

	c.Reply().NoContent()
}

// BeforePost verify is user is allowed to link station with a train
func (c *StopController) BeforePost() {
	c.needRole("admin")
}

// Post allow to create new stop between one station and one train
func (c *StopController) Post(stationID uint, code string, i *models.StopInput) {
	s, sErr := repositories.Station.FindOne(stationID)
	t, tErr := repositories.Train.FindOneByCode(code)
	exist, eErr := repositories.Stop.IsExist(stationID, code)

	if c.serverError(sErr) || c.serverError(tErr) || c.serverError(eErr) || // 500
		c.notFound(s) || c.notFound(t) || // 404
		!c.validatePost(i) { // 400
		return
	}

	if exist {
		c.Reply().Conflict().JSON(aah.Error{})
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
