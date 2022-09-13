package models

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	Event   string
	Key     string
	Value   string
	Deleted bool
}
