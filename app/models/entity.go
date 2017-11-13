package models

import "time"

type (
	// Entity regroup all necessary field for database
	Entity struct {
		ID        uint       `json:"id" gorm:"primary_key"`
		CreatedAt time.Time  `json:"-"`
		UpdatedAt time.Time  `json:"-"`
		DeletedAt *time.Time `json:"deleted_at,omitempty" sql:"index"`
	}
)
