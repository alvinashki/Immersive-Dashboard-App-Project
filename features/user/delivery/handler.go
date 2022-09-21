package delivery

import (
	"fmt"
	"gp3/features/user"
	"gp3/utils/helper"
	"net/http"

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
}

func (deliv *Delivery) PostUser(c echo.Context) error {
	var dataRequest UserRequest
	errBind := c.Bind(&dataRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("error binding data"))
	}

	fmt.Println("error =", dataRequest)
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

	fmt.Println("ini dari handler =", fromCoreList(result))
	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("Succes get data", fromCoreList(result)))
}
