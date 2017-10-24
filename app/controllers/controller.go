package controllers

import (
	"net/http"
	"reflect"
	"strings"

	"aahframework.org/aah.v0"
	"gopkg.in/go-playground/validator.v9"
	"github.com/train-sh/api-train/app/rest"
	"github.com/train-sh/api-train/app/validators"
)

type Controller struct {
	*aah.Context
}


func (c *Controller) serverError(err error) bool {
	if err == nil {
		return false
	}

	c.Log().Error(err)
	c.Reply().InternalServerError().JSON(aah.Error{})

	return true
}

func (c *Controller) validatePost(input interface{}) bool {
	errs := validators.Struct(input)

	if errs != nil {
		c.badRequestResponse(errs)
		return false
	}

	return true
}

func (c *Controller) get(s rest.Hateoasable, err error) {
	if c.serverError(err) ||
		c.notFound(s) ||
		c.serverError(c.hateoas(s)) {
		return
	}

	c.Reply().Ok().JSON(s)
}

func (c *Controller) notFound(s interface{}) bool {
	if !reflect.ValueOf(s).IsNil() {
		return false
	}

	c.Reply().NotFound().JSON(aah.Error{
		Code: http.StatusNotFound,
		Message: "Not Found",
	})

	return true
}

func (c *Controller) badRequestResponse(errs error) {
	m := map[string]string{}

	for _, err := range errs.(validator.ValidationErrors) {
		// TODO improve method to get field name
		field := strings.ToLower(err.Field())
		m[field] = err.Translate(nil)
	}

	c.Reply().BadRequest().JSON(m)
}

func (c *Controller) hateoas(s rest.Hateoasable) error {
	return s.GenerateHateoas(c.Context)
}
