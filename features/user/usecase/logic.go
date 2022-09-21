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
	if newUser.Email == "" || newUser.Password == "" || newUser.Role == "" {
		return -1, errors.New("email, password dan role tidak boleh kosong")
	}

	row, err := usecase.userData.InsertData(newUser)
	return row, err
}

func (usecase *userUsecase) GetAllUser() ([]user.Core, error) {
	results, err := usecase.userData.SelectAlUser()
	return results, err
}

func (usecase *userUsecase) PutUser(newUser user.Core, id int) (int, error) {

	row, err := usecase.userData.UpdateData(newUser, id)
	return row, err
}

func (usecase *userUsecase) DeleteAkun(id int) (int, error) {
	row, err := usecase.userData.DeleteData(id)
	return row, err
}

func (usecase *userUsecase) GetUserId(id int) (user.Core, error) {
	result, err := usecase.userData.SelectUserId(id)
	if err != nil {
		return user.Core{}, err
	}
	return result, nil
}
