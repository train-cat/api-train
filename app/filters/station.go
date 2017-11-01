package filters

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Station filter
type Station struct {
	Pagination
	Name string `bind:"name"`
}

// ApplyFilter on query
func (f *Station) ApplyFilter(db *gorm.DB) *gorm.DB {
	if f.Name != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", *f.Name))
	}

	return db
}
