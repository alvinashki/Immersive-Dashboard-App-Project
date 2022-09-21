package usecase

import (
	"errors"
	"gp3/features/mentee"
)

type menteeUsecase struct {
	menteeData mentee.DataInterface
}

func New(data mentee.DataInterface) mentee.UsecaseInterface {
	return &menteeUsecase{
		menteeData: data,
	}
}

func (usecase *menteeUsecase) InsertMentee(dataMentee mentee.Core) (int, error) {

	if dataMentee.Name == "" || dataMentee.Class_Id == 0 {
		return -1, errors.New("mentee name or class must be filled")
	}

	rowCreate, errCreate := usecase.menteeData.CreateMentee(dataMentee)
	return rowCreate, errCreate

}

func (usecase *menteeUsecase) SelectMentee() ([]mentee.ResponseCore, error) {
	dataMentee, err := usecase.menteeData.FindMentee()
	return dataMentee, err

}

func (usecase *menteeUsecase) SelectMenteeById(mentee_id int) (mentee.ResponseCore, error) {
	dataMentee, err := usecase.menteeData.FindMenteeById(mentee_id)
	return dataMentee, err
}
