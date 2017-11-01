package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// JSON is alias of map[string]string
type JSON map[string]string

// Value transform Json to friendly value for database
func (j JSON) Value() (driver.Value, error) {
	return json.Marshal(j)
}

// Scan retrieve value from database
func (j *JSON) Scan(value interface{}) error {
	if value == nil {
		*j = JSON{}
		return nil
	}

	if bs, ok := value.([]byte); ok {
		json.Unmarshal(bs, j)
		return nil
	}

	return errors.New("can't cast Json to map[string]string")
}

// Get return value in a given key, or fallback if key not exist
func (j *JSON) Get(field string, fallback string) string {
	m := map[string]string(*j)

	if v, ok := m[field]; ok {
		return v
	}

	return fallback
}
