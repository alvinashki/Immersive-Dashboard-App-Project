package delivery

import "gp3/features/user"

type UserResponse struct {
	ID         int    `json:"id" form:"id"`
	Name       string `json:"name" form:"name"`
	Email      string `json:"email" form:"email"`
	Role       string `json:"role" form:"role"`
	Status     string `json:"status" form:"status"`
	DivisionId int    `json:"division_id" form:"division_id"`
}

func fromCore(dataCore user.Core) UserResponse {
	return UserResponse{
		ID:         dataCore.ID,
		Name:       dataCore.Name,
		Email:      dataCore.Email,
		Role:       dataCore.Role,
		Status:     dataCore.Status,
		DivisionId: dataCore.DivisionId,
	}
}

func fromCoreList(dataCore []user.Core) []UserResponse {
	var dataResponse []UserResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
