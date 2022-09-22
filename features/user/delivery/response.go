package delivery

import "gp3/features/user"

type UserResponse struct {
	ID       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Role     string `json:"role" form:"role"`
	Status   string `json:"status" form:"status"`
	Division string `json:"division" form:"division"`
}

func fromCore(dataCore user.ResponseCore) UserResponse {
	return UserResponse{
		ID:       dataCore.ID,
		Name:     dataCore.Name,
		Email:    dataCore.Email,
		Role:     dataCore.Role,
		Status:   dataCore.Status,
		Division: dataCore.Division,
	}
}

func fromCoreList(dataCore []user.ResponseCore) []UserResponse {
	var dataResponse []UserResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
