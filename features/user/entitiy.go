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
}

type DataInterface interface {
	InsertData(data Core) (row int, err error)
}
