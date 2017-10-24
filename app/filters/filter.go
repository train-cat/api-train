package filters

import "github.com/jinzhu/gorm"

type Filter interface {
	ApplyFilter(db *gorm.DB) *gorm.DB
	ApplyPagination(db *gorm.DB) *gorm.DB
	GetPage() int
	GetLimit() int
}
