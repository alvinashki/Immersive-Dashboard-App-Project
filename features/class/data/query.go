package data

import (
	"gp3/features/class"

	"gorm.io/gorm"
)

type classData struct {
	db *gorm.DB
}

func New(db *gorm.DB) class.DataInterface {
	return &classData{
		db: db,
	}

}
