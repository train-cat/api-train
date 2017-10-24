package filters

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Stop struct {
	Pagination
	Terminus string `bind:"filter[terminus_name]"`
	Mission  string `bind:"filter[mission]"`
}

func (f *Stop) ApplyFilter(db *gorm.DB) *gorm.DB {
	if f.Mission != "" {
		db = db.Where("mission = ?", f.Mission)
	}

	if f.Terminus != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", f.Terminus))
	}

	return db
}
