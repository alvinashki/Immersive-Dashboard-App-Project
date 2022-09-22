package delivery

import (
	"gp3/features/logs"
)

type LogsRequest struct {
	Feedback string `json:"Feedback" form:"Feedback"`
	Status   string `json:"status" form:"status"`
	UserId   uint   `json:"user_id" form:"user_id"`
	MenteeId uint   `json:"mentee_id" form:"mentee_id"`
}

func toCore(dataRequest LogsRequest) logs.Core {
	return logs.Core{
		Feedback: dataRequest.Feedback,
		Status:   dataRequest.Status,
		UserId:   dataRequest.UserId,
		MenteeId: dataRequest.MenteeId,
	}
}
