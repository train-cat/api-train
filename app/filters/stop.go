package filters

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Stop filters
type Stop struct {
	Pagination
	TerminusID            uint   `bind:"terminus_id"`
	TerminusName          string `bind:"terminus_name"`
	TrainThroughStationID uint   `bind:"train_through_station_id"`
	Mission               string `bind:"mission"`
}

// ApplyFilter on query
func (f *Stop) ApplyFilter(db *gorm.DB) *gorm.DB {
	if f.Mission != "" {
		db = db.Where("mission = ?", f.Mission)
	}

	if f.TerminusName != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", f.TerminusName))
	}

	if f.TerminusID != 0 {
		db = db.Where("terminus_id = ?", f.TerminusID)
	}

	if f.TrainThroughStationID != 0 {
		db = db.
			Joins("LEFT JOIN stop AS stop_at ON train.id = stop_at.train_id").
			Where("(stop_at.station_id = ? OR terminus_id = ?)", f.TrainThroughStationID, f.TrainThroughStationID)
	}

	return db
}
