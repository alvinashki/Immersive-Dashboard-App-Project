package delivery

import (
	"gp3/features/user"
	"gp3/middlewares"
	"gp3/utils/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Delivery struct {
	userUsecase user.UsecaseInterface
}

func New(e *echo.Echo, usecase user.UsecaseInterface) {
	handler := &Delivery{
		userUsecase: usecase,
	}
	e.POST("/adduser", handler.PostUser)
	e.GET("/getalldata", handler.GetAllUser)
	e.PUT("/updateprofile/:id", handler.PutUser, middlewares.JWTMiddleware())
}

func (deliv *Delivery) PostUser(c echo.Context) error {
	var dataRequest UserRequest
	errBind := c.Bind(&dataRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("error binding data"))
	}

	// fmt.Println("error =", dataRequest)
	row, err := deliv.userUsecase.CreateData(toCore(dataRequest))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("error insert data"))
	}
	if row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("error insert data"))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponseHelper("success insert data"))
}

func (deliv *Delivery) GetAllUser(c echo.Context) error {
	result, err := deliv.userUsecase.GetAllUser()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("Failed to get data"))
	}
	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("Succes get data", fromCoreList(result)))
}

func (deliv *Delivery) PutUser(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	idToken := middlewares.ExtractToken(c)

	if idToken != idConv {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("you dont have acces"))
	}

	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("param must be a number"))
	}

	var dataUpdate UserRequest
	errBind := c.Bind(&dataUpdate)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("error binding data"))
	}

	row, err := deliv.userUsecase.PutUser(toCore(dataUpdate), idConv)

	if err != nil || row == 0 {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("failed to update data"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("success update data"))
}
