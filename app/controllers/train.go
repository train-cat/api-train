package controllers

import (
	"github.com/train-cat/api-train/app/models"
	"github.com/train-cat/api-train/app/repositories"
)

// TrainController regroup all endpoints for train
type TrainController struct {
	Controller
}

// Head return 204 if train exist
func (c *TrainController) Head(code string) {
	exist := repositories.Train.IsExist(code)

	if !exist {
		c.notFoundResponse()
		return
	}

	c.Reply().NoContent()
}

// Get return one train
func (c *TrainController) Get(code string) {
	t, err := repositories.Train.FindOneByCode(code)

	c.get(t, err)
}

// BeforePost assure user has role 'admin' before execute Post method
func (c *TrainController) BeforePost() {
	c.needRole("admin")
}

// Post create new train
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
