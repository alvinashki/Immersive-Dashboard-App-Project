package usecase

import (
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
