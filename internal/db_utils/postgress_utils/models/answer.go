package models

import "gorm.io/gorm"

type Answer struct {
	gorm.Model
	Key     string
	Value   string
	Deleted bool
}
