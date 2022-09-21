package delivery

import (
	"gp3/features/mentee"
	"gp3/utils/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type MenteeDelivery struct {
	menteeUsecase mentee.UsecaseInterface
}

func New(e *echo.Echo, usecase mentee.UsecaseInterface) {
	handler := &MenteeDelivery{
		menteeUsecase: usecase,
	}

	e.POST("/mentees", handler.PostNewMentee)
	e.GET("/mentees", handler.GetAllMentee)
	e.GET("/mentees/:id", handler.GetMenteeById)

}

func (delivery *MenteeDelivery) PostNewMentee(c echo.Context) error {
	var menteeRequestData MenteeRequest

	errBind := c.Bind(&menteeRequestData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("fail bind data"))
	}

	menteeRequestData.Status = "Interview"

	row, err := delivery.menteeUsecase.InsertMentee(ToCore(menteeRequestData))

	if err != nil || row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("success add new mentee"))
}

func (delivery *MenteeDelivery) GetAllMentee(c echo.Context) error {
	dataMentee, err := delivery.menteeUsecase.SelectMentee()

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

	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("get all data mentee success", FromCore(dataMentee)))
}
