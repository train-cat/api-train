package controllers

import (
	"github.com/train-cat/api-train/app/models"
	"github.com/train-cat/api-train/app/repositories"
)

// ActionController regroup all endpoints concern action
type ActionController struct {
	Controller
}

// BeforePost verify if the user is allowed to create an action
func (c *ActionController) BeforePost() {
	c.needRole("bot")
}

// Post create a new action
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
