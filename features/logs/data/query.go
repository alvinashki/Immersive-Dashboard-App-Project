package data

import (
	"gp3/features/logs"

	"gorm.io/gorm"
)

type logData struct {
	db *gorm.DB
}

func New(db *gorm.DB) logs.DataInterface {
	return &logData{
		db: db,
	}
}

func (repo *logData) InsertData(newLogs logs.Core) (int, error) {
	logsModel := fromCore(newLogs)

	tx := repo.db.Create(&logsModel)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return int(tx.RowsAffected), nil
}

func (repo *logData) FindFeedback(mentee_id int) ([]logs.ResponseCore, error) {
	var dataLogs []Logs

	if mentee_id != 0 {
		tx := repo.db.Where("mentee_id = ?", mentee_id).Joins("Mentee").Joins("User").Find(&dataLogs)

		if tx.Error != nil {
			return []logs.ResponseCore{}, tx.Error
		}
	}
	return toCoreList(dataLogs), nil

}
