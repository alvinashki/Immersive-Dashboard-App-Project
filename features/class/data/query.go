package data

import (
	"errors"
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

func (repo *classData) FindClass() ([]class.Core, error) {
	var dataClass []Class

	tx := repo.db.Find(&dataClass)

	if tx.Error != nil {
		return []class.Core{}, tx.Error
	}

	return toCoreList(dataClass), nil

}

func (repo *classData) UpdateClass(dataClass class.Core) (int, error) {
	var newDataClass Class

	tx_OldData := repo.db.Find(&newDataClass, dataClass.ID)

	if tx_OldData.Error != nil {
		return -1, tx_OldData.Error
	}

	if dataClass.Class != "" {
		newDataClass.Class = dataClass.Class
	}

	tx_newData := repo.db.Save(&newDataClass)

	if tx_newData.Error != nil {
		return -1, tx_newData.Error
	}

	if tx_newData.RowsAffected == 0 {
		return 0, errors.New("zero row affected, fail update")
	}

	return int(tx_newData.RowsAffected), nil
}

func (repo *classData) DeleteClass(class_id int) (int, error) {
	var dataClass Class
	tx := repo.db.Delete(&dataClass, class_id)

	if tx.Error != nil {
		return -1, tx.Error
	}

	return int(tx.RowsAffected), nil

}
