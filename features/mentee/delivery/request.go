package delivery

import "gp3/features/mentee"

type MenteeRequest struct {
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

func ToCore(dataRequest MenteeRequest) mentee.Core {
	return mentee.Core{
		Name:         dataRequest.Name,
		Gender:       dataRequest.Gender,
		Address:      dataRequest.Address,
		Home_Address: dataRequest.Home_Address,
		Class_Id:     dataRequest.Class_Id,
		Email:        dataRequest.Email,
		Telegram:     dataRequest.Telegram,
		Phone:        dataRequest.Phone,
		Status:       dataRequest.Status,
		Category:     dataRequest.Category,
		Name_Ed:      dataRequest.Name_Ed,
		Phone_Ed:     dataRequest.Phone_Ed,
		Status_Ed:    dataRequest.Status_Ed,
		Major:        dataRequest.Major,
		Graduate:     dataRequest.Graduate,
	}

}
