package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Json map[string]string

func (j Json) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *Json) Scan(value interface{}) error {
	if value == nil {
		*j = Json{}
		return nil
	}

	if bs, ok := value.([]byte); ok {
		json.Unmarshal(bs, j)
		return nil
	}

	return errors.New("can't cast Json to map[string]string")
}

func (j *Json) Get(field string, fallback string) string {
	m := map[string]string(*j)

	if v, ok := m[field]; ok {
		return v
	}

	return fallback
}
