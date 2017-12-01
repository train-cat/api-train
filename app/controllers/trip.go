package controllers

import "github.com/train-cat/api-train/app/repositories"

// TripController regroup all endpoints for train
type TripController struct {
	Controller
}

// Get return one trip
func (c *TripController) Get(id int) {
	t, err := repositories.Trip.Find(id)

	c.get(t, err)
}
