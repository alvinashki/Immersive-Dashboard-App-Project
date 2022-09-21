package delivery

import "gp3/features/user"

type UserRequest struct {
	Name       string `json:"name" form:"name"`
	Email      string `json:"email" form:"email"`
	Password   string `json:"password" form:"password"`
	Role       string `json:"role" form:"role"`
	Status     string `json:"status" form:"status"`
	DivisionId int    `json:"division_id" form:"division_id"`
}

func toCore(dataRequest UserRequest) user.Core {
	return user.Core{
		Name:       dataRequest.Name,
		Email:      dataRequest.Email,
		Password:   dataRequest.Password,
		Role:       dataRequest.Role,
		Status:     dataRequest.Status,
		DivisionId: dataRequest.DivisionId,
	}
}
