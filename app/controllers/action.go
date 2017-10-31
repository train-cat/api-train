package controllers

import (
	"github.com/train-cat/api-train/app/models"
	"github.com/train-cat/api-train/app/repositories"
)

type ActionController struct {
	Controller
}

func (c *ActionController) BeforePost() {
	c.needRole("bot")
}

func (c *ActionController) Post(i *models.ActionInput) {
	if !c.validatePost(i) {
		return
	}

	a, err := repositories.Action.Persist(i)

	if c.serverError(err) || c.serverError(c.hateoas(a)) {
		return
	}

	c.Reply().Created().JSON(a)
}
