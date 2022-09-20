package delivery

import "gp3/features/class"

type ClassResponse struct {
	ID    uint   `json:"id" form:"id"`
	Class string `json:"class" form:"class"`
}

func FromCore(dataCore class.Core) ClassResponse {
	return ClassResponse{
		ID:    dataCore.ID,
		Class: dataCore.Class,
	}

}

func FromCoreList(dataCore []class.Core) []ClassResponse {
	var dataResponse []ClassResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, FromCore(v))
	}
	return dataResponse

}
