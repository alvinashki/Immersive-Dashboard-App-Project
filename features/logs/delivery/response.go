package delivery

import (
	"gp3/features/logs"
	"time"
)

type LogsResponse struct {
	ID          uint      `json:"id" form:"id"`
	Feedback    string    `json:"Feedback" form:"Feedback"`
	Status      string    `json:"status" form:"status"`
	CreatedAt   time.Time `json:"created_At" form:"created_at"`
	Name_User   string    `json:"Name_User" form:"Name_User"`
	Name_Mentee string    `json:"Name_Mentee" form:"Name_Mentee"`
}

func fromCore(dataCore logs.ResponseCore) LogsResponse {
	return LogsResponse{
		ID:          dataCore.ID,
		Feedback:    dataCore.Feedback,
		Status:      dataCore.Status,
		CreatedAt:   dataCore.CreatedAt,
		Name_User:   dataCore.Name_User,
		Name_Mentee: dataCore.Name_Mentee,
	}
}

func fromCoreList(dataCore []logs.ResponseCore) []LogsResponse {
	var dataResponse []LogsResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
