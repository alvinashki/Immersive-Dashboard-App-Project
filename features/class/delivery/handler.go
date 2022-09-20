package delivery

import (
	"gp3/features/class"

	"github.com/labstack/echo/v4"
)

type CLassDelivery struct {
	classUsecase class.UsecaseInterface
}

func New(e *echo.Echo, usecase class.UsecaseInterface) {
	// handler := &CLassDelivery{
	// 	classUsecase: usecase,
	// }

}
