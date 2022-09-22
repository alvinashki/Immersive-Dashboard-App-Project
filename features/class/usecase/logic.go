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

func (usecase *classUsecase) SelectClass() ([]class.Core, error) {
	dataClass, err := usecase.classData.FindClass()
	return dataClass, err
}

func (usecase *classUsecase) UpdateClassData(dataClass class.Core) (int, error) {
	rowUpdate, errUpdate := usecase.classData.UpdateClass(dataClass)
	return rowUpdate, errUpdate

}

func (usecase *classUsecase) DeleteClassData(class_id int) (int, error) {
	rowDelete, errDelete := usecase.classData.DeleteClass(class_id)
	return rowDelete, errDelete

}
