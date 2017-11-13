package repositories

import (
	"github.com/train-cat/api-train/app/filters"
	"github.com/train-cat/api-train/app/models"
)

type alert struct{}

// Alert namespace
var Alert alert

// Persist on alert in database
func (r alert) Persist(s *models.Stop, i *models.AlertInput) (*models.Alert, error) {
	a := i.ToEntity()
	var err error

	a.CodeTrain = *s.Train.Code
	a.StationID = s.StationID
	a.Station = s.Station
	a.Action, err = Action.FindOne(*a.ActionID)

	if err != nil {
		return nil, err
	}

	err = db.Save(a).Error

	return a, err
}

// FindAll alerts filtered
func (r alert) FindAll(f filters.Filter) (*models.Collection, error) {
	alerts := []*models.Alert{}

	db := db.Model(models.Alert{}).
		Preload("Station").
		Preload("Action")

	return NewCollection(f, db, &alerts)
}
