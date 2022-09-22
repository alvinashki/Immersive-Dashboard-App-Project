package data

import (
	"gp3/features/logs"
	"time"

	"gorm.io/gorm"
)

type Logs struct {
	gorm.Model
	Feedback    string
	Status      string
	CreatedAt   time.Time
	UserId      int
	Name_User   User
	MenteeId    int
	Name_Mentee Mentee
}

type User struct {
	gorm.Model
	Name string
	Logs []Logs
}

type Mentee struct {
	gorm.Model
	Name string
	Logs []Logs
}

func fromCore(dataCore logs.Core) Logs {
	return Logs{
		Feedback:  dataCore.Feedback,
		Status:    dataCore.Status,
		CreatedAt: dataCore.CreatedAt,
		UserId:    dataCore.UserId,
		MenteeId:  dataCore.MenteeId,
	}
}

func (dataLogs *Logs) toCore() logs.ResponseCore {
	return logs.ResponseCore{
		ID:          int(dataLogs.ID),
		Feedback:    dataLogs.Feedback,
		Status:      dataLogs.Status,
		CreatedAt:   dataLogs.CreatedAt,
		Name_User:   dataLogs.Name_User.Name,
		Name_Mentee: dataLogs.Name_Mentee.Name,
	}
}

func toCoreList(dataLogs []Logs) []logs.ResponseCore {
	var dataCore []logs.ResponseCore

	for key := range dataLogs {
		dataCore = append(dataCore, dataLogs[key].toCore())

	}

	return dataCore

}
