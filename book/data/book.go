package data

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string
	Genres pq.StringArray `gorm:"type:text[]"`
}
