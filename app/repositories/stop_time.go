package repositories

import (
	"github.com/train-cat/api-train/app/filters"
	"github.com/train-cat/api-train/app/models"
	"github.com/jinzhu/gorm"
	"time"
	"strconv"
)

type stopTime struct{}

// StopTime namespace
var StopTime stopTime

// FindAllByStation stops
func (r stopTime) FindAllByStation(stationID int, f filters.Filter) (*models.Collection, error) {
	stops := []*models.StopTime{}

	today, _ := strconv.Atoi(time.Now().Format("20060102"))

	db := db.
		Model(models.StopTime{}).
		Preload("Trip").
		Preload("Trip.Calendar").
		Joins("LEFT JOIN trip ON trip_id = trip.id").
		Joins("LEFT JOIN calendar ON trip.calendar_id = calendar.id").
		Order("stop_time.schedule ASC").
		Where("stop_time.station_id = ?", stationID).
		Where("calendar.start < ? AND ? < calendar.end", today, today)

	return NewCollection(f, db, &stops)
}

func (r stopTime) Find(id int) (*models.StopTime, error) {
	stop := &models.StopTime{}

	err := db.
		Preload("Trip").
		Preload("Station").
		Where("id = ?", id).First(stop).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return stop, nil
}
