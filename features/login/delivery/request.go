package delivery

import "gp3/features/login"

type Request struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Role     string `json:"role" form:"role"`
}

func ToCore(data Request) login.Core {
	return login.Core{
		Email:    data.Email,
		Password: data.Password,
		Role:     data.Role,
	}
}
