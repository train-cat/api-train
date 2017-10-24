package filters

import (
	"github.com/jinzhu/gorm"
	"fmt"
)

type Station struct {
	Pagination
	Name string `bind:"filter[name]"`
}

// ApplyFilter for station
func (f *Station) ApplyFilter(db *gorm.DB) *gorm.DB {
	if f.Name != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", f.Name))
	}

	return db
}
