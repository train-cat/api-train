package filters

import "github.com/jinzhu/gorm"

// StopTime filters
type StopTime struct {
	Pagination
	TrainThroughStationID uint   `bind:"train_through_station_id"`
	Mission               string `bind:"mission"`
	ScheduledBefore       string `bind:"scheduled_before"`
	ScheduledAfter        string `bind:"scheduled_after"`
	ScheduledAt           string `bind:"scheduled_at"`
}

// ApplyFilter on query
func (f *StopTime) ApplyFilter(db *gorm.DB) *gorm.DB {
	if f.Mission != "" {
		db = db.Where("mission = ?", f.Mission)
	}

	if f.TrainThroughStationID != 0 {
		db = db.
			Joins("LEFT JOIN stop_time AS stop_at ON trip.id = stop_at.trip_id").
			Where("stop_at.station_id = ? AND stop_time.schedule < stop_at.schedule", f.TrainThroughStationID)
	}

	if f.ScheduledBefore != "" {
		db = db.Where("stop_time.schedule < ?", f.ScheduledBefore)
	}

	if f.ScheduledAfter != "" {
		db = db.Where("stop_time.schedule > ?", f.ScheduledAfter)
	}

	if f.ScheduledAt != "" {
		db = db.Where("stop_time.schedule = ?", f.ScheduledAt)
	}

	return db
}
