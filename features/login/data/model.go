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
<<<<<<< HEAD
	Role     string `json:"role" form:"role"` //coba
=======
	Role     string `json:"role" form:"role"`
>>>>>>> 476986a755198c287082c23df981b94e38b8b7b9
}

func toCore(user User) login.Core {

	core := login.Core{
		ID:       int(user.ID),
		Email:    user.Email,
		Password: user.Password,
<<<<<<< HEAD
		Role:     user.Role, //coba
=======
		Role:     user.Role,
>>>>>>> 476986a755198c287082c23df981b94e38b8b7b9
	}
	fmt.Println(core.ID)
	fmt.Println("role =", core.Role)
	return core

}
