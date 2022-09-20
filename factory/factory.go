package factory

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	loginData "gp3/features/login/data"
	loginDelivery "gp3/features/login/delivery"
	loginUsecase "gp3/features/login/usecase"

	menteeData "gp3/features/mentee/data"
	menteeDelivery "gp3/features/mentee/delivery"
	menteeUsecase "gp3/features/mentee/usecase"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	loginDataFactory := loginData.New(db)
	loginUsecaseFactory := loginUsecase.New(loginDataFactory)
	loginDelivery.New(e, loginUsecaseFactory)

	menteeDataFactory := menteeData.New(db)
	menteeUsecaseFactory := menteeUsecase.New(menteeDataFactory)
	menteeDelivery.New(e, menteeUsecaseFactory)
}
