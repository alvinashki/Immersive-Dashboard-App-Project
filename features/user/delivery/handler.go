package delivery

import (
	"fmt"
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
	e.POST("/users", handler.PostUser, middlewares.JWTMiddleware())
	e.GET("/users", handler.GetAllUser, middlewares.JWTMiddleware())
	e.PUT("/users/:id", handler.PutUser, middlewares.JWTMiddleware())
	e.DELETE("/users/:id", handler.DeleteAkun, middlewares.JWTMiddleware())
	e.GET("/users/:id", handler.SelectUserId, middlewares.JWTMiddleware())
}

func (deliv *Delivery) PostUser(c echo.Context) error {
	_, role := middlewares.ExtractToken(c)

	if role != "admin" {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("you dont have acces"))
	}
	fmt.Println(role)

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
	idToken, role := middlewares.ExtractToken(c) //coba

	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("param must be a number"))
	}

	if idConv == idToken || role == "admin" {
		var dataUpdate UserRequest
		errBind := c.Bind(&dataUpdate)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("error binding data"))
		}

		row, err := deliv.userUsecase.PutUser(toCore(dataUpdate), idConv)

		if err != nil || row == 0 {
			return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("failed to update data"))
		}
	}

	return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("you dont have acces"))

}

func (deliv *Delivery) DeleteAkun(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	_, role := middlewares.ExtractToken(c)

	if role != "admin" {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("you dont have acces"))
	}

	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("param must be a number"))
	}

	row, err := deliv.userUsecase.DeleteAkun(idConv)
	if err != nil || row == 0 {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("failed to delete account"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("succes delete account"))
}

func (deliv *Delivery) SelectUserId(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)

	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("param must be a number"))
	}

	result, err := deliv.userUsecase.GetUserId(idConv)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("failed to get data"))
	}
	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("succes get data", fromCore(result)))

}
