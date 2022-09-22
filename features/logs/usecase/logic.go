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

/*
func (usecase *logUsecase) GetAllUser() ([]user.ResponseCore, error) {
	results, err := usecase.logData.SelectAlUser()
	return results, err
}

func (usecase *logUsecase) PutUser(newUser user.Core, id int) (int, error) {

	row, err := usecase.logData.UpdateData(newUser, id)
	return row, err
}

func (usecase *logUsecase) DeleteAkun(id int) (int, error) {
	row, err := usecase.logData.DeleteData(id)
	return row, err
}

func (usecase *logUsecase) GetUserId(id int) (user.ResponseCore, error) {
	result, err := usecase.logData.SelectUserId(id)
	if err != nil {
		return user.ResponseCore{}, err
	}
	return result, nil
}
*/
