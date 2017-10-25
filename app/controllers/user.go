package controllers

import (
	"github.com/train-cat/api-train/app/models"
	"github.com/train-cat/api-train/app/repositories"
)

type UserController struct {
	Controller
}

func (c *UserController) Post(i *models.UserInput) {
	if !c.validatePost(i) {
		return
	}

	u, err := repositories.User.Persist(i)

	if c.serverError(err) || c.serverError(c.hateoas(u)) {
		return
	}

	c.Reply().Created().JSON(u)
}
