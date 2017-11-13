package filters

import "github.com/jinzhu/gorm"

// Alert filter
type Alert struct {
	Pagination
	Code      string `bind:"code_train"`
	StationID int    `bind:"station_id"`
}

// ApplyFilter on query
func (f *Alert) ApplyFilter(db *gorm.DB) *gorm.DB {
	if f.Code != "" {
		db = db.Where("code_train = ?", f.Code)
	}

	if f.StationID != 0 {
		db = db.Where("station_id = ?", f.StationID)
	}

	return db
}
