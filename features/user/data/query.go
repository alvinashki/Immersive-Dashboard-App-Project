package data

import (
	"fmt"
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

func (repo *userData) SelectAlUser() ([]user.Core, error) {
	var allUserData []User
	tx := repo.db.Find(&allUserData)

	if tx.Error != nil {
		return nil, tx.Error
	}

	userList := toCoreList(allUserData)
	fmt.Println(userList)
	fmt.Println(allUserData)
	return userList, nil
}
