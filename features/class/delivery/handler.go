package delivery

import (
	"gp3/features/class"
	"gp3/utils/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CLassDelivery struct {
	classUsecase class.UsecaseInterface
}

func New(e *echo.Echo, usecase class.UsecaseInterface) {
	handler := &CLassDelivery{
		classUsecase: usecase,
	}

	e.POST("/class", handler.PostNewClass)

}

func (delivery *CLassDelivery) PostNewClass(c echo.Context) error {
	var classRequestData ClassRequest

	errBind := c.Bind(&classRequestData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("fail bind data"))
	}

	row, err := delivery.classUsecase.InsertClass(ToCore(classRequestData))

	if err != nil || row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("success add new class"))
}
