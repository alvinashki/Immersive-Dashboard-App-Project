package delivery

import "gp3/features/class"

type ClassRequest struct {
	Class string `json:"class" form:"class"`
}

func ToCore(dataRequest ClassRequest) class.Core {
	return class.Core{
		Class: dataRequest.Class,
	}

}
