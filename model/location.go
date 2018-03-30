package model

import (
	"github.com/jinzhu/gorm"
)

// Location struct
type Location struct {
	*gorm.Model
	PlaceID string `gorm:"column:place_id,not null"`
}
