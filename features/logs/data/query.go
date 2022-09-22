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

// func (repo *logData) SelectAlUser() ([]user.ResponseCore, error) {
// 	var alllogData []User
// 	tx := repo.db.Preload("Division").Find(&alllogData)

// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}

// 	return toCoreList(alllogData), nil

// }

// func (repo *logData) UpdateData(data user.Core, id int) (row int, err error) {
// 	tx := repo.db.Model(&User{}).Where("id = ?", id).Updates(fromCore(data))
// 	if tx.Error != nil {
// 		return -1, tx.Error
// 	}
// 	if tx.RowsAffected == 0 {
// 		return 0, errors.New("failed to update data")
// 	}
// 	return int(tx.RowsAffected), nil
// }

// func (repo *logData) DeleteData(id int) (row int, err error) {
// 	tx := repo.db.Delete(&User{}, id)
// 	if tx.Error != nil {
// 		return -1, tx.Error
// 	}
// 	if tx.RowsAffected == 0 {
// 		return 0, errors.New("failed delete akun")
// 	}
// 	return int(tx.RowsAffected), nil
// }

// func (repo *logData) SelectUserId(id int) (user.ResponseCore, error) {
// 	var logData User
// 	// logData.ID = uint(id)

// 	tx := repo.db.Preload("Division").First(&logData, id)

// 	if tx.Error != nil {
// 		return user.ResponseCore{}, tx.Error
// 	}
// 	return logData.toCore(), nil
// }
