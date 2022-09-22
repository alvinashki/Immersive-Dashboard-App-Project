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

func (repo *classData) CreateClass(dataClass class.Core) (int, error) {
	newClass := fromCore(dataClass)

	tx := repo.db.Create(&newClass)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return int(tx.RowsAffected), nil
}
