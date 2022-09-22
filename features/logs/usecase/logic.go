package usecase

import (
	// "errors"
	"gp3/features/logs"
)

type logUsecase struct {
	logData logs.DataInterface
}

func New(dataLogs logs.DataInterface) logs.UsecaseInterface {
	return &logUsecase{
		logData: dataLogs,
	}
}

func (usecase *logUsecase) CreateData(newLogs logs.Core) (int, error) {
	row, err := usecase.logData.InsertData(newLogs)
	return row, err
}

func (usecase *logUsecase) SelectFeedback(mentee_id int) ([]logs.ResponseCore, error) {
	dataLogs, err := usecase.logData.FindFeedback(mentee_id)
	return dataLogs, err
}
