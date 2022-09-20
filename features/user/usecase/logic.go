package usecase

import (
	"errors"
	"gp3/features/user"
)

type userUsecase struct {
	userData user.DataInterface
}

func New(dataUser user.DataInterface) user.UsecaseInterface {
	return &userUsecase{
		userData: dataUser,
	}
}

func (usecase *userUsecase) CreateData(newUser user.Core) (int, error) {
	if newUser.Email == "" || newUser.Password == "" {
		return -1, errors.New("email dan password tidak boleh kosong")
	}

	row, err := usecase.userData.InsertData(newUser)
	return row, err
}
