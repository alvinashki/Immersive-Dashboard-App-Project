package data

import (
	"gp3/features/login"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func toCore(user User) login.Core {

	core := login.Core{
		ID:       int(user.ID),
		Email:    user.Email,
		Password: user.Password,
	}

	return core

}
