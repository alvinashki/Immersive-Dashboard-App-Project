package delivery

import "gp3/features/mentee"

type MenteeResponse struct {
	ID           uint   `json:"id" form:"id"`
	Name         string `json:"name" form:"name"`
	Gender       string `json:"gender" form:"gender"`
	Address      string `json:"address" form:"address"`
	Home_Address string `json:"home_address" form:"home_address"`
	Class_Id     uint   `json:"class_id" form:"class_id"`
	Email        string `json:"email" form:"email"`
	Telegram     string `json:"telegram" form:"telegram"`
	Phone        string `json:"phone" form:"phone"`
	Status       string `json:"status" form:"status"`
	Category     string `json:"category" form:"category"`
	Name_Ed      string `json:"name_ed" form:"name_ed"`
	Phone_Ed     string `json:"phone_ed" form:"phone_ed"`
	Status_Ed    string `json:"status_ed" form:"status_ed"`
	Major        string `json:"major" form:"major"`
	Graduate     string `json:"graduate" form:"graduate"`
}

func FromCore(dataCore mentee.Core) MenteeResponse {
	return MenteeResponse{
		ID:           dataCore.ID,
		Name:         dataCore.Name,
		Gender:       dataCore.Gender,
		Address:      dataCore.Address,
		Home_Address: dataCore.Home_Address,
		Class_Id:     dataCore.Class_Id,
		Email:        dataCore.Email,
		Telegram:     dataCore.Telegram,
		Phone:        dataCore.Phone,
		Status:       dataCore.Status,
		Category:     dataCore.Category,
		Name_Ed:      dataCore.Name_Ed,
		Phone_Ed:     dataCore.Phone_Ed,
		Status_Ed:    dataCore.Status_Ed,
		Major:        dataCore.Major,
		Graduate:     dataCore.Graduate,
	}

}

func FromCoreList(dataCore []mentee.Core) []MenteeResponse {
	var dataResponse []MenteeResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, FromCore(v))
	}
	return dataResponse

}
