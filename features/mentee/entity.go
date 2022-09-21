package mentee

type Core struct {
	ID           uint
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
}

type ResponseCore struct {
	ID           uint
	Name         string
	Gender       string
	Address      string
	Home_Address string
	Class        string
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
}

type UsecaseInterface interface {
	SelectMentee() (data []ResponseCore, err error)
	InsertMentee(data Core) (row int, err error)
}

type DataInterface interface {
	FindMentee() (data []ResponseCore, err error)
	CreateMentee(data Core) (row int, err error)
}
