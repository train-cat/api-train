package models

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// example https://play.golang.org/p/EpcL6nzlfV
const format = "%02d:%02d"

// Schedule is struct for time (not datetime)
type Schedule struct {
	Hour   uint `validate:"min=0,max=23"`
	Minute uint `validate:"min=0,max=59"`
	Valid  bool `validate:"true"`
}

// Value convert Schedule to friendly database value
func (s Schedule) Value() (driver.Value, error) {
	return s.String(), nil
}

// String return representation of Schedule in string
func (s Schedule) String() string {
	return fmt.Sprintf(format, s.Hour, s.Minute)
}

// Scan retrieve value from database
func (s *Schedule) Scan(value interface{}) error {
	if value == nil {
		*s = Schedule{0, 0, false}
		return nil
	}

	if strTime, ok := value.([]uint8); ok {
		t, err := time.Parse("15:04:05", string(strTime))

		if err != nil {
			return err
		}

		*s = Schedule{uint(t.Hour()), uint(t.Minute()), true}
		return nil
	}

	return errors.New("can't cast Time to Schedule")
}

// UnmarshalJSON convert JSON send by client to Schedule struct
func (s *Schedule) UnmarshalJSON(data []byte) error {
	str := strings.Trim(string(data), "\"")

	if str == "" {
		return nil
	}

	t := strings.Split(str, ":")

	if len(t) != 2 {
		return errors.New("invalid format for schedule")
	}

	h, err := strconv.Atoi(t[0])
	if err != nil {
		return err
	}

	m, err := strconv.Atoi(t[1])
	if err != nil {
		return err
	}

	s.Hour = uint(h)
	s.Minute = uint(m)
	s.Valid = true

	return nil
}
