package delivery

import (
	"gp3/features/mentee"

	"github.com/labstack/echo/v4"
)

type MenteeDelivery struct {
	menteeUsecase mentee.UsecaseInterface
}

func New(e *echo.Echo, usecase mentee.UsecaseInterface) {
	// handler := &MenteeDelivery{
	// 	menteeUsecase: usecase,
	// }

}
