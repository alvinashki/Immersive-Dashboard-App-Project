package user

type Core struct {
	ID         int
	Name       string
	Email      string
	Password   string
	Role       string
	Status     string
	DivisionId int
	Division   DivisionCore
}

type DivisionCore struct {
	ID       int
	Division string
}

type UsecaseInterface interface {
	CreateData(data Core) (row int, err error)
	GetAllUser() (data []Core, err error)
	PutUser(data Core, id int) (row int, err error)
}

type DataInterface interface {
	InsertData(data Core) (row int, err error)
	SelectAlUser() (data []Core, err error)
	UpdateData(data Core, id int) (row int, err error)
}
