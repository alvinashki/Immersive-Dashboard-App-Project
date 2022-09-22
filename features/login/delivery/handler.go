package delivery

import (
	"gp3/features/login"
	"gp3/utils/helper"

	"github.com/labstack/echo/v4"
)

type Delivery struct {
	authUsecase login.UsecaseInterface
}

func New(e *echo.Echo, usecase login.UsecaseInterface) {

	handler := Delivery{
		authUsecase: usecase,
	}

	e.POST("/login", handler.Auth)

}

// ini komen
func (delivery *Delivery) Auth(c echo.Context) error {

	var req Request
	errBind := c.Bind(&req)
	if errBind != nil {
		return c.JSON(400, helper.FailedResponseHelper("wrong request"))
	}

	str := delivery.authUsecase.LoginAuthorized(req.Email, req.Password, req.Role) //coba
	if str == "please input email and password" || str == "email not found" || str == "wrong password" {
		return c.JSON(400, helper.FailedResponseHelper(str))
	} else if str == "failed to created token" {
		return c.JSON(500, helper.FailedResponseHelper(str))
	} else {
		return c.JSON(200, helper.SuccessDataResponseHelper("Login Success", str))
	}

}
