package data

import (
	"fmt"
	"gp3/features/login"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Role     string `json:"role" form:"role"`
}

func toCore(user User) login.Core {

	core := login.Core{
		ID:       int(user.ID),
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
	}
	fmt.Println(core.ID)
	fmt.Println("role =", core.Role)
	return core

}
