package controllers

import (
	"github.com/train-cat/api-train/app/models"
	"github.com/train-cat/api-train/app/repositories"
)

type TrainController struct {
	Controller
}

func (c *TrainController) Post(i *models.TrainInput) {
	if !c.validatePost(i) {
		return
	}

	t, err := repositories.Train.Persist(i)

	if c.serverError(err) || c.serverError(c.hateoas(t)) {
		return
	}

	c.Reply().Created().JSON(t)
}
