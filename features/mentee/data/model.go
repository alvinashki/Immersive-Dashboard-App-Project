package data

import (
	// modelCLass "gp3/features/class/data"
	"gp3/features/mentee"

	"gorm.io/gorm"
)

type Mentee struct {
	gorm.Model
	Name         string
	Gender       string
	Address      string
	Home_Address string
	Class_Id     uint
	Email        string
	Telegram     string
	Phone        string
	Status       string
	Category     string
	Name_Ed      string
	Phone_Ed     string
	Status_Ed    string
	Major        string
	Graduate     string
	Class        Class
}

type Class struct {
	gorm.Model
	Class   string
	Mentees []Mentee
}

func fromCore(dataCore mentee.Core) Mentee {
	return Mentee{
		Name:         dataCore.Name,
		Gender:       dataCore.Gender,
		Address:      dataCore.Address,
		Home_Address: dataCore.Home_Address,
		Class_Id:     dataCore.Class_Id,
		Email:        dataCore.Email,
		Telegram:     dataCore.Telegram,
		Phone:        dataCore.Phone,
		Status:       dataCore.Status,
		Category:     dataCore.Category,
		Name_Ed:      dataCore.Name_Ed,
		Phone_Ed:     dataCore.Phone_Ed,
		Status_Ed:    dataCore.Status_Ed,
		Major:        dataCore.Major,
		Graduate:     dataCore.Graduate,
	}
}

func (dataMentee *Mentee) toCore() mentee.ResponseCore {
	return mentee.ResponseCore{
		ID:           dataMentee.ID,
		Name:         dataMentee.Name,
		Gender:       dataMentee.Gender,
		Address:      dataMentee.Address,
		Home_Address: dataMentee.Home_Address,
		Class:        dataMentee.Class.Class,
		Email:        dataMentee.Email,
		Telegram:     dataMentee.Telegram,
		Phone:        dataMentee.Phone,
		Status:       dataMentee.Status,
		Category:     dataMentee.Category,
		Name_Ed:      dataMentee.Name_Ed,
		Phone_Ed:     dataMentee.Phone_Ed,
		Status_Ed:    dataMentee.Status_Ed,
		Major:        dataMentee.Major,
		Graduate:     dataMentee.Graduate,
	}
}

func toCoreList(dataMentee []Mentee) []mentee.ResponseCore {
	var dataCore []mentee.ResponseCore

	for key := range dataMentee {
		dataCore = append(dataCore, dataMentee[key].toCore())

	}

	return dataCore

}
