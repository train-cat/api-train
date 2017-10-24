package models

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"

	"aahframework.org/aah.v0"
	"github.com/train-sh/api-train/app/rest"
)

type (
	Collection struct {
		Page  int `json:"page"`
		Limit int `json:"limit"`
		Pages int `json:"pages"`
		Total int `json:"total"`
		rest.Hateoas
	}
)

func (c *Collection) GenerateHateoas(ctx *aah.Context) error {
	u := ctx.Req.Unwrap().URL

	links, err := generateLinks(u, c.Pages)

	if err != nil {
		return err
	}

	c.Links = *links

	return c.hateoasItems(ctx)
}

func (c *Collection) hateoasItems(ctx *aah.Context) error {
	v := reflect.ValueOf(c.Embedded["items"]).Elem()
	var err error

	switch v.Kind() {
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {

			h, ok := v.Index(i).Interface().(rest.Hateoasable)

			if ok {
				if err = h.GenerateHateoas(ctx); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func generateLinks(u *url.URL, lastPage int) (*rest.Links, error) {
	links := rest.Links{}
	q := u.Query()
	currentPage := 1

	if q.Get(rest.QueryPage) != "" {
		var err error
		currentPage, err = strconv.Atoi(q.Get(rest.QueryPage))

		if err != nil {
			return nil, err
		}
	}

	links[rest.LinkSelf] = rest.Link{Href: u.RequestURI()}

	for link, page := range map[string]int{rest.LinkFirst: 1, rest.LinkLast: lastPage, rest.LinkPrevious: currentPage - 1, rest.LinkNext: currentPage + 1} {
		if page > 0 && page <= lastPage {
			q.Set(rest.QueryPage, fmt.Sprintf("%d", page))
			query, err := url.QueryUnescape(q.Encode())
			if err != nil {
				return nil, err
			}
			links[link] = rest.Link{Href: fmt.Sprintf("%s?%s", u.Path, query)}
		}
	}

	return &links, nil
}
