package data

import (
	"errors"
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

func (repo *menteeData) FindMentee(class_id int, category, status string) ([]mentee.ResponseCore, error) {
	var dataMentee []Mentee

	if class_id != 0 && category != "" && status != "" {
		tx := repo.db.Where("class_id = ? AND category = ? AND status = ?", class_id, category, status).Preload("Class").Find(&dataMentee)

		if tx.Error != nil {
			return []mentee.ResponseCore{}, tx.Error
		}
	} else if class_id != 0 && category != "" {
		tx := repo.db.Where("class_id = ? AND category = ?", class_id, category).Preload("Class").Find(&dataMentee)

		if tx.Error != nil {
			return []mentee.ResponseCore{}, tx.Error
		}
	} else if class_id != 0 && status != "" {
		tx := repo.db.Where("class_id = ? AND status = ?", class_id, status).Preload("Class").Find(&dataMentee)

		if tx.Error != nil {
			return []mentee.ResponseCore{}, tx.Error
		}
	} else if category != "" && status != "" {
		tx := repo.db.Where("category = ? AND status = ?", category, status).Preload("Class").Find(&dataMentee)

		if tx.Error != nil {
			return []mentee.ResponseCore{}, tx.Error
		}
	} else if class_id != 0 {
		tx := repo.db.Where("class_id = ?", class_id).Preload("Class").Find(&dataMentee)

		if tx.Error != nil {
			return []mentee.ResponseCore{}, tx.Error
		}
	} else if category != "" {
		tx := repo.db.Where("category = ?", category).Preload("Class").Find(&dataMentee)

		if tx.Error != nil {
			return []mentee.ResponseCore{}, tx.Error
		}
	} else if status != "" {
		tx := repo.db.Where("status = ?", status).Preload("Class").Find(&dataMentee)

		if tx.Error != nil {
			return []mentee.ResponseCore{}, tx.Error
		}
	} else {
		tx := repo.db.Preload("Class").Find(&dataMentee)

		if tx.Error != nil {
			return []mentee.ResponseCore{}, tx.Error
		}
	}

	return toCoreList(dataMentee), nil

}

func (repo *menteeData) FindMenteeById(mentee_id int) (mentee.ResponseCore, error) {
	var dataMentee Mentee

	tx := repo.db.Preload("Class").First(&dataMentee, mentee_id)

	if tx.Error != nil {
		return mentee.ResponseCore{}, tx.Error
	}

	return dataMentee.toCore(), nil

}

func (repo *menteeData) UpdateMentee(dataMentee mentee.Core) (int, error) {
	var newDataMentee Mentee

	tx_OldData := repo.db.First(&newDataMentee, dataMentee.ID)

	if tx_OldData.Error != nil {
		return -1, tx_OldData.Error
	}

	if dataMentee.Name != "" {
		newDataMentee.Name = dataMentee.Name
	}

	if dataMentee.Gender != "" {
		newDataMentee.Gender = dataMentee.Gender
	}

	if dataMentee.Address != "" {
		newDataMentee.Address = dataMentee.Address
	}

	if dataMentee.Home_Address != "" {
		newDataMentee.Home_Address = dataMentee.Home_Address
	}

	if dataMentee.Class_Id != 0 {
		newDataMentee.Class_Id = dataMentee.Class_Id
	}

	if dataMentee.Email != "" {
		newDataMentee.Email = dataMentee.Email
	}

	if dataMentee.Telegram != "" {
		newDataMentee.Telegram = dataMentee.Telegram
	}

	if dataMentee.Phone != "" {
		newDataMentee.Phone = dataMentee.Phone
	}

	if dataMentee.Status != "" {
		newDataMentee.Status = dataMentee.Status
	}

	if dataMentee.Category != "" {
		newDataMentee.Category = dataMentee.Category
	}

	if dataMentee.Name_Ed != "" {
		newDataMentee.Name_Ed = dataMentee.Name_Ed
	}

	if dataMentee.Phone_Ed != "" {
		newDataMentee.Phone_Ed = dataMentee.Phone_Ed
	}

	if dataMentee.Status_Ed != "" {
		newDataMentee.Status_Ed = dataMentee.Status_Ed
	}

	if dataMentee.Major != "" {
		newDataMentee.Major = dataMentee.Major
	}

	if dataMentee.Graduate != "" {
		newDataMentee.Graduate = dataMentee.Graduate
	}

	tx_newData := repo.db.Save(&newDataMentee)

	if tx_newData.Error != nil {
		return -1, tx_newData.Error
	}

	if tx_newData.RowsAffected != 1 {
		return 0, errors.New("zero row affected, fail update")
	}

	return int(tx_newData.RowsAffected), nil

}

func (repo *menteeData) DeleteMentee(mentee_id int) (int, error) {
	var dataMentee Mentee
	tx := repo.db.Delete(&dataMentee, mentee_id)

	if tx.Error != nil {
		return -1, tx.Error
	}

	return int(tx.RowsAffected), nil

}
