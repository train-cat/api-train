package controllers

import (
	"github.com/train-cat/api-train/app/filters"
	"github.com/train-cat/api-train/app/repositories"
)

// StopTimeController regroup all endpoints concern stop
type StopTimeController struct {
	Controller
}

// CGetByStation return all stops for one given station
func (c *StopTimeController) CGetByStation(stationID int, f *filters.StopTime) {
	ss, err := repositories.StopTime.FindAllByStation(stationID, f)

	c.get(ss, err)
}

// TODO CGetByTrip
