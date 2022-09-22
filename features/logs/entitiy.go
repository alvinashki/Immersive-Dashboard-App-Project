package logs

import "time"

type Core struct {
	ID       uint
	Feedback string
	Status   string
	UserId   uint
	MenteeId uint
	File     string //s3
}

type ResponseCore struct {
	ID          uint
	Feedback    string
	Status      string
	CreatedAt   time.Time
	Name_User   string
	Name_Mentee string
	File        string //s3
}

type UsecaseInterface interface {
	CreateData(data Core) (row int, err error)
	SelectFeedback(mentee_id int) (data []ResponseCore, err error)
}

type DataInterface interface {
	InsertData(data Core) (row int, err error)
	FindFeedback(mentee_id int) (data []ResponseCore, err error)
}
