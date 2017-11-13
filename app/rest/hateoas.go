package rest

import (
	"fmt"
	"strings"

	"aahframework.org/aah.v0"
)

// List of available links for collection
const (
	LinkSelf     = "self"
	LinkFirst    = "first"
	LinkLast     = "last"
	LinkPrevious = "previous"
	LinkNext     = "next"

	// QueryPage is query param use for the pagination
	QueryPage = "_page"
)

type (
	// Hateoas struct with all necessary component for hateoas response
	Hateoas struct {
		Embedded Embedded `json:"_embedded,omitempty" gorm:"-"`
		Links    Links    `json:"_links,omitempty"    gorm:"-"`
	}

	// Embedded struct in hateoas response
	Embedded map[string]interface{}
	// Links to another resources
	Links map[string]Link

	// Link to another resource
	Link struct {
		Href string `json:"href"`
	}

	// Hateoasable interface for generate hateoas content
	Hateoasable interface {
		GenerateHateoas(ctx *aah.Context) error
	}
)

// GenerateURI return uri for one given route
func GenerateURI(ctx *aah.Context, routeName string, args ...interface{}) string {
	uri := ctx.ReverseURL(routeName, args...)

	return strings.TrimPrefix(uri, fmt.Sprintf("//%s", ctx.Req.Host))
}
