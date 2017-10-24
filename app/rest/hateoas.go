package rest

import (
	"fmt"
	"strings"

	"aahframework.org/aah.v0"
)

const (
	LinkSelf     = "self"
	LinkFirst    = "first"
	LinkLast     = "last"
	LinkPrevious = "previous"
	LinkNext     = "next"

	QueryPage = "_page"
)

type (
	Hateoas struct {
		Embedded Embedded `json:"_embedded,omitempty" gorm:"-"`
		Links    Links    `json:"_links,omitempty"    gorm:"-"`
	}

	Embedded map[string]interface{}
	Links    map[string]Link

	Link struct {
		Href string `json:"href"`
	}

	Hateoasable interface {
		GenerateHateoas(ctx *aah.Context) error
	}
)

func GenerateURI(ctx *aah.Context, routeName string, args ...interface{}) string {
	uri := ctx.ReverseURL(routeName, args...)

	return strings.TrimPrefix(uri, fmt.Sprintf("//%s", ctx.Req.Host))
}
