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

	pages := int(math.Ceil(float64(count) / float64(f.GetLimit())))

	if pages < f.GetPage() {
		return nil, nil
	}

	return &models.Collection{
		Page:  f.GetPage(),
		Limit: f.GetLimit(),
		Pages: pages,
		Total: count,
		Hateoas: rest.Hateoas{
			Embedded: rest.Embedded{
				"items": v,
			},
		},
	}, nil
}
