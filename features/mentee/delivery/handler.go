package delivery

import (
	"gp3/features/mentee"
	"gp3/middlewares"
	"gp3/utils/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MenteeDelivery struct {
	menteeUsecase mentee.UsecaseInterface
}

func New(e *echo.Echo, usecase mentee.UsecaseInterface) {
	handler := &MenteeDelivery{
		menteeUsecase: usecase,
	}

	e.POST("/mentees", handler.PostNewMentee, middlewares.JWTMiddleware())
	e.GET("/mentees", handler.GetAllMentee, middlewares.JWTMiddleware())
	e.GET("/mentees/:id", handler.GetMenteeById, middlewares.JWTMiddleware())
	e.PUT("/mentees/:id", handler.UpdateMenteeData, middlewares.JWTMiddleware())
	e.DELETE("/mentees/:id", handler.DeleteMenteeData, middlewares.JWTMiddleware())

}

func (delivery *MenteeDelivery) PostNewMentee(c echo.Context) error {
	var menteeRequestData MenteeRequest

	errBind := c.Bind(&menteeRequestData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("fail bind data"))
	}

	row, err := delivery.menteeUsecase.InsertMentee(ToCore(menteeRequestData))

	if err != nil || row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("success add new mentee"))
}

func (delivery *MenteeDelivery) GetAllMentee(c echo.Context) error {
	category := c.QueryParam("category")
	status := c.QueryParam("status")
	class_id, err := strconv.Atoi(c.QueryParam("class_id"))

	if err != nil && class_id != 0 {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("fail converse class_id param"))
	}

	dataMentee, err := delivery.menteeUsecase.SelectMentee(class_id, category, status)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("get all data mentees success", FromCoreList(dataMentee)))
}

func (delivery *MenteeDelivery) GetMenteeById(c echo.Context) error {

	mentee_id := helper.ParamInt(c)

	dataMentee, err := delivery.menteeUsecase.SelectMenteeById(mentee_id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("get all data mentees success", FromCore(dataMentee)))
}

func (delivery *MenteeDelivery) UpdateMenteeData(c echo.Context) error {

	mentee_id := helper.ParamInt(c)

	var menteeRequestData MenteeRequest
	errBind := c.Bind(&menteeRequestData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("fail bind data"))
	}

	requestMenteeCore := ToCore(menteeRequestData)
	requestMenteeCore.ID = uint(mentee_id)

	row, err := delivery.menteeUsecase.UpdateMenteeData(requestMenteeCore)

	if err != nil || row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("update success"))
}

func (delivery *MenteeDelivery) DeleteMenteeData(c echo.Context) error {

	mentee_id := helper.ParamInt(c)

	row, err := delivery.menteeUsecase.DeleteMenteeData(mentee_id)

	if err != nil || row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("delete success"))
}
