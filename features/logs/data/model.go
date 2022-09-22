package data

import (
	"gp3/features/logs"
	modelMentee "gp3/features/mentee/data"
	modelUser "gp3/features/user/data"

	"gorm.io/gorm"
)

type Logs struct {
	gorm.Model
	Feedback string
	Status   string
	UserId   uint
	User     modelUser.User `gorm:"foreignKey:UserId"`
	MenteeId uint
	Mentee   modelMentee.Mentee `gorm:"foreignKey:MenteeId"`
}

func fromCore(dataCore logs.Core) Logs {
	return Logs{
		Feedback: dataCore.Feedback,
		Status:   dataCore.Status,
		UserId:   dataCore.UserId,
		MenteeId: dataCore.MenteeId,
	}
}

func (dataLogs *Logs) toCore() logs.ResponseCore {
	return logs.ResponseCore{
		ID:          dataLogs.ID,
		Feedback:    dataLogs.Feedback,
		Status:      dataLogs.Status,
		CreatedAt:   dataLogs.CreatedAt,
		Name_User:   dataLogs.User.Name,
		Name_Mentee: dataLogs.Mentee.Name,
	}
}

func toCoreList(dataLogs []Logs) []logs.ResponseCore {
	var dataCore []logs.ResponseCore

	for key := range dataLogs {
		dataCore = append(dataCore, dataLogs[key].toCore())

	}

	return dataCore

}
