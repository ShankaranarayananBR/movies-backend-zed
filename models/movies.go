package models

import "gorm.io/gorm"

// We will have the models for each DB table

type Movies struct {
	gorm.Model
	MovieName   string `gorm:"type:varchar(200); not null"`
	Description string `gorm:"type:varchar(200); not null"`
	MovieGenre  string `gorm:"type:varchar(200); not null"`
	MovieYear   int64  `gorm:"type:varchar(200); not null"`
}

func (mv *Movies) Tablename() string {
	return "movies"
}
