package user

type Core struct {
	ID         int
	Name       string
	Email      string
	Password   string
	Role       string
	Status     string
	DivisionId int
	// Division   Division
}

// type DivisionCore struct {
// 	ID       int
// 	Division string
// }

type ResponseCore struct {
	ID       int
	Name     string
	Email    string
	Password string
	Role     string
	Status   string
	Division string
}

type UsecaseInterface interface {
	CreateData(data Core) (row int, err error)
	GetAllUser() (data []ResponseCore, err error)
	PutUser(data Core, id int) (row int, err error)
	DeleteAkun(id int) (row int, err error)
	GetUserId(id int) (data ResponseCore, err error)
}

type DataInterface interface {
	InsertData(data Core) (row int, err error)
	SelectAlUser() (data []ResponseCore, err error)
	UpdateData(data Core, id int) (row int, err error)
	DeleteData(id int) (row int, err error)
	SelectUserId(id int) (data ResponseCore, err error)
}
