package repositories

import "github.com/train-cat/api-train/app/models"

type alert struct{}

var Alert alert

func (_ alert) Persist(s *models.Stop, i *models.AlertInput) (*models.Alert, error) {
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
