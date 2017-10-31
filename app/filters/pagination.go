package filters

import "github.com/jinzhu/gorm"

const (
	defaultPage  = 1
	defaultLimit = 20
	maxLimit     = 100
)

type Pagination struct {
	Page  int `bind:"_page"`
	Limit int `bind:"_limit_per_page"`
}

// GetPage return safe page (no negative)
func (p *Pagination) GetPage() int {
	if p.Page <= 0 {
		return defaultPage
	}

	return p.Page
}

// GetLimit return safe limit_per_page (no negative or too large)
func (p *Pagination) GetLimit() int {
	if p.Limit <= 0 || p.Limit > maxLimit {
		return defaultLimit
	}

	return p.Limit
}

func (p *Pagination) ApplyPagination(db *gorm.DB) *gorm.DB {
	return db.
		Offset((p.GetPage() - 1) * p.GetLimit()).
		Limit(p.GetLimit())
}

func (p *Pagination) ApplyFilter(db *gorm.DB) *gorm.DB {
	return db
}
