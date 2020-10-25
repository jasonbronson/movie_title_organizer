package main

import (
	"gorm.io/gorm"
)

type MovieList struct {
	gorm.Model
	Title    string `gorm:"index:idx_movie"`
	Filename string `gorm:"index:idx_movie"`
	Type     string `gorm:"index:idx_movie"`
}
