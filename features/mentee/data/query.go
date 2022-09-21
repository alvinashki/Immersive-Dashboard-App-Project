package data

import (
	"gp3/features/mentee"

	"gorm.io/gorm"
)

type menteeData struct {
	db *gorm.DB
}

func New(db *gorm.DB) mentee.DataInterface {
	return &menteeData{
		db: db,
	}

}

func (repo *menteeData) CreateMentee(dataMentee mentee.Core) (int, error) {
	newMentee := fromCore(dataMentee)

	tx := repo.db.Create(&newMentee)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return int(tx.RowsAffected), nil
}

func (repo *menteeData) FindMentee() ([]mentee.ResponseCore, error) {
	var dataMentee []Mentee

	tx := repo.db.Preload("Class").Find(&dataMentee)

	if tx.Error != nil {
		return []mentee.ResponseCore{}, tx.Error
	}

	return toCoreList(dataMentee), nil

}
