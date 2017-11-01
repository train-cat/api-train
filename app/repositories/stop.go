package repositories

import (
	"github.com/train-cat/api-train/app/filters"
	"github.com/train-cat/api-train/app/models"
)

type stop struct{}

// Stop namespace
var Stop stop


// Persist one stop
func (r stop) Persist(s *models.Stop) error {
	return db.Save(s).Error
}

// FindAllByStation stops
func (r stop) FindAllByStation(stationID int, f filters.Filter) (*models.Collection, error) {
	stops := []*models.Stop{}

	db := db.
	Model(models.Stop{}).
		Preload("Train").
		Preload("Train.Terminus").
		Joins("LEFT JOIN train ON train_id = train.id").
		Joins("LEFT JOIN station ON terminus_id = station.id").
		Order("schedule ASC").
		Where("stop.station_id = ?", stationID)

	return NewCollection(f, db, &stops)
}

// FindAllByTrain stops
func (r stop) FindAllByTrain(code string, f filters.Filter) (*models.Collection, error) {
	stops := []*models.Stop{}

	db := db.
	Model(models.Stop{}).
		Preload("Station").
		Preload("Train").
		Preload("Train.Terminus").
		Joins("LEFT JOIN train ON train_id = train.id").
		Order("schedule ASC").
		Where("code = ?", code)

	return NewCollection(f, db, &stops)
}

// FindOneByStationAndCode stop
func (r stop) FindOneByStationAndCode(stationID int, code string) (*models.Stop, error) {
	stop := &models.Stop{}

	err := db.
		Preload("Station").
		Preload("Train").
		Joins("LEFT JOIN train ON train_id = train.id").
		Where("code = ? AND station_id = ?", code, stationID).
		First(stop).Error

	return stop, err
}

// IsExist return true if stop between stationID and code train exist
func (r stop) IsExist(stationID int, code string) (bool, error) {
	count := 0

	err := db.Model(&models.Stop{}).
		Joins("LEFT JOIN train ON train_id = train.id").
		Where("train.code = ? AND station_id = ?", code, stationID).
		Count(&count).Error

	return count > 0, err
}
