package delivery

import (
	"gp3/features/class"
	"gp3/middlewares"
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

	e.POST("/class", handler.PostNewClass, middlewares.JWTMiddleware())
	e.GET("/class", handler.GetAllClass, middlewares.JWTMiddleware())
	e.PUT("/class/:id", handler.UpdateClassData, middlewares.JWTMiddleware())
	e.DELETE("/class/:id", handler.DeleteClassData, middlewares.JWTMiddleware())

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

func (delivery *CLassDelivery) GetAllClass(c echo.Context) error {
	dataClass, err := delivery.classUsecase.SelectClass()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("get all data class success", FromCoreList(dataClass)))
}

func (delivery *CLassDelivery) UpdateClassData(c echo.Context) error {
	class_id := helper.ParamInt(c)

	var classRequestData ClassRequest
	errBind := c.Bind(&classRequestData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("fail bind data"))
	}

	requestClassCore := ToCore(classRequestData)
	requestClassCore.ID = uint(class_id)

	row, err := delivery.classUsecase.UpdateClassData(requestClassCore)

	if err != nil || row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("update success"))
}

func (delivery *CLassDelivery) DeleteClassData(c echo.Context) error {
	class_id := helper.ParamInt(c)

	row, err := delivery.classUsecase.DeleteClassData(class_id)

	if err != nil || row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("delete success"))
}
