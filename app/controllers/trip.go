package controllers

import "github.com/train-cat/api-train/app/repositories"

// TripController regroup all endpoints for train
type TripController struct {
	Controller
}

// Get return one trip
func (c *TripController) Get(tripID int) {
	t, err := repositories.Trip.Find(tripID)

	c.get(t, err)
}

// GetTerminus return terminus of the trip
func (c *TripController) GetTerminus(tripID uint) {
	s, err := repositories.Station.FindTerminus(tripID)

	c.get(s, err)
}
