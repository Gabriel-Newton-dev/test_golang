package models

import "gorm.io/gorm"

type Car struct {
	gorm.Model
	ID     int
	Brand  string //marca
	Name   string
	Models string
	Value  float64
	Turbo  bool
}

var Cars []Car
