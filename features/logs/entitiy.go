package logs

import "time"

type Core struct {
	ID        int
	Feedback  string
	Status    string
	CreatedAt time.Time
	UserId    int
	MenteeId  int
}

type ResponseCore struct {
	ID          int
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
