package repositories

import (
	"math"

	"github.com/jinzhu/gorm"
	"github.com/train-cat/api-train/app/filters"
	"github.com/train-cat/api-train/app/models"
	"github.com/train-cat/api-train/app/rest"
)

// NewCollection create collection for 'v' and apply filter
func NewCollection(f filters.Filter, db *gorm.DB, v interface{}) (*models.Collection, error) {
	db = f.ApplyFilter(db)

	count := 0
	err := db.Count(&count).Error

	if err != nil {
		return nil, err
	}

	db = f.ApplyPagination(db)

	err = db.Find(v).Error

	if err != nil {
		return nil, err
	}

	countPages := int(math.Ceil(float64(count) / float64(f.GetLimit())))
	currentPage := f.GetPage()

	// page doesn't exist (except page 1)
	if countPages < currentPage && currentPage != 1 {
		return nil, nil
	}

	return &models.Collection{
		Page:  currentPage,
		Limit: f.GetLimit(),
		Pages: countPages,
		Total: count,
		Hateoas: rest.Hateoas{
			Embedded: rest.Embedded{
				"items": v,
			},
		},
	}, nil
}
