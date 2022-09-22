package data

import (
	"errors"
	"gp3/features/user"

	"gorm.io/gorm"
)

type userData struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.DataInterface {
	return &userData{
		db: db,
	}
}

func (repo *userData) InsertData(newUser user.Core) (int, error) {
	userModel := fromCore(newUser)

	tx := repo.db.Create(&userModel)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return int(tx.RowsAffected), nil
}

func (repo *userData) SelectAlUser() ([]user.ResponseCore, error) {
	var allUserData []User
	tx := repo.db.Preload("Division").Find(&allUserData)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return toCoreList(allUserData), nil

}

func (repo *userData) UpdateData(data user.Core, id int) (row int, err error) {
	tx := repo.db.Model(&User{}).Where("id = ?", id).Updates(fromCore(data))
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("failed to update data")
	}
	return int(tx.RowsAffected), nil
}

func (repo *userData) DeleteData(id int) (row int, err error) {
	tx := repo.db.Delete(&User{}, id)
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("failed delete akun")
	}
	return int(tx.RowsAffected), nil
}

func (repo *userData) SelectUserId(id int) (user.ResponseCore, error) {
	var userData User
	// userData.ID = uint(id)

	tx := repo.db.Preload("Division").First(&userData, id)

	if tx.Error != nil {
		return user.ResponseCore{}, tx.Error
	}
	return userData.toCore(), nil
}
