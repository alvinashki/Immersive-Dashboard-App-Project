package usecase

import (
	"errors"
	"gp3/features/class"
)

type classUsecase struct {
	classData class.DataInterface
}

func New(data class.DataInterface) class.UsecaseInterface {
	return &classUsecase{
		classData: data,
	}
}

func (usecase *classUsecase) InsertClass(dataClass class.Core) (int, error) {

	if dataClass.Class == "" {
		return -1, errors.New("class name must be filled")
	}

	rowCreate, errCreate := usecase.classData.CreateClass(dataClass)
	return rowCreate, errCreate

}
