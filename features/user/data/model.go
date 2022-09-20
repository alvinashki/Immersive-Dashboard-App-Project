package data

import (
	"gp3/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name       string
	Email      string `gorm:"unique"`
	Password   string
	Role       string
	Status     string
	DivisionId int
	Division   user.DivisionCore
}

type Division struct {
	gorm.Model
	Division string
	User     []User
}

func fromCore(dataCore user.Core) User {
	return User{
		Name:       dataCore.Name,
		Email:      dataCore.Email,
		Password:   dataCore.Password,
		Role:       dataCore.Role,
		Status:     dataCore.Status,
		DivisionId: dataCore.Division.ID,
	}
}

func (dataUser *User) toCore() user.Core {
	return user.Core{
		ID:         int(dataUser.ID),
		Name:       dataUser.Name,
		Email:      dataUser.Email,
		Password:   dataUser.Password,
		Role:       dataUser.Role,
		Status:     dataUser.Status,
		DivisionId: dataUser.DivisionId,
	}
}
