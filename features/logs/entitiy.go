package logs

import "time"

type Core struct {
	ID       uint
	Feedback string
	Status   string
	UserId   uint
	MenteeId uint
}

type ResponseCore struct {
	ID          uint
	Feedback    string
	Status      string
	CreatedAt   time.Time
	Name_User   string
	Name_Mentee string
}

type UsecaseInterface interface {
	CreateData(data Core) (row int, err error)
}

type DataInterface interface {
	InsertData(data Core) (row int, err error)
}
