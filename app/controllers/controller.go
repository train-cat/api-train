package controllers

import (
	"net/http"
	"reflect"
	"strings"

	"aahframework.org/aah.v0"
	"gopkg.in/go-playground/validator.v9"
	"github.com/train-cat/api-train/app/rest"
	"github.com/train-cat/api-train/app/validators"
)

// Controller extend aah.Context
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

func (c *Controller) needRole(role string) {
	abort := false

	if c.Subject() == nil {
		c.unauthorizedResponse()
		abort = true
	}

	if !c.Subject().HasRole(role) {
		c.forbiddenResponse()
		abort = true
	}

	if abort {
		c.Abort()
	}
}


func (c *Controller) notFound(s interface{}) bool {
	if !reflect.ValueOf(s).IsNil() {
		return false
	}

	c.notFoundResponse()

	return true
}

// === shortcut response ===
func (c *Controller) unauthorizedResponse() {
	c.Reply().Unauthorized().JSON(aah.Data{
		"message": "access denied",
	})
}

func (c *Controller) forbiddenResponse() {
	c.Reply().Forbidden().JSON(aah.Data{
		"message": "access denied",
	})
}

func (c *Controller) notFoundResponse() {
	c.Reply().NotFound().JSON(aah.Error{
		Code: http.StatusNotFound,
		Message: "Not Found",
	})
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
// === end shortcut response ===

func (c *Controller) hateoas(s rest.Hateoasable) error {
	return s.GenerateHateoas(c.Context)
}
