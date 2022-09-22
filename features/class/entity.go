package class

type Core struct {
	ID    uint
	Class string
}

type UsecaseInterface interface {
	InsertClass(data Core) (row int, err error)
	SelectClass() (data []Core, err error)
	UpdateClassData(data Core) (row int, err error)
	DeleteClassData(class_id int) (row int, err error)
}

type DataInterface interface {
	CreateClass(data Core) (row int, err error)
	FindClass() (data []Core, err error)
	UpdateClass(data Core) (row int, err error)
	DeleteClass(class_id int) (row int, err error)
}
