package delivery

import (
	"gp3/features/logs"
	"time"
)

type LogsRequest struct {
	Feedback  string    `json:"Feedback" form:"Feedback"`
	Status    string    `json:"status" form:"status"`
	CreatedAt time.Time `json:"created_At" form:"created_at"`
	UserId    int       `json:"user_id" form:"user_id"`
	MenteeId  int       `json:"mentee_id" form:"mentee_id"`
}

func toCore(dataRequest LogsRequest) logs.Core {
	return logs.Core{
		Feedback:  dataRequest.Feedback,
		Status:    dataRequest.Status,
		CreatedAt: dataRequest.CreatedAt,
		UserId:    dataRequest.UserId,
		MenteeId:  dataRequest.MenteeId,
	}
}
