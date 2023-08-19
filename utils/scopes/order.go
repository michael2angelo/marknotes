package scopes

import (
	"fmt"

	"gorm.io/gorm"
)

type Direction string

const (
	Ascending  Direction = "ASC"
	Descending Direction = "DESC"
)

func OrderBy(field string, direction Direction) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(fmt.Sprintf("%s %s", field, direction))
	}
}