package models

import (
	"database/sql/driver"
	"errors"
	"strings"
)

const separator = ","

// SliceString alias for []string
type SliceString []string

// Value return value can be store in database
func (s SliceString) Value() (driver.Value, error) {
	return strings.Join(s, separator), nil
}

// Scan retrieve value from database
func (s *SliceString) Scan(value interface{}) error {
	if value == nil {
		*s = SliceString{}
		return nil
	}

	if bs, ok := value.([]byte); ok {
		*s = strings.Split(string(bs), separator)
		return nil
	}

	return errors.New("can't cast SliceString to string")
}
