package class

type Core struct {
	ID    uint
	Class string
}

type UsecaseInterface interface {
	InsertClass(data Core) (row int, err error)
	SelectClass() (data []Core, err error)
}

type DataInterface interface {
	CreateClass(data Core) (row int, err error)
	FindClass() (data []Core, err error)
}
